# Pre-Commit hook

### Install

`pip install --user pre-commit`
or
`curl https://pre-commit.com/install-local.py | python -`

### install hooks

`pre-commit install` if there is `.pre-commit-config.yaml` file in side the repo.

### run hooks

`pre-commit run` or just give a git commit. It will run precommit only on staged files. For unstaged one???

option `--all-files` for running all files.

### pre-commit config yaml

ref: [.pre-commit-config.yaml](Git_Precommit_config.yaml)
