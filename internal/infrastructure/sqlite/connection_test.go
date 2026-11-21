package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"path/filepath"
	"testing"
)

func TestOpenPingClose(t *testing.T) {
	path := filepath.Join(t.TempDir(), "test.db")

	db, err := Open(context.Background(), path)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}

	if err := db.PingContext(context.Background()); err != nil {
		t.Fatalf("PingContext() error = %v", err)
	}

	if err := db.Close(); err != nil {
		t.Fatalf("Close() error = %v", err)
	}
}

func TestForeignKeysEnabledOnEveryConnection(t *testing.T) {
	const connectionCount = 4

	path := filepath.Join(t.TempDir(), "foreign_keys.db")

	db, err := Open(context.Background(), path)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer db.Close()

	db.SetMaxOpenConns(connectionCount)

	_, err = db.Exec(`
		CREATE TABLE parent (
			id INTEGER PRIMARY KEY
		)
	`)
	if err != nil {
		t.Fatalf("create parent table: %v", err)
	}

	_, err = db.Exec(`
		CREATE TABLE child (
			id INTEGER PRIMARY KEY,
			parent_id INTEGER NOT NULL,
			FOREIGN KEY (parent_id) REFERENCES parent(id)
		)
	`)
	if err != nil {
		t.Fatalf("create child table: %v", err)
	}

	connections := make([]*sql.Conn, 0, connectionCount)

	for i := range connectionCount {
		conn, err := db.Conn(context.Background())
		if err != nil {
			t.Fatalf("acquire connection %d: %v", i, err)
		}

		connections = append(connections, conn)
	}

	defer func() {
		for _, conn := range connections {
			_ = conn.Close()
		}
	}()

	for i, conn := range connections {
		var enabled int

		err := conn.QueryRowContext(
			context.Background(),
			"PRAGMA foreign_keys",
		).Scan(&enabled)
		if err != nil {
			t.Fatalf("connection %d: read foreign_keys: %v", i, err)
		}

		if enabled != 1 {
			t.Errorf(
				"connection %d: foreign_keys = %d, want 1",
				i,
				enabled,
			)
		}

		_, err = conn.ExecContext(
			context.Background(),
			"INSERT INTO child (id, parent_id) VALUES (?, ?)",
			i+1,
			999999,
		)

		if err == nil {
			t.Errorf(
				"connection %d: expected foreign key violation",
				i,
			)
		}
	}
}

func TestTransactionRollback(t *testing.T) {
	path := filepath.Join(t.TempDir(), "transaction.db")

	db, err := Open(context.Background(), path)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer db.Close()

	_, err = db.Exec(`
		CREATE TABLE items (
			id INTEGER PRIMARY KEY,
			name TEXT NOT NULL
		)
	`)
	if err != nil {
		t.Fatalf("create table: %v", err)
	}

	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		t.Fatalf("BeginTx() error = %v", err)
	}

	_, err = tx.Exec(
		"INSERT INTO items (id, name) VALUES (?, ?)",
		1,
		"temporary",
	)
	if err != nil {
		t.Fatalf("insert: %v", err)
	}

	if err := tx.Rollback(); err != nil {
		t.Fatalf("Rollback() error = %v", err)
	}

	var count int

	err = db.QueryRow(
		"SELECT COUNT(*) FROM items",
	).Scan(&count)
	if err != nil {
		t.Fatalf("count rows: %v", err)
	}

	if count != 0 {
		t.Errorf("row count = %d, want 0", count)
	}
}

func TestOpenRejectsEmptyPath(t *testing.T) {
	_, err := Open(context.Background(), " ")

	if err == nil {
		t.Fatal("expected error for empty path")
	}
}

func ExampleOpen() {
	db, err := Open(context.Background(), "example.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("SQLite database opened")
	// Output: SQLite database opened
}
