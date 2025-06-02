# -------- Stage 1: Build the Go binary --------
    FROM golang:1.24.3 AS builder

    WORKDIR /app
    
    # Copy go.mod and go.sum to cache dependencies
    COPY go.mod go.sum ./
    RUN go mod download
    
    # Copy the rest of your application
    COPY . .
    
    # Build a static binary for Alpine
    RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .
    
    # -------- Stage 2: Run using a minimal image --------
    FROM alpine:latest
    
    WORKDIR /root/
    
    # Copy binary from the builder stage
    COPY --from=builder /app/main .
    
    # Set the entry point
    CMD ["./main"]
    