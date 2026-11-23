# `emacs-ci.yml`

Runs Emacs Lisp CI across a matrix of Emacs versions using `purcell/setup-emacs`, with default `batch-byte-compile` and ERT test execution (or a custom test command), plus an aggregated `Emacs CI` summary job for repository rulesets.

## Usage

```yaml
jobs:
  ci:
    permissions:
      contents: read # checkout repository for Emacs CI
    uses: berquerant/reusable-workflows/.github/workflows/emacs-ci.yml@v0.9.0
    with:
      emacs-versions: '["27.1", "30.1"]'
```

## Inputs

| Input | Type | Required | Default | Description |
| :--- | :--- | :---: | :--- | :--- |
| `emacs-versions` | `string` | no | `'["30.1"]'` | JSON array of Emacs versions to test |
| `byte-compile-error-on-warn` | `boolean` | no | `true` | Whether byte-compile warnings fail the build when `test-command` is empty |
| `setup-command` | `string` | no | `""` | Optional shell command to run before byte-compile and tests |
| `test-command` | `string` | no | `""` | Shell command to run tests (runs `batch-byte-compile` and `tests/*.el` ERT when empty) |
| `submodules` | `boolean` | no | `false` | Whether to checkout git submodules |
| `timeout-minutes` | `number` | no | `5` | Job timeout in minutes |

## Permissions

Requires `contents: read` on the caller job.

## Required Status Checks

This workflow includes a `summary` job named `Emacs CI` that aggregates the Emacs version matrix. Register `<caller-job-id> / Emacs CI` (e.g. `ci / Emacs CI`) in repository rulesets.
