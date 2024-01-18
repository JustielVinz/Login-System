# Use the official Golang base image
FROM golang:1.21

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

# Download and install any dependencies
RUN go mod download

# Build the Go application
RUN go build -o myapp

# Expose the port the application runs on
EXPOSE 3000

# Define the command to run your application
CMD ["./myapp"]
