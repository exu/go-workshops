# The smallest final Go image 4.19MB
# :) Thanks to Tihomir :)

# build stage
FROM golang:alpine AS build-env
RUN mkdir /src
ADD main.go /src
RUN cd /src && CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static" -w -s' -o goapp

# final stage
FROM scratch
WORKDIR /app
COPY --from=build-env /src/goapp /app/
ENTRYPOINT ["/app/goapp"]
