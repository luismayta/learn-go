#!/usr/bin/env bash
# -*- coding: utf-8 -*-

PROJECT_NAME=go_examples
PYTHON_VERSION=2.7.9
PYENV_NAME="${PROJECT_NAME}"

GVM_NAME="${PROJECT_NAME}"
GVM_PATHS_NAME="{src, pkg, bin}"

RVM_NAME="${PROJECT_NAME}"
RVM_VERSION=2.1.0
# Vars Dir
ROOT_DIR="`pwd`"
RESOURCES_DIR="$ROOT_DIR/resources"
RESOURCES_DB_DIR="$RESOURCES_DIR/db"

source "${ROOT_DIR}/.env" >> /dev/null 2>&1
