# --- Stage 1: Build the App ---
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Download dependencies first (Cached layer)
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the application
# -o main: output file name
# ./cmd/api: entry point
RUN go build -o main ./cmd/api

# --- Stage 2: Run the App ---
FROM alpine:latest

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Expose the port (Documentation only, standard for Go apps)
EXPOSE 8080

# Run the binary
CMD ["./main"]