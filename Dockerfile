# Use an official Golang image as the base image
FROM golang:1.20 AS builder

# Set the current working directory inside the container
WORKDIR /go/src/app

# Copy the current directory contents into the container
COPY . .

# Install dependencies (if any) and build the Go app
RUN go mod tidy
RUN go build -o /go/bin/app

# Use a smaller base image for the final image (alpine)
FROM alpine:latest

# Install necessary dependencies in the smaller image (Alpine)
RUN apk --no-cache add ca-certificates

# Set the working directory to the location of the built binary
WORKDIR /root/

# Copy the Go binary from the builder stage
COPY --from=builder /go/bin/app /root/

# Set the entry point to run the Go app
ENTRYPOINT ["./app"]

# Expose the port if the Go app is a web server
EXPOSE 8080

# Optionally, run the Go app (uncomment if you want to automatically run it)
# CMD ["./app"]
