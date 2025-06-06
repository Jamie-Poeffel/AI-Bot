# Build stage
FROM golang:1.23.5-alpine AS builder

WORKDIR /app

COPY .env .env

# Copy Go modules and download deps early (for caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the binary inside /app
RUN go build -o server .

# Run stage
FROM alpine:latest

# Install CA certificates (for HTTPS calls)
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/server .

# Expose the port your app listens on
EXPOSE 4000

# Run the server
CMD ["./server"]
