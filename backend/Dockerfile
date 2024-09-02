# Use the official Golang image as the base image
FROM golang:1.23-alpine

# Set the working directory
WORKDIR /app

# Copy the Go modules files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the application code
COPY . .

# Build the Go application
RUN go build -o main .

# Expose port
EXPOSE 8080

# Run the executable
CMD ["./main"]
