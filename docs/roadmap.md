# Product Roadmap

Dynamic DB Admin is a desktop SQLite database administration and design tool. This roadmap describes planned work; unchecked items are not implemented functionality.

## Sprint 1 — Project Foundation

- [x] Initialize Wails v3 and Vue 3 frontend.
- [x] Establish Go package boundaries and frontend source layout.
- [x] Add application shell, sidebar, dashboard and initial routes.
- [x] Add reusable base UI components.
- [ ] Align project documentation and GitHub planning with the SQLite product scope.
- [ ] Verify development and production build configuration.

## Sprint 2 — SQLite Database Management

- [ ] Select and integrate the SQLite driver.
- [ ] Implement create/open/close operations for SQLite files.
- [ ] Support multiple open database files and active database selection.
- [ ] Add recent files and application metadata storage.
- [ ] Handle invalid, read-only, locked and unsupported database files.

## Sprint 3 — Database Explorer

- [ ] Introspect tables, views, indexes and triggers.
- [ ] Read columns, declared types, defaults and constraints.
- [ ] Read primary keys and foreign-key relationships, including composite keys.
- [ ] Build a navigable schema explorer.
- [ ] Refresh metadata after schema changes.

## Sprint 4 — Schema Designer

- [ ] Create, rename and drop tables.
- [ ] Add, rename and remove columns where supported.
- [ ] Configure primary keys, foreign keys, unique, not-null, default and check constraints.
- [ ] Create and manage indexes.
- [ ] Implement safe table-rebuild migrations for unsupported ALTER operations.
- [ ] Preserve dependent indexes, triggers and views.
- [ ] Back up and validate before risky migrations.

## Sprint 5 — Relationship Designer

- [ ] Display tables as ER diagram nodes.
- [ ] Visualize actual foreign-key constraints.
- [ ] Navigate from relationships to table/column details.
- [ ] Arrange nodes and persist diagram coordinates in application metadata.
- [ ] Refresh diagram after schema changes.

## Sprint 6 — Data Editor

- [ ] Implement schema-driven data grid.
- [ ] Add, edit and delete rows safely.
- [ ] Generate dynamic record forms from column metadata.
- [ ] Handle NULL, blobs, dates, booleans and SQLite storage classes.
- [ ] Respect primary keys, composite keys and tables without stable row identity.

## Sprint 7 — Advanced Data Management

- [ ] Add sorting, filtering and text search.
- [ ] Add pagination and bounded data loading.
- [ ] Add related-record selection for foreign keys.
- [ ] Improve large-table performance and editing feedback.

## Sprint 8 — Safety and Data Integrity

- [ ] Add reliable backup and restore workflows.
- [ ] Add integrity and foreign-key checks.
- [ ] Add migration preview and recovery flows.
- [ ] Add operation history and safe error reporting.

> Safety requirements are not postponed until this sprint: backup and integrity safeguards are required as soon as schema migrations are implemented in Sprint 4.

## Sprint 9 — Testing and Stabilization

- [ ] Add domain and application unit tests.
- [ ] Add SQLite integration tests using temporary database files.
- [ ] Cover malformed databases, composite keys, triggers, views and no-PK tables.
- [ ] Test migrations and recovery paths.
- [ ] Profile large tables and improve reliability.

## Sprint 10 — Release and Distribution

- [ ] Configure Windows and Linux packaging.
- [ ] Verify platform-specific WebView/runtime requirements.
- [ ] Add release automation and versioning.
- [ ] Write user documentation and release notes.
- [ ] Publish the first stable release.

## Later / not in initial release

- SQL editor and query workspace.
- Import/export workflows.
- Advanced trigger and view editors.
- Additional database engines (not currently planned).
