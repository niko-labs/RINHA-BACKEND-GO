######################
#    dependencies    #
######################

FROM golang:1.22.0-alpine AS dependencies

RUN apk update && apk add --no-cache gcc libc-dev make

# Set the Current Working Directory inside the container
WORKDIR /.dependencies

# Copy Only the go mod and sum files to take advantage of caching
COPY ./app/go.mod .
COPY ./app/go.sum .

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download


#####################
#      builder      #
#####################

# Build the Go from dependencies
FROM dependencies AS builder
WORKDIR /.builder

# Copy discarding the go.mod and go.sum files
COPY ./app .

# Build the Go app
ENV CGO_ENABLED=0

# Build the Go app
RUN make build-server



#####################
#       FINAL       #
#####################

FROM alpine:latest AS final

# Set the Current Working Directory inside the container
WORKDIR /.app
RUN apk update && apk add --no-cache curl
ENV GIN_MODE=release

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /.builder/bin/server /.app/server.so

# Command to run the executable
CMD ["/.app/server.so"]
