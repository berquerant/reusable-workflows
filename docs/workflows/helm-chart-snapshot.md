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

## Repository Layout

The caller repository is expected to provide:

- `<generate-script> ls <chart-name>`: prints the values files under `tests/<chart-name>/values/`.
- `<generate-script> gen <chart-name> generated`: writes manifests to `tests/<chart-name>/generated/<value>`.
- `tests/<chart-name>/golden/<value>`: golden snapshots to compare against.

`generate-script` is executed as-is; pass only a trusted path.

## Permissions

Requires `contents: read` on the caller job.

## Required Status Checks

This workflow includes a `summary` job named `Snapshot` that aggregates the results of `generate-matrix`, `generate-manifests`, and all matrix `test` jobs. Register `<caller-job-id> / Snapshot` (e.g. `snapshot / Snapshot`) in repository rulesets instead of listing individual matrix entries.
