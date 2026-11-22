# `goreleaser-release.yml`

Runs GoReleaser v2 to build and publish GitHub Releases, with optional Go or `mise` setup and pre-release checks.

## Usage

```yaml
jobs:
  release:
    permissions:
      contents: write
    uses: berquerant/reusable-workflows/.github/workflows/goreleaser-release.yml@v0.3.0
    with:
      setup-mode: mise
      pre-release-command: mise run
      version-package: github.com/berquerant/example/version
```

## Inputs

| Input | Type | Required | Default | Description |
| :--- | :--- | :---: | :--- | :--- |
| `setup-mode` | `string` | no | `go` | Toolchain setup mode (`go` or `mise`) |
| `go-version-file` | `string` | no | `go.mod` | Path to `go.mod` when `setup-mode` is `go` |
| `pre-release-command` | `string` | no | `""` | Optional shell command to run before GoReleaser |
| `version-package` | `string` | no | `""` | Optional `VERSION_PACKAGE` environment variable for GoReleaser |
| `goreleaser-version` | `string` | no | `~> v2` | GoReleaser version constraint |
| `goreleaser-args` | `string` | no | `release --clean` | Arguments passed to GoReleaser |
| `timeout-minutes` | `number` | no | `15` | Job timeout in minutes |

## Environment Variables Set

- `REVISION`: Short Git commit SHA (`git rev-parse --short HEAD`)
- `VERSION`: Exact Git tag at `HEAD` when present (`git describe --tags --abbrev=0 --exact-match`)
- `VERSION_PACKAGE`: Set when `inputs.version-package` is non-empty

## Permissions

Requires `contents: write` on the caller job.
