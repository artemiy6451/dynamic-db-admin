# Desktop build

The Wails v3 build embeds the Vite production output from `frontend/dist` into the Go executable. The application does not need a separate frontend server after it has been built.

## Development

```bash
wails3 dev
```

Wails runs the development build and starts Vite as a separate development server. This is expected only in development mode.

## Production build

```bash
wails3 build
```

The executable is written to `bin/`. The root Taskfile delegates to the platform-specific Taskfile under `build/linux` or `build/windows`.

## Linux

Build on a supported Linux distribution with a C compiler, pkg-config, GTK4 and WebKitGTK 6.0 development packages installed. On Arch Linux, install the matching GTK4, WebKitGTK 6.0 and base development toolchain packages.

The resulting application uses the system WebKitGTK runtime; it is not bundled into the executable.

## Windows

Build on Windows with Go, Node.js/npm and the Wails v3 CLI installed. The application requires the Microsoft Edge WebView2 Runtime on the target machine. The Windows build task uses CGO-disabled native Go compilation by default; validate this against the Wails and project dependency requirements when adding native CGO dependencies.

A Windows executable must be built and launched on Windows (or a supported Windows CI runner) before marking Windows release verification complete. A Linux cross-compile alone is not proof of runtime compatibility.

## Data locations

- User-selected SQLite database files stay at the paths chosen by the user. The app must not copy them beside the executable.
- Application configuration is loaded from the per-user configuration directory (or the explicit `DYNAMIC_DB_ADMIN_CONFIG` path).
- Future application metadata storage belongs in a per-user application data/configuration directory, never in `bin/` or the installation directory.
- Build outputs and generated frontend assets are not user data.
