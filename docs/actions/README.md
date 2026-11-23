# Composite Actions

Callers should pin action references to a release tag (e.g. `@v0.10.0`) or full commit SHA.

## `actions/setup-go`

Sets up Go with `actions/setup-go` and runs `go mod download`.

```yaml
- uses: berquerant/reusable-workflows/actions/setup-go@v0.10.0
```

| Input | Required | Default | Description |
| :--- | :---: | :--- | :--- |
| `go-version-file` | no | `go.mod` | Path to `go.mod` relative to repository root |
| `cache-dependency-path` | no | `go.sum` | Path to `go.sum` relative to repository root |
| `cache` | no | `"true"` | Whether to cache Go modules and build cache |
| `working-directory` | no | `.` | Working directory containing `go.mod` |

## `actions/setup-rust`

Sets up the Rust toolchain via `rustup` and caches Cargo artifacts with `Swatinem/rust-cache`.

```yaml
- uses: berquerant/reusable-workflows/actions/setup-rust@v0.10.0
```

| Input | Required | Default | Description |
| :--- | :---: | :--- | :--- |
| `toolchain` | no | `stable` | Rust toolchain to install via `rustup` |
| `components` | no | `clippy rustfmt` | Space-separated `rustup` components to install |
| `workspaces` | no | `.` | Workspaces path passed to `Swatinem/rust-cache` |

## `actions/shellcheck`

Registers `shellcheck-matcher.json` and runs `shellcheck -f gcc` on tracked `*.sh` files (or specified `files`).

```yaml
- uses: berquerant/reusable-workflows/actions/shellcheck@v0.10.0
```

| Input | Required | Default | Description |
| :--- | :---: | :--- | :--- |
| `files` | no | `""` | Space-separated file paths to check (auto-discovers tracked `*.sh` files when empty) |
| `exclude` | no | `""` | Comma-separated ShellCheck error codes to exclude |

## `actions/setup-buildx-mirror`

Sets up Docker Buildx configured with `mirror.gcr.io` as a mirror for `docker.io`.

```yaml
- uses: berquerant/reusable-workflows/actions/setup-buildx-mirror@v0.10.0
```
