# Start from the official Golang image
FROM golang:1.23-alpine

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application files
COPY . .

# Build the application
RUN go build -o main .

# Expose the application's port
EXPOSE 8080

# Run the application
CMD ["./main"]
