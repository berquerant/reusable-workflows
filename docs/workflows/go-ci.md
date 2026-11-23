# `go-ci.yml`

Runs standard Go CI jobs (`lint`, `test`, `build`, and `govulncheck`) in parallel using `actions/setup-go`, plus an aggregated `Go CI` summary job for repository rulesets.

For repositories using `mise` (`jdx/mise-action`), use [`go-mise-ci.yml`](go-mise-ci.md).

## Usage

```yaml
jobs:
  ci:
    permissions:
      contents: read # checkout repository for Go CI
    uses: berquerant/reusable-workflows/.github/workflows/go-ci.yml@v0.5.0
    with:
      lint-command: make lint
      test-command: make test
      build-command: make
```

## Inputs

| Input | Type | Required | Default | Description |
| :--- | :--- | :---: | :--- | :--- |
| `working-directory` | `string` | no | `.` | Working directory containing the Go module |
| `go-version-file` | `string` | no | `go.mod` | Path to `go.mod` relative to repository root |
| `cache-dependency-path` | `string` | no | `go.sum` | Path to `go.sum` relative to repository root |
| `submodules` | `boolean` | no | `false` | Whether to checkout git submodules |
| `lint-command` | `string` | no | `""` | Shell command for linting (skipped when empty) |
| `test-command` | `string` | no | `make test` | Shell command for testing (skipped when empty) |
| `build-command` | `string` | no | `""` | Shell command for building (skipped when empty) |
| `run-govulncheck` | `boolean` | no | `true` | Whether to run `berquerant/govulncheck-action` |
| `timeout-minutes` | `number` | no | `10` | Job timeout in minutes |

## Permissions

Requires `contents: read` on the caller job.

## Required Status Checks

This workflow includes a `summary` job named `Go CI` that aggregates `lint`, `test`, `build`, and `govulncheck`. Register `<caller-job-id> / Go CI` (e.g. `ci / Go CI`) in repository rulesets.
