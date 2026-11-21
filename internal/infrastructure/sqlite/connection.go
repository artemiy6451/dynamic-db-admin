package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"

	sqlitedriver "modernc.org/sqlite"
)

const (
	driverName   = "sqlite"
	maxOpenConns = 4
	maxIdleConns = 4
	busyTimeout  = 5000
)

var strictPragmasOnce sync.Once

func Open(ctx context.Context, path string) (*sql.DB, error) {
	if strings.TrimSpace(path) == "" {
		return nil, fmt.Errorf("database path cannot be empty")
	}

	strictPragmasOnce.Do(func() {
		sqlitedriver.StrictPragmas(true)
	})

	absolutePath, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("resolve database path: %w", err)
	}

	dsn, err := buildDSN(absolutePath)
	if err != nil {
		return nil, fmt.Errorf("build database DSN: %w", err)
	}

	db, err := sql.Open(driverName, dsn)
	if err != nil {
		return nil, fmt.Errorf("open SQLite database: %w", err)
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxLifetime(0)
	db.SetConnMaxIdleTime(0)

	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := db.PingContext(pingCtx); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("ping SQLite database: %w", err)
	}

	return db, nil
}

func buildDSN(path string) (string, error) {
	path = filepath.ToSlash(path)

	if runtime.GOOS == "windows" &&
		len(path) >= 2 && path[1] == ':' {
		path = "/" + path
	}

	query := url.Values{}
	query.Add("_pragma", "foreign_keys(1)")
	query.Add("_pragma", fmt.Sprintf("busy_timeout(%d)", busyTimeout))

	uri := url.URL{
		Scheme:   "file",
		Path:     path,
		RawQuery: query.Encode(),
	}

	return uri.String(), nil
}
