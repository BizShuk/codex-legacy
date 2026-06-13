#!/bin/bash

### Installation of pre-commit cli
pip install pre-commit

### Create .pre-commit-config.yaml with sample
pre-commit sample-config

### Install Hooks
pre-commit install

### Run hook cross all project
pre-commit run --all-files