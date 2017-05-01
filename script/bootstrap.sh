#!/usr/bin/env bash
# -*- coding: utf-8 -*-

export PROJECT_NAME=golang-examples

export PYTHON_VERSION=2.7.9
export PYENV_NAME="${PROJECT_NAME}"

export GVM_NAME="${PROJECT_NAME}"
export GVM_PATHS_NAME=(
    "src"
    "pkg"
    "bin"
)

export GRIP_PORT=6430
export DEPLOY_ACCOUNT=ubuntu

# Vars Dir
export ROOT_DIR
ROOT_DIR=$(pwd)
export RESOURCES_DIR="$ROOT_DIR/resources"
export RESOURCES_DB_DIR="$RESOURCES_DIR/db"
export REQUIREMENTS_DIR="${ROOT_DIR}/requirements/"
export SOURCE_DIR="${ROOT_DIR}/"
export REQUIREMENTS_DIR="${SOURCE_DIR}/requirements/"
