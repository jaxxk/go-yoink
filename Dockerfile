# Use the official Golang image as a base image
FROM golang:1.22

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Install goose
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Expose port 8080 to the outside world
EXPOSE 8080

# Make the entrypoint script executable
RUN chmod +x /app/entrypoint.sh

# Use the entrypoint script to run migrations and then the tests
ENTRYPOINT ["/app/entrypoint.sh"]
