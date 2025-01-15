FROM golang:1.23-alpine3.20 AS build

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o service ./cmd/service/main.go

FROM alpine:3.21 AS app

RUN addgroup -S app-group && adduser -S app-user -G app-group

WORKDIR /app
COPY --from=build /build/service ./

RUN chown -R app-user:app-group /app

RUN apk --no-cache add curl

USER app-user

EXPOSE 8080

CMD ./service
