FROM golang:1.22-alpine3.20 as build
WORKDIR /build

COPY . .
RUN go mod download
RUN GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o service ./cmd/service/main.go

FROM alpine:3.20 as app
WORKDIR /app

COPY --from=build /build/service .

RUN apk --no-cache add curl
HEALTHCHECK \
    --start-period=10s \
    --start-interval=1s \
    --interval=1m \
    --timeout=20s \
    --retries=3 \
    CMD curl -f http://localhost:8080/health || exit 1

CMD ./service
