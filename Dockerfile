# Use the official Golang image as the base image
FROM golang:1.20.4-alpine3.17

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY ../go.mod ./
COPY ../go.sum ./
COPY ../body.json ./

# Download the Go module dependencies
RUN go mod download

# Copy the source code into the container
COPY /main.go /app/

# Build the Go application
RUN go build -o main .

# Expose the port the application will run on
EXPOSE 8080

# Set the entry point command for the container
CMD ["./main"]