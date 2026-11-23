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

## Documentation Version References

- Usage examples in `docs/workflows/<name>.md` and `docs/actions/README.md` pin to the release tag in which that workflow/action was **last changed** (not necessarily the latest tag). Mismatched tags across docs are expected; do not bump unchanged ones.
- When a workflow/action changes, update its doc example to the new tag after the release is tagged (follow-up commit).
- `docs/README.md` uses `@v0.5.0` only as an illustrative example of pinning; it need not track releases.

## Composite Actions

- Reusable workflows in this repository do **not** call the composite actions under `actions/`. Inside a reusable workflow, a local reference (`uses: ./actions/...`) resolves against the **caller's** workspace, not this repository. An external reference (`berquerant/reusable-workflows/actions/...@<tag>`) would require the tag to exist before the workflow that uses it is released.
- Setup steps (Go, Rust, etc.) are therefore implemented inline in each workflow. Duplication between workflows (e.g. `prepare` matrix generation, `summary` result checks) is accepted for this reason; when changing one, check its siblings.
- Composite actions are provided for callers to use directly. Keep their `uses:` references pinned to a commit SHA with a version comment.
