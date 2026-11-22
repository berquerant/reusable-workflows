# `helm-chart-release.yml`

Packages and releases Helm charts via `helm/chart-releaser-action` and pushes new chart versions to an OCI registry (default: GHCR).

## Usage

```yaml
jobs:
  release:
    permissions:
      contents: write
      packages: write
    uses: berquerant/reusable-workflows/.github/workflows/helm-chart-release.yml@v0.2.0
```

## Inputs

| Input | Type | Required | Default | Description |
| :--- | :--- | :---: | :--- | :--- |
| `helm-version` | `string` | no | `latest` | Helm version to install via `azure/setup-helm` |
| `oci-registry` | `string` | no | `""` | OCI registry URL (defaults to `oci://ghcr.io/<owner>/<repo>/charts`) |
| `timeout-minutes` | `number` | no | `10` | Job timeout in minutes |

## Permissions

Requires `contents: write` and `packages: write` on the caller job.
