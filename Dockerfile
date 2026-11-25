# Build stage
FROM golang:1.24.5-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o api ./api

# Runtime stage
FROM scratch

WORKDIR /app

# Copy migrations from builder
COPY --from=builder /app/migrations ./migrations

# Copy built application from builder
COPY --from=builder /app/api .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Expose port
EXPOSE 8080
ENV MIGRATIONS_DIR=/app/migrations

# Run application
CMD ["./api"]
