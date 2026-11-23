# AGENTS.md

## Versioning

- Start at `v0.1.0` (breaking changes may occur before `v1.0.0`).
- Callers should pin to a specific version tag (e.g. `@v0.1.0`) or commit SHA.

## Design Conventions

- **Least-privilege permissions**: Declare top-level `permissions: {}` and grant only the minimum required permissions per job, documenting the purpose of each permission with an inline YAML comment.
- **Inline rationale comments**: Add inline YAML comments explaining notable non-default settings (e.g. `fetch-depth: 0`, `cache: false`).
- **Checkout safety**: Set `persist-credentials: false` on `actions/checkout`.
- **Timeouts**: Set explicit `timeout-minutes` on every job.
- **Concurrency**: Managed at the caller workflow's top level.
- **Supply-chain integrity**: Keep `uses:` references consistent with `.github/workflows/actions.lock` (via `gh actions-lock`) and verify `.github/workflows/actions.lock` in CI (`gh actions-lock --verify`).
- **Linting**: Run `mise run lint` (`actionlint`, `ghalint`, `zizmor`) before committing workflow changes.
- **Documentation**: Document each workflow in `docs/workflows/<name>.md` and add an entry to `docs/README.md`.
- **Cross-repository edits**: When updating another repository to call workflows from this repository, create a git worktree under `.git-worktree/` inside this repository.
