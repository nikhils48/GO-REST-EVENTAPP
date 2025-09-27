# Stage 1: Build the Go binaries
FROM golang:1.25-alpine AS builder

# Install dependencies
RUN apk add --no-cache git

# Set workdir
WORKDIR /app

# Copy Go modules and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the API binary
RUN go build -o main ./cmd/api

# Build the migration binary
RUN go build -o migrate ./cmd/migrate

# Stage 2: Minimal runtime image
FROM alpine:latest

# Install SQLite if needed
RUN apk add --no-cache sqlite

# Set workdir
WORKDIR /app

# Copy the compiled binaries from builder
COPY --from=builder /app/main .
COPY --from=builder /app/migrate .

# Copy migration files
COPY --from=builder /app/cmd/migrate/migrations ./cmd/migrate/migrations

# Expose port
EXPOSE 8081

# Set environment variables if needed
ENV GIN_MODE=release

# Run migrations first, then start the server
CMD ./migrate up && ./main
