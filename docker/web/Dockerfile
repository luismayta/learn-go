FROM golang:1.8.0

MAINTAINER Luis Mayta <@slovacus>

COPY ./glide.yaml /go/src/golang-examples/
COPY ./glide.lock /go/src/golang-examples/

# set working directory to /app/
WORKDIR /go/src/golang-examples/

RUN apt-get update \
  && apt-get install -y zsh \
  && go get -u github.com/alecthomas/gometalinter \
  && gometalinter --install \
  && go get -u github.com/nsf/gocode \
  && go get -u golang.org/x/tools/cmd/goimports \
  && go get -u github.com/dougm/goflymake \
  && go get -u github.com/pengwynn/flint \
  && go get -u github.com/rogpeppe/godef \
  && go get -u github.com/golang/lint \
  && curl https://glide.sh/get | sh \
  && glide install
