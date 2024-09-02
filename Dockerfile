# Stage 1: Build the Go app
FROM golang:1.20-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main ./api
# Stage 2: A minimal docker image with only the necessary dependencies
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .
COPY --from=builder /app/api/helper/errorcode.json ./api/helper/
COPY --from=builder /app/api/config/config.uat.json ./api/config/

# Set environment variable to distinguish between environments
ENV APP_ENV=uat
ENV ERROR_PATH="/root/api/helper/errorcode.json"
ENV CONFIG_PATH="/root/api/config/config.uat.json"

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
