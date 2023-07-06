# Use the official Go image as the base image
FROM golang:1.20.5

# Set the working directory in the container
WORKDIR /app

# Copy the application files into the working directory
COPY . /app

# Build the application
RUN go build -o main .

# Expose port 3000
EXPOSE 3000

# Define the entry point for the container
CMD ["./main"]