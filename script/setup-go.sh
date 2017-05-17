#!/usr/bin/env bash
# -*- coding: utf-8 -*-

# shellcheck source=/dev/null
[ -r "script/bootstrap.sh" ] && source "script/bootstrap.sh"

go get -u github.com/alecthomas/gometalinter
gometalinter --install
go get -u github.com/nsf/gocode
go get -u golang.org/x/tools/cmd/goimports
go get -u github.com/dougm/goflymake
go get -u github.com/pengwynn/flint
go get -u github.com/rogpeppe/godef
go get -u github.com/golang/lint
