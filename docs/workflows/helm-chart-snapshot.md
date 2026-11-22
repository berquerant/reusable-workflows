# `helm-chart-snapshot.yml`

Generates rendered Helm chart manifests for each test values file and compares them against golden snapshots using `berquerant/k8s-object-diff`.

## Usage

```yaml
jobs:
  snapshot:
    permissions:
      contents: read
    uses: berquerant/reusable-workflows/.github/workflows/helm-chart-snapshot.yml@v0.2.0
    with:
      chart-name: fluent-sandbox
```

## Inputs

| Input | Type | Required | Default | Description |
| :--- | :--- | :---: | :--- | :--- |
| `chart-name` | `string` | yes | — | Chart name under `tests/` (e.g. `fluent-sandbox`) |
| `generate-script` | `string` | no | `./tests/generate.sh` | Path to the manifest generation script |
| `timeout-minutes` | `number` | no | `5` | Job timeout in minutes |

## Permissions

Requires `contents: read` on the caller job.
