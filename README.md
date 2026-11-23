# reusable-workflows

Shared GitHub Actions reusable workflows for `berquerant` repositories.

See [docs/README.md](docs/README.md) for available workflows and usage.

## Local Development

Install tools with [mise](https://mise.jdx.dev/) and run linters:

```bash
mise install
mise run lint
```

## Smoke Tests

[`smoke-test.yml`](.github/workflows/smoke-test.yml) runs on every push and pull request to `main`. It calls the Go, Rust, and Python CI workflows against minimal fixtures in [`tests/fixtures/`](tests/fixtures) to verify that they actually run. Helm, release, Emacs, and `go-mise-ci` workflows are not covered. See `AGENTS.md` for maintenance rules.

## Versioning of Examples

Usage examples in `docs/` pin to the release in which each workflow or action last changed, so tags differ between pages. Check the [releases](https://github.com/berquerant/reusable-workflows/releases) page for the newest tag. See `AGENTS.md` for the maintenance rule.
