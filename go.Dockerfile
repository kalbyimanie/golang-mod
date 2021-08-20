FROM golang:alpine3.13 AS builder
RUN apk update && \
    apk upgrade && \
    apk add bash \
            git
ENV GO111MODULE=on GOOS=linux GOARCH=amd64
WORKDIR /opt
COPY go.mod /opt/go.mod
COPY main.go /opt/main.go
RUN go mod vendor && \
    go build -o main main.go

FROM alpine:edge
WORKDIR /go
COPY --from=builder /opt/main /go/main
CMD [ "/bin/sh", "-c", "/go/main" ]