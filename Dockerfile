FROM docker.io/golang:1.21.6-alpine3.19 AS builder

RUN mkdir -p /go/src/github.com/elct9620/rknpu-device-plugin
ADD . /go/src/github.com/elct9620/rknpu-device-plugin
WORKDIR /go/src/github.com/elct9620/rknpu-device-plugin
RUN go install

FROM alpine:3.19.1
WORKDIR /root/
COPY --from=0 /go/bin/rknpu-device-plugin .
CMD ["./rknpu-device-plugin"]
