# Use the official Golang image to create a build artifact.
FROM golang:1.22 as builder

# Set the working directory outside $GOPATH to enable the support for modules.
WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the application, disable cgo to create a statically linked binary.
RUN CGO_ENABLED=0 GOOS=linux go build -o /to-do ./cmd

# Use scratch (empty) container image
FROM scratch

# Copy the binary to the production image from the builder stage.
COPY --from=builder /to-do /to-do

EXPOSE 8080

COPY .env.sample .env

# Run the web service on container startup.
CMD ["/to-do"]
