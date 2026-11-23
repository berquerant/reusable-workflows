# Workflows

| Workflow | Description |
| :--- | :--- |
| [`actions-lock-verify.yml`](workflows/actions-lock-verify.md) | Verifies `.github/workflows/actions.lock` with `gh actions-lock --verify` |
| [`helm-chart-ci.yml`](workflows/helm-chart-ci.md) | Runs `helm lint --strict`, `ct lint`, and `ct install` on a `kind` cluster |
| [`helm-chart-snapshot.yml`](workflows/helm-chart-snapshot.md) | Generates Helm chart manifests and compares against golden snapshots |
| [`helm-chart-release.yml`](workflows/helm-chart-release.md) | Releases Helm charts via `chart-releaser` and pushes OCI packages to GHCR |
| [`goreleaser-release.yml`](workflows/goreleaser-release.md) | Builds and publishes GitHub Releases with GoReleaser v2 (`actions/setup-go`) |
| [`goreleaser-mise-release.yml`](workflows/goreleaser-mise-release.md) | Builds and publishes GitHub Releases with GoReleaser v2 (`jdx/mise-action`) |
| [`go-ci.yml`](workflows/go-ci.md) | Runs Go `lint`, `test`, `build`, and `govulncheck` with `actions/setup-go` |
| [`go-mise-ci.yml`](workflows/go-mise-ci.md) | Runs Go `lint`, `test`, `build`, and `govulncheck` with `jdx/mise-action` |
| [`rust-ci.yml`](workflows/rust-ci.md) | Runs Rust `lint`, `test`, `build`, and `cargo audit` with `Swatinem/rust-cache` |

## Security Notes for Callers

- **Pin references**: Pin workflow references to a specific release tag (e.g. `@v0.5.0`) or full commit SHA.
- **Avoid `pull_request_target`**: Never call these workflows from `pull_request_target` or pass untrusted user inputs (such as `workflow_dispatch` inputs or PR titles/bodies) into `*-command` inputs.
- **Restrict release triggers**: Only call `goreleaser-release.yml`, `goreleaser-mise-release.yml`, and `helm-chart-release.yml` from protected branch pushes (`main`) or release tag pushes (`v*`).
