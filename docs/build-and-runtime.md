# Runtime and data storage requirements

## Linux

The Linux desktop build links against native desktop libraries. Install a C compiler, `pkg-config`, GTK4 development headers and WebKitGTK 6.0 development headers before building.

On Arch Linux, install the corresponding packages from the official repositories (typically `base-devel`, `pkgconf`, `gtk4` and `webkitgtk-6.0`).

The target machine must have compatible GTK4 and WebKitGTK 6.0 runtime libraries installed. These libraries are not embedded into the Go executable.

## Windows

The Windows application uses Microsoft Edge WebView2. The target machine must have the WebView2 Runtime installed. The Evergreen Runtime is the recommended deployment option; the installer should detect or bootstrap it before the application is launched.

Build and launch verification must be performed on Windows or a supported Windows CI runner. A cross-compiled executable alone does not prove runtime compatibility.

## Frontend assets

The Vite production output is written to `frontend/dist` and embedded into the Go binary by the `//go:embed all:frontend/dist` directive in `main.go`. Production execution serves these embedded files; it does not start Vite or require a separate frontend server.

The `wails3 dev` workflow is different by design: it starts a development build and a Vite dev server for hot reload.

## User data locations

- User-selected SQLite files remain at the exact locations selected by the user.
- Application configuration is read from the operating system's per-user configuration directory, under `Dynamic DB Admin/config.yaml`. Set `DYNAMIC_DB_ADMIN_CONFIG` to override the config path for local development.
- The application must never use its executable directory as a default location for user databases, preferences, metadata or backups.
- The private metadata database is not implemented yet. When added, it must be placed under a per-user application data/configuration directory and remain separate from user-selected SQLite files.
