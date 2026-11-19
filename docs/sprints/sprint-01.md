# Sprint 01 — Project Foundation

**Status:** In progress  
**Branch:** `dev`

## Goal

Establish the desktop application foundation and reusable frontend/backend structure for future SQLite database management features.

## Scope

- Wails v3 application bootstrap and lifecycle.
- Go package boundaries and manual dependency initialization.
- Vue 3 + TypeScript + Vite frontend.
- Pinia, Vue Router and Tailwind CSS.
- Main application shell, sidebar, content area and dashboard.
- Reusable UI primitives and initial Wails binding integration.
- Development/build workflow and project documentation.

## Completed foundation visible in the repository

- Wails v3 dependency and application bootstrap.
- Go packages for app, domain, application, infrastructure and Wails adapter.
- Vue root component and app entry point.
- Main layout and sidebar.
- Dashboard, Connections placeholder and Settings page.
- Initial routes.
- BaseButton, BaseInput, LoadingState and ErrorNotification components.
- Greeting feature composable and Wails service.

## Remaining work

- Replace the PostgreSQL-oriented placeholder with SQLite-oriented Databases navigation and dashboard language.
- Update README, architecture and roadmap to describe the selected product scope.
- Remove outdated Wails v2 references from planning issues.
- Verify frontend type-check and production build.
- Verify Wails v3 development and production asset embedding.
- Document Windows build requirements separately from the Linux development environment.

## Acceptance criteria

- [ ] Application starts in Wails v3 development mode.
- [ ] Main layout and sidebar render correctly.
- [ ] Dashboard, Databases and Settings routes are reachable.
- [ ] Reusable UI components are available.
- [ ] Vue can invoke generated Wails v3 bindings.
- [ ] Frontend type-check and production build pass.
- [ ] Documentation describes SQLite-only scope accurately.
- [ ] No PostgreSQL/server-connection assumptions remain in the user-facing shell.

## Out of scope

- Actual SQLite file management (Sprint 2).
- Schema introspection (Sprint 3).
- Schema editing and migrations (Sprint 4).
- ER diagram and data editing (later sprints).

## Notes

Do not mark this sprint complete until the acceptance criteria are verified in the working development environment. A repository snapshot alone does not prove runtime correctness.
