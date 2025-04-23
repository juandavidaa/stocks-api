FROM golang:1.24-alpine AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /migrator cmd/migrate/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o /api      cmd/api/main.go

FROM alpine:latest
RUN apk add --no-cache ca-certificates
WORKDIR /app

COPY --from=builder /migrator .
COPY --from=builder /api .
COPY data.json ./seed/data.json

RUN chmod +x migrator && \
    chmod +x api

EXPOSE 8080