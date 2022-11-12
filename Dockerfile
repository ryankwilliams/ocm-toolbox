FROM golang:1.19
ARG TMP_DIR=/tmp/ocm-toolbox
WORKDIR ${TMP_DIR}
ADD . ${TMP_DIR}
RUN go build -o /usr/local/bin/ocm-toolbox cmd/main.go
RUN rm -rf ${TMP_DIR}
ENTRYPOINT [ "ocm-toolbox" ]