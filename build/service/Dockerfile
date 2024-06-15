FROM golang:1.22-alpine3.19 as build
WORKDIR /build

COPY . .
RUN go mod download
RUN GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o service ./cmd/service/main.go

FROM alpine:3.19 as app
WORKDIR /app

COPY --from=build /build/service .
COPY --from=build /build/cmd/service/config.yaml .

CMD ./service
