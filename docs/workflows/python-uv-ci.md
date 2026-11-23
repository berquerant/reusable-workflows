# `python-uv-ci.yml`

Runs Python CI using `astral-sh/setup-uv` with dependency caching, `uv python install`, `uv sync`, and a configurable check command.

## Usage

```yaml
jobs:
  ci:
    permissions:
      contents: read # checkout repository for Python uv CI
    uses: berquerant/reusable-workflows/.github/workflows/python-uv-ci.yml@v0.8.0
    with:
      python-version: "3.14"
      command: make ci
```

## Inputs

| Input | Type | Required | Default | Description |
| :--- | :--- | :---: | :--- | :--- |
| `working-directory` | `string` | no | `.` | Working directory containing `pyproject.toml` |
| `python-version` | `string` | no | `""` | Python version to install via `uv` (uses `.python-version` when empty) |
| `sync-command` | `string` | no | `uv sync --locked --all-extras --dev` | Shell command to sync Python dependencies |
| `command` | `string` | no | `make ci` | Shell command to run CI checks |
| `submodules` | `boolean` | no | `false` | Whether to checkout git submodules |
| `timeout-minutes` | `number` | no | `20` | Job timeout in minutes |

## Permissions

Requires `contents: read` on the caller job.

## Required Status Checks

Register `<caller-job-id> / Python UV CI` (e.g. `ci / Python UV CI`) in repository rulesets.
