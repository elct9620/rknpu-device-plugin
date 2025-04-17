FROM docker.io/golang:1.23-alpine3.21 AS builder

RUN mkdir -p /go/src/github.com/elct9620/rknpu-device-plugin
ADD . /go/src/github.com/elct9620/rknpu-device-plugin
WORKDIR /go/src/github.com/elct9620/rknpu-device-plugin
RUN go install

FROM alpine:3.21
WORKDIR /root/
COPY --from=0 /go/bin/rknpu-device-plugin .
CMD ["./rknpu-device-plugin"]
