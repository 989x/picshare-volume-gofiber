# Start from the official Golang base image
FROM golang:1.22.4-alpine

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./
COPY go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN go build -o main ./cmd/main.go

# Expose port
EXPOSE 8080

# Run the binary
CMD ["./main"]
