# `goreleaser-mise-release.yml`

Runs GoReleaser v2 to build and publish GitHub Releases using `jdx/mise-action` and optional pre-release checks.

For repositories that do not allow `jdx/mise-action` and use `actions/setup-go` directly, use [`goreleaser-release.yml`](goreleaser-release.md).

## Usage

```yaml
jobs:
  release:
    permissions:
      contents: write # create GitHub releases and upload release assets
    uses: berquerant/reusable-workflows/.github/workflows/goreleaser-mise-release.yml@v0.5.0
    with:
      pre-release-command: mise run
      version-package: github.com/berquerant/example/version
```

## Inputs

| Input | Type | Required | Default | Description |
| :--- | :--- | :---: | :--- | :--- |
| `pre-release-command` | `string` | no | `""` | Optional shell command to run before GoReleaser |
| `version-package` | `string` | no | `""` | Optional `VERSION_PACKAGE` environment variable for GoReleaser |
| `goreleaser-version` | `string` | no | `~> v2` | GoReleaser version constraint (tracks latest v2 by default; pin for reproducibility) |
| `goreleaser-args` | `string` | no | `release --clean` | Arguments passed to GoReleaser |
| `timeout-minutes` | `number` | no | `15` | Job timeout in minutes |

## Environment Variables Set

- `REVISION`: Short Git commit SHA (`git rev-parse --short HEAD`)
- `VERSION`: Exact Git tag at `HEAD` when present (`git describe --tags --abbrev=0 --exact-match`)
- `VERSION_PACKAGE`: Set when `inputs.version-package` is non-empty

## Security

`pre-release-command` is executed as a shell command in a job holding `contents: write`. Pass only trusted, static values, and call this workflow only from protected branches or release tags.

## Permissions

Requires `contents: write` on the caller job.
