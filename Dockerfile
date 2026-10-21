FROM docker.io/golang:1.27-alpine3.24 AS builder

RUN mkdir -p /go/src/github.com/elct9620/rknpu-device-plugin
ADD . /go/src/github.com/elct9620/rknpu-device-plugin
WORKDIR /go/src/github.com/elct9620/rknpu-device-plugin
RUN go install

FROM alpine:3.24
# Upgraded explicitly so the trust store comes from Alpine security updates
RUN apk upgrade --no-cache ca-certificates-bundle
WORKDIR /root/
COPY --from=0 /go/bin/rknpu-device-plugin .
CMD ["./rknpu-device-plugin"]
