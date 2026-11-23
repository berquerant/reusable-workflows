# `rust-ci.yml`

Runs standard Rust CI jobs (`lint`, `test`, `build`, and `cargo audit`) in parallel using `rustup`, `Swatinem/rust-cache`, and `actions/cache` (for `cargo-audit`), plus an aggregated `Rust CI` summary job for repository rulesets.

## Usage

```yaml
jobs:
  ci:
    permissions:
      contents: read # checkout repository for Rust CI
    uses: berquerant/reusable-workflows/.github/workflows/rust-ci.yml@v0.7.1
    with:
      lint-command: make lint
      test-command: make test
      build-command: make build
```

## Inputs

| Input | Type | Required | Default | Description |
| :--- | :--- | :---: | :--- | :--- |
| `working-directory` | `string` | no | `.` | Working directory containing `Cargo.toml` |
| `toolchain` | `string` | no | `stable` | Rust toolchain to install via `rustup` |
| `components` | `string` | no | `clippy rustfmt` | Space-separated `rustup` components to install |
| `submodules` | `boolean` | no | `false` | Whether to checkout git submodules |
| `lint-command` | `string` | no | `""` | Shell command for linting (skipped when empty) |
| `test-command` | `string` | no | `make test` | Shell command for testing (skipped when empty) |
| `build-command` | `string` | no | `""` | Shell command for building (skipped when empty) |
| `run-audit` | `boolean` | no | `true` | Whether to run `cargo audit` |
| `timeout-minutes` | `number` | no | `20` | Job timeout in minutes |

## Security

`lint-command`, `test-command` and `build-command` are executed as shell commands on the runner. Pass only trusted, static values; never pass untrusted input (PR titles/bodies, `workflow_dispatch` inputs).

## Permissions

Requires `contents: read` on the caller job.

## Required Status Checks

This workflow includes a `summary` job named `Rust CI` that aggregates `lint`, `test`, `build`, and `audit`. Register `<caller-job-id> / Rust CI` (e.g. `ci / Rust CI`) in repository rulesets.
