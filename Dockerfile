# Use the official Go image as the base image
FROM golang:1.20

# Set the working directory in the container
WORKDIR /app

# Copy the application files into the working directory
COPY . /app

# Download the application dependencies
RUN go mod download

# Install the package
RUN go install -v ./...

# Installing postgresql

RUN apt-get update && apt-get install -y postgresql-client

# Build the application
RUN go build -o main .

# Expose port 8080
EXPOSE 8080

# Define the entry point for the container
CMD ["./main"]