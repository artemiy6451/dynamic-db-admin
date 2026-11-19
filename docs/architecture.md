# Architecture

## Product scope

Dynamic DB Admin is a desktop application for designing, inspecting and managing SQLite database files. SQLite is the only supported target database engine.

The application opens user-selected database files (`.db`, `.sqlite`, `.sqlite3`). It does not manage remote database servers or PostgreSQL connections.

## Technology

- Go and Wails v3 for the desktop application and backend services.
- Vue 3 Composition API, TypeScript, Vite, Vue Router and Pinia.
- Tailwind CSS v4 for styling.
- SQLite for user databases and a separate private application metadata store.
- Go `database/sql` with a SQLite driver. The driver choice is finalized when Sprint 2 implements the database layer.

## Architectural style

Use a modular monolith with explicit boundaries. Dependency injection is assembled manually in the application bootstrap; do not add a DI framework without a concrete need.

```text
Go backend
├── domain             # database, schema, table, column and relationship rules
├── application        # use cases and ports
├── infrastructure
│   ├── sqlite         # driver, introspection, DDL, migrations and data access
│   ├── storage        # private application metadata store
│   └── logging
├── adapters/wails     # Wails services exposed to the frontend
└── app                # composition root and lifecycle
```

The domain and application layers must not depend on Wails or Vue. SQLite-specific implementation belongs in infrastructure. Wails services translate typed frontend requests into application use cases.

## Database boundaries

There are two distinct kinds of SQLite files:

1. **User databases** are selected by the user and contain the data being managed.
2. **Application metadata storage** is private to Dynamic DB Admin. It stores preferences, recent files, window/UI metadata, ER diagram coordinates and application history.

Never mix application metadata into a user's database unless a future feature explicitly asks the user to do so.

## Backend/frontend contract

- Expose typed operations through generated Wails v3 bindings.
- Keep Wails binding methods thin; business rules belong in application/domain.
- The frontend sends typed operation descriptions, not arbitrary SQL fragments.
- The backend validates identifiers and constructs SQL. Bind all data values as parameters.
- Generated bindings are build artifacts; do not hand-edit them.

## SQLite-specific correctness

- Enable foreign-key enforcement for every physical database connection, not just one arbitrary pooled connection.
- Introspect existing databases instead of assuming they were created by this application.
- Support composite primary keys and tables without primary keys. Where safe row identity cannot be established, disable unsafe row-level mutations and explain why.
- Treat SQLite's dynamic typing and declared column types carefully; do not assume strict typing for every file.
- Schema changes that SQLite cannot express directly require a migration plan: create replacement table, copy data, restore indexes/triggers/views and constraints, validate, then replace the original.
- Analyze dependent objects before rebuilding a table. Never silently discard triggers, indexes, views or foreign-key relationships.
- Use transactions where supported and create a recoverable backup before risky schema migrations.
- Verify integrity and foreign-key consistency after structural changes.

## Data editor

The data editor combines a table/grid view and automatically generated forms. It must support safe CRUD, sorting, filtering, search and pagination. Related-record selection should be driven by foreign-key metadata.

## ER designer

The ER designer visualizes actual foreign-key relationships. Users can arrange table nodes and save diagram positions in application metadata storage. A visual connector must not imply a database constraint that does not exist.

## Security and safety

- Parameterize values and safely quote validated identifiers.
- Never accept raw SQL fragments from the frontend for schema/data operations.
- Confirm destructive actions.
- Preview risky migrations and preserve recoverability.
- Do not log database contents, secrets or sensitive values.
- Do not assume a database file is trusted; handle malformed or hostile files defensively.

## Frontend organization

Use feature-oriented modules under `frontend/src/features` (databases, schema explorer, table designer, data editor, relationship designer and settings). Shared primitives belong in `components/ui`; shared shell components belong in `components/layout`. Centralize Wails binding calls in frontend services and stores where appropriate.

## Out of scope for the initial version

A general-purpose SQL editor, PostgreSQL/MySQL support, remote database server management and server-mode deployment are not part of the initial product scope.
