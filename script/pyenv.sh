#!/usr/bin/env bash
# -*- coding: utf-8 -*-

[ -r "script/bootstrap.sh" ] && source "script/bootstrap.sh"

# load source files externals
if [ -e "$HOME/.pyenv" ]; then
    eval "$(pyenv init -)"
    eval "$(pyenv virtualenv-init -)"
fi

cd "${ROOT_DIR}"

pyenv versions | grep -q "${PYENV_NAME}"
response=$?

if [[ ! ${response} -eq '0' ]]; then
    pyenv virtualenv "${PYTHON_VERSION}" "${PYENV_NAME}" >> /dev/null 2>&1
fi

pyenv activate "${PROJECT_NAME}" >> /dev/null 2>&1