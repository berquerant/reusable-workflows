# `helm-chart-ci.yml`

Runs Helm chart linting and installation tests on a `kind` cluster (`helm lint --strict`, `ct lint`, `ct install`).

## Usage

```yaml
jobs:
  ci:
    permissions:
      contents: read
    uses: berquerant/reusable-workflows/.github/workflows/helm-chart-ci.yml@v0.2.0
    with:
      chart-dir: charts/fluent-sandbox
```

## Inputs

| Input | Type | Required | Default | Description |
| :--- | :--- | :---: | :--- | :--- |
| `chart-dir` | `string` | yes | — | Path to the Helm chart directory (e.g. `charts/my-chart`) |
| `helm-version` | `string` | no | `latest` | Helm version to install via `azure/setup-helm` |
| `timeout-minutes` | `number` | no | `10` | Job timeout in minutes |

## Permissions

Requires `contents: read` on the caller job.
