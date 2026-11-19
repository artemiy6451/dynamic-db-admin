# Sprint 02 — SQLite Database Management

**Status:** Planned  
**Branch:** `dev`

## Goal

Implement the first real database workflow: create, open, track, switch between and close SQLite database files.

## Scope

1. Select and integrate a Go SQLite driver.
2. Open and close SQLite database files through a backend application service.
3. Support multiple simultaneously open database files and an active database.
4. Create a separate private application metadata store.
5. Build the Databases workspace in Vue and connect it to generated Wails v3 bindings.

## Suggested implementation order

1. Driver evaluation and a small `database/sql` proof of concept.
2. Domain types and application use cases for database file lifecycle.
3. SQLite infrastructure implementation and integration tests.
4. Private metadata store for recent files and preferences.
5. Wails adapter DTOs/methods and binding generation.
6. Frontend service/store and Databases page integration.

## Safety and correctness requirements

- Opening a nonexistent path must not silently create a database; creation is an explicit action.
- Handle read-only, invalid, locked and already-open files.
- Close database handles during explicit close and application shutdown.
- Keep user database files separate from the private application metadata store.
- Enable foreign-key enforcement on every physical connection.
- Do not accept arbitrary SQL from the frontend.
- Avoid storing database contents or credentials in recent-file metadata.

## Acceptance criteria

- [ ] A new SQLite database can be created at a user-selected path.
- [ ] Existing SQLite files can be opened.
- [ ] Multiple files can remain open simultaneously.
- [ ] The active database can be changed.
- [ ] A database can be closed without affecting other open files.
- [ ] Recent file references persist across application restarts.
- [ ] Invalid/read-only/locked-file errors are handled.
- [ ] Wails v3 bindings are regenerated and used by the frontend.
- [ ] Databases page shows open and recent files with loading, empty and error states.
- [ ] Unit and integration tests cover the database lifecycle.

## Out of scope

- Schema introspection (Sprint 3).
- Schema editing and migrations (Sprint 4).
- ER diagram (Sprint 5).
- Data grid and record forms (Sprint 6).
