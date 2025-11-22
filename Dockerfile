# ==== Build stage ====
FROM golang:1.25 AS builder

WORKDIR /app

# Download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source
COPY . .

# Build the app (static binary)
RUN CGO_ENABLED=1 GOOS=linux go build -o app ./cmd/app/main.go

# ==== Run stage ====
FROM debian:stable-slim

WORKDIR /app

# SQLite dependencies
RUN apt-get update && apt-get install -y sqlite3 && apt-get clean

# Copy binary + migrations
COPY --from=builder /app/app .
COPY ./scripts/migrations ./scripts/migrations

# Default port (change if needed)
EXPOSE 3000

CMD ["./app"]
