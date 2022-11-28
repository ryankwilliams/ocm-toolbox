FROM golang:1.19 AS build
WORKDIR /tmp/src
COPY . /tmp/src
RUN CGO_ENABLED=0 go build -o /tmp/ocm-toolbox

FROM alpine:latest
COPY --from=build /tmp/ocm-toolbox /usr/local/bin/
ENTRYPOINT [ "ocm-toolbox" ]