FROM golang:1.23-alpine3.20 AS build

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o migrator ./cmd/migrator/

FROM alpine:3.20 AS app

RUN addgroup -S app-group && adduser -S app-user -G app-group

WORKDIR /app
COPY --from=build /build/migrator ./

RUN chown -R app-user:app-group /app

RUN apk --no-cache add curl

USER app-user

CMD ./migrator
