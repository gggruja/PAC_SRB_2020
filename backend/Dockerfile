# Golang base image used to build the application in the first phase
FROM golang:alpine AS builder

# Maintainer
LABEL maintainer="Goran Grujic <goran.grujic@prodyna.com>"

# Environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Setting current work directory inside the container
WORKDIR /build

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Build the Go application
RUN go build -o main .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/main .

# Build a small image
FROM scratch

# Copy the Pre-built binary file from the previous stage. Also copy the .env file
COPY --from=builder /dist/main .
COPY --from=builder /build/.env .

# Command to run
ENTRYPOINT ["/main"]
