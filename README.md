# GitHub Action - Inject secrets into file

[![Action Template](https://img.shields.io/badge/Action%20Template-Go%20Container%20Action-blue.svg?colorA=24292e&colorB=0366d6&style=flat&longCache=true&logo=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAA4AAAAOCAYAAAAfSC3RAAAABHNCSVQICAgIfAhkiAAAAAlwSFlzAAAM6wAADOsB5dZE0gAAABl0RVh0U29mdHdhcmUAd3d3Lmlua3NjYXBlLm9yZ5vuPBoAAAERSURBVCiRhZG/SsMxFEZPfsVJ61jbxaF0cRQRcRJ9hlYn30IHN/+9iquDCOIsblIrOjqKgy5aKoJQj4O3EEtbPwhJbr6Te28CmdSKeqzeqr0YbfVIrTBKakvtOl5dtTkK+v4HfA9PEyBFCY9AGVgCBLaBp1jPAyfAJ/AAdIEG0dNAiyP7+K1qIfMdonZic6+WJoBJvQlvuwDqcXadUuqPA1NKAlexbRTAIMvMOCjTbMwl1LtI/6KWJ5Q6rT6Ht1MA58AX8Apcqqt5r2qhrgAXQC3CZ6i1+KMd9TRu3MvA3aH/fFPnBodb6oe6HM8+lYHrGdRXW8M9bMZtPXUji69lmf5Cmamq7quNLFZXD9Rq7v0Bpc1o/tp0fisAAAAASUVORK5CYII=)](https://github.com/CezaryTarnowski-TomTom/gha-inject-secrets-into-file)
[![Execute pre-commit](https://github.com/CezaryTarnowski-TomTom/gha-inject-secrets-into-file/actions/workflows/precommit.yml/badge.svg)](https://github.com/CezaryTarnowski-TomTom/gha-inject-secrets-into-file/actions/workflows/precommit.yml)
[![Integration Test](https://github.com/CezaryTarnowski-TomTom/gha-inject-secrets-into-file/actions/workflows/integration.yml/badge.svg)](https://github.com/CezaryTarnowski-TomTom/gha-inject-secrets-into-file/actions/workflows/integration.yml)

This is a GitHub Action to replace placeholders in files with values from secrets or key vaults.

## Usage

This action uses go text file templating and replace values in files with secrets taken from other step or repo secrets.


### Example workflow

```yaml
name: My Workflow
on: [push, pull_request]
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@master
    - name: Run action
      uses: CezaryTarnowski-TomTom/gha-inject-secrets-into-file@v1
      with:
        secrets: ${{ toJson(secrets) }}
    - run: |
        echo .env
```

### Inputs

| Input                 | Description                                                                                                         |
|-----------------------|---------------------------------------------------------------------------------------------------------------------|
| `secrets`             | A JSON with secrets to use as replacement                                                                           |
| `file` _(optional)_   | Name of the input file a go text template to have secrets replaced with values form _secrets_ JSON (default _.env_) |
| `output` _(optional)_ | The file name of the output by default it would be the same as input file - it would get overwritten                |

> The file to be processed is using the [go text template](https://pkg.go.dev/text/template#hdr-Text_and_spaces)

### Outputs

No direct output apart from file with replaced placeholders

## Examples

### Using the action with Azure Key Vault

This is how to use the action with Azure Key Vault.

```yaml
steps:
  - uses: Azure/login@v1
    with:
      creds: ${{ secrets.AZURE_CREDENTIALS }}
  - uses: Azure/get-keyvault-secrets@v1
    with:
      keyvault: "MyKeyVault"
      secrets: '*'
    id: kv
  - uses: CezaryTarnowski-TomTom/gha-inject-secrets-into-file@v1
    with:
      secrets: ${{ toJson(steps.kv.outputs) }}
```

### Example .env file
```
SOME_VALUE={{ index . "my-secret" }}
OTHER_VALUE={{ .otherSecret }}
```

> NOTE: for variable names that contains dash/hyphen you need to use special syntax with index function. It is not possible to use {{ .name-with-hyphen }} as hyphen has a special meaning in the go template syntax.
