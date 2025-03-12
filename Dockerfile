FROM golang:1.23.4-alpine3.19 AS builder

WORKDIR /app

ENV CGO_ENABLED=0
ENV GOPATH=/go
ENV GOCACHE=/go-build

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod/cache \
    go mod download

COPY . .
RUN --mount=type=cache,target=/go/pkg/mod/cache \
    --mount=type=cache,target=/go-build \
    go build -o main cmd/main.go

FROM alpine:3.19

RUN apk add --no-cache ca-certificates

WORKDIR /app
COPY --from=builder /app/main .
COPY .env .

EXPOSE 8080
CMD ["/app/main"]
