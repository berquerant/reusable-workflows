# Workflows

| Workflow | Description |
| :--- | :--- |
| [`actions-lock-verify.yml`](workflows/actions-lock-verify.md) | Verifies `.github/workflows/actions.lock` with `gh actions-lock --verify` |
| [`helm-chart-ci.yml`](workflows/helm-chart-ci.md) | Runs `helm lint --strict`, `ct lint`, and `ct install` on a `kind` cluster |
| [`helm-chart-snapshot.yml`](workflows/helm-chart-snapshot.md) | Generates Helm chart manifests and compares against golden snapshots |
| [`helm-chart-release.yml`](workflows/helm-chart-release.md) | Releases Helm charts via `chart-releaser` and pushes OCI packages to GHCR |
| [`goreleaser-release.yml`](workflows/goreleaser-release.md) | Builds and publishes GitHub Releases with GoReleaser v2 |
