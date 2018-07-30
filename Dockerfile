# build stage
FROM golang:alpine AS build-env

ENV BUILD_DIR=/go/src/github.com/exu/go-workshops/
RUN mkdir -p $BUILD_DIR
ADD . $BUILD_DIR
WORKDIR $BUILD_DIR

ENTRYPOINT go
