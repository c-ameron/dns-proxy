FROM golang:1.12 as builder
WORKDIR /go/src/dns-proxy
COPY main.go .
RUN go build

FROM debian:stretch-slim
WORKDIR /app
RUN apt-get update && \
      apt-get install -y ca-certificates
COPY --from=builder /go/src/dns-proxy/dns-proxy ./
EXPOSE 8853
CMD ["./dns-proxy"]