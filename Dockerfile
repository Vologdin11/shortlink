FROM golang:alpine AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

COPY . /build

WORKDIR /build

RUN go build -ldflags='-w -s' -o shortlink ./cmd/shortlink

#image
FROM alpine:latest

COPY --from=builder /build/shortlink /

CMD ./shortlink