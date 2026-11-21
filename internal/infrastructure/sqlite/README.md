# SQLite Infrastructure

## Driver decision

The application uses `modernc.org/sqlite` as its SQLite driver.

### Why modernc.org/sqlite

- Implements SQLite in pure Go.
- Does not require CGO.
- Does not require a C compiler or SQLite development libraries.
- Supports Linux and Windows.
- Integrates with Go's standard `database/sql` package.
- Supports SQLite PRAGMAs, transactions and SQLite URI parameters.

### Version

`modernc.org/sqlite v1.62.0`

### Licensing

The driver is licensed under BSD-3-Clause.

The package also includes third-party components and dependencies with
their own licenses. Their notices must be reviewed and included in the
application's distribution materials where required.

Official license information:
https://github.com/modernc-org/sqlite

## Connection management

All database connections must be opened through the infrastructure
package's `Open` function.

The connection pool is configured with:

- Maximum open connections: 4
- Maximum idle connections: 4
- Connection lifetime: unlimited
- Connection idle lifetime: unlimited

## SQLite configuration

The following PRAGMAs are applied to every physical connection:

- `foreign_keys = ON`
- `busy_timeout = 5000`

`foreign_keys` must never be enabled only through a one-time
`Exec` call on `sql.DB`, because `sql.DB` represents a connection pool.

## Transactions

Transactions use the standard `database/sql` API.

Application services are responsible for transaction boundaries.
Infrastructure code provides database connectivity.

## Backup strategy

Backups must not rely on copying an active database file directly.

A dedicated backup implementation should use SQLite's Online Backup API
or another SQLite-supported consistent backup mechanism.

WAL mode is not enabled globally by default. Journal mode must be
considered when implementing backup and database lifecycle operations.
