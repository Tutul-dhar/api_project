# STAGE 1: Build the Go binary

FROM golang:1.22 AS builder

WORKDIR /app

# Copy go.mod and go.sum first (for caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source code
COPY . .

# Build the Go app binary
RUN go build -o main .


# STAGE 2: Run the binary in minimal image

FROM debian:bookworm-slim

# Set working directory in the smaller image
WORKDIR /app

# Copy only the binary from the builder stage
COPY --from=builder /app/main .

# Expose the port (used in flag default or overridden)
EXPOSE 8080

# Run the binary
CMD ["./main"]
