# AGENTS.md

## Versioning

- Start at `v0.1.0` (breaking changes may occur before `v1.0.0`).
- Callers should pin to a specific version tag (e.g. `@v0.1.0`) or commit SHA.

## Design Conventions

- **Least-privilege permissions**: Declare top-level `permissions: {}` and grant only the minimum required permissions per job.
- **Checkout safety**: Set `persist-credentials: false` on `actions/checkout`.
- **Timeouts**: Set explicit `timeout-minutes` on every job.
- **Concurrency**: Managed at the caller workflow's top level.
- **Supply-chain integrity**: All actions referenced inside this repository are locked via `.github/workflows/actions.lock` and verified in CI (`gh actions-lock --verify`).
- **Linting**: Run `mise run lint` (`actionlint`, `ghalint`, `zizmor`) before committing workflow changes.
- **Documentation**: Document each workflow in `docs/workflows/<name>.md` and add an entry to `docs/README.md`.
