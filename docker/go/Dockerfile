# build stage
FROM golang:alpine AS build-env

ENV BUILD_DIR=/go/src/
ADD . $BUILD_DIR
WORKDIR $BUILD_DIR
RUN go build -o app

FROM alpine
COPY --from=build-env /go/src/app /app
ENTRYPOINT ["/app"]
