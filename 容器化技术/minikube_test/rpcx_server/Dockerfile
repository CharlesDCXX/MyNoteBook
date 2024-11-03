# Dockerfile
FROM golang:1.16-buster AS builder
RUN mkdir /src
ADD . /src
WORKDIR /src
ENV GOPROXY goproxy.cn
RUN go env -w GO111MODULE=auto
RUN go build -o main .

FROM debian:buster-slim

WORKDIR /

COPY --from=builder /src/main /main
EXPOSE 3000
ENTRYPOINT ["/main"]