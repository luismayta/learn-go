FROM golang:1.8.0

MAINTAINER Luis Mayta <@slovacus>

COPY ./glide.yaml /go/src/golang-examples/
COPY ./glide.lock /go/src/golang-examples/

# set working directory to /app/
WORKDIR /go/src/golang-examples/

RUN apt-get update \
  && apt-get install -y zsh \
  && curl https://glide.sh/get | sh \
  && glide install
