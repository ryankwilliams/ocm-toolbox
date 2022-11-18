FROM golang:1.19 AS build
WORKDIR /tmp/src
COPY . /tmp/src
RUN go build -o /tmp/ocm-toolbox cmd/main.go

FROM alpine:latest
COPY --from=build /tmp/ocm-toolbox /usr/local/bin/ocm-toolbox
ENTRYPOINT [ "ocm-toolbox" ]