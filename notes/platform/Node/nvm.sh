#!/bin/bash

NODE_VER=""

### Install node
nvm install --lts
nvm install ${NODE_VER}
nvm uninstall ${NODE_VER}

### List out remote node versions
nvm ls-remote

### choose default
nvm use ${NODE_VER}