FROM golang:1.22-alpine3.20 as build
WORKDIR /build

COPY . .
RUN go mod download
RUN GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o service ./cmd/service/main.go

FROM alpine:3.20 as app
WORKDIR /app

COPY --from=build /build/service .

CMD ./service
