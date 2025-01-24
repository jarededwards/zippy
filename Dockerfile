FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY go.mod .

RUN go mod download

COPY . .

RUN go build -o zippy .

FROM alpine:3.18.2

LABEL org.opencontainers.image.source=https://github.com/jarededwards/zippy
LABEL org.opencontainers.image.description="simple golang http server"


COPY --from=builder /build/zippy /

EXPOSE 8080

ENTRYPOINT ["/zippy"]
