# `actions-lock-verify.yml`

Verifies `.github/workflows/actions.lock` with `gh actions-lock --verify`.

## Usage

```yaml
jobs:
  actions-lock-verify:
    permissions:
      contents: read
    uses: berquerant/reusable-workflows/.github/workflows/actions-lock-verify.yml@v0.1.0
```

## Inputs

| Input | Type | Default | Description |
| :--- | :--- | :--- | :--- |
| `timeout-minutes` | `number` | `5` | Job timeout in minutes |
| `working-directory` | `string` | `"."` | Working directory containing `.github/workflows/actions.lock` |

## Permissions

Requires `contents: read` on the caller job.
