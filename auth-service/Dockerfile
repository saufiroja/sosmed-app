FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o auth-service ./cmd

# Path: Dockerfile
FROM alpine:latest

RUN apk update && apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/auth-service .

EXPOSE 8080

CMD ["./auth-service"]