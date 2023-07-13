# Stage 1: Build the Go binary
FROM golang:latest

WORKDIR /app

# Copy the rest of the application source code
COPY . /app

# Build the Go binary
RUN go build -o bookstore

# Set the entry point to start the application
ENTRYPOINT ["./bookstore"]
