# CookieCutter

## Requirement

- ~/.cookiecutterrc , check [example](#rc-example)

## How to use

`cookiecutter [-v] (template folder|git repo)`

## Template

The template folder/repository should contain `cookiecutter.json`file.
Also check [jinja template desigin engine](https://jinja.palletsprojects.com/en/2.11.x/templates/) for more detial template expression (function, condition, ...).

#### cookiecutter.json

template project config

```jsonc
{
  "project_name": "My New Project",
  "project_slug": "{{ cookiecutter.project_name }}",
  "_copy_without_render": ["*.html"], // placeholder won't be replaced
  "license": ["MIT", "BSD-3", "GNU GPL v3.0", "Apache Software License 2.0"]
  // prompt for selecting one in the list
}
```

#### Basic template variable

`{{cookiecutter.<variable}}`

## .cookiecutterrc

```yaml
default_context: # value for inject
  aws_account_id: 1211111
abbreviations:
  gn: https://github.com/{0}.git
cookiecutters_dir: "e:\path\to\template\dir"
replay_dir: "e:\path\to\replay\dir"
```

## hooks

- hooks/pre_gen_project.py
- hooks/post_gen_project.py

## replay

Locate at `~/.cookiecutter_replay`. Stored used template json file.

Using option `--replay` will regenerate the same project.

## options

- `-o`, where to output.
- `-c`, checkout after clone if it is git repo
- `-f`, overwrite if exists
- `-s`, skip if exists
- `--config-file`, point to specific config file.

### issues: subprocess: CalledProcessError: for git clone

1. `git config core.longpaths=true`
2. Naming Files, Paths, and Namespaces for windows 10
3. Windows 10 "Enable NTFS long paths policy" option missing
