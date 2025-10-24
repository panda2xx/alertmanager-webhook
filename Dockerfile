# Stage 1: Build the Go application
FROM golang:1.25.3-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go module files
# This leverages Docker's layer caching
COPY go.* ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application as a static binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o alertmanager-webhook .

# ---

# Stage 2: Create the final, minimal image
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/alertmanager-webhook .

# Expose port 9095 to the outside world
EXPOSE 9095

# Command to run the application when the container starts
CMD ["./alertmanager-webhook"]
