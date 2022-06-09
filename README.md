# update-terraform-lockfile

A GitHub Action which updates the Terraform lockfile

[![.github/workflows/gotest.yml](https://github.com/champ-oss/update-terraform-lockfile/actions/workflows/gotest.yml/badge.svg?branch=main)](https://github.com/champ-oss/update-terraform-lockfile/actions/workflows/gotest.yml)
[![.github/workflows/golint.yml](https://github.com/champ-oss/update-terraform-lockfile/actions/workflows/golint.yml/badge.svg?branch=main)](https://github.com/champ-oss/update-terraform-lockfile/actions/workflows/golint.yml)
[![.github/workflows/release.yml](https://github.com/champ-oss/update-terraform-lockfile/actions/workflows/release.yml/badge.svg)](https://github.com/champ-oss/update-terraform-lockfile/actions/workflows/release.yml)
[![.github/workflows/sonar.yml](https://github.com/champ-oss/update-terraform-lockfile/actions/workflows/sonar.yml/badge.svg)](https://github.com/champ-oss/update-terraform-lockfile/actions/workflows/sonar.yml)

[![SonarCloud](https://sonarcloud.io/images/project_badges/sonarcloud-black.svg)](https://sonarcloud.io/summary/new_code?id=champ-oss_update-terraform-lockfile)

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=champ-oss_update-terraform-lockfile&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=champ-oss_update-terraform-lockfile)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=champ-oss_update-terraform-lockfile&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=champ-oss_update-terraform-lockfile)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=champ-oss_update-terraform-lockfile&metric=reliability_rating)](https://sonarcloud.io/summary/new_code?id=champ-oss_update-terraform-lockfile)

## Features
- Updates the Terraform lockfile automatically
- Opens a pull request if updates are available

## Example Usage

```yaml
jobs:
  run:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
        with:
          token: ${{ secrets.GITHUB_TOKEN }}

      - uses: hashicorp/setup-terraform@v1.3.2
        with:
          terraform_version: 1.1.4
          terraform_wrapper: false

      - uses: champ-oss/update-terraform-lockfile
        with:
          token: ${{ secrets.GITHUB_TOKEN }}
```

## Token
By default the `GITHUB_TOKEN` should be passed to the `actions/checkout` step as well as this action (see example usage). This is necessary for the action to be allowed to push changes to a branch as well as open a pull request.


## Parameters
| Parameter | Required | Description |
| --- | --- | --- |
| token | false | GitHub Token or PAT |
| target-branch | false | Target branch for pull request |
| pull-request-branch | false | Branch to push changes |
| user | false | Git username |
| email | false | Git email |
| commit-message | false | Commit message to use |

## Contributing

