#!/usr/bin/env bash
# -*- coding: utf-8 -*-

[ -r "script/bootstrap.sh" ] && source "script/bootstrap.sh"

rvm gemset create "${RVM_NAME}" >> /dev/null 2>&1
rvm use "${RVM_VERSION}"@"${RVM_NAME}" >> /dev/null 2>&1
