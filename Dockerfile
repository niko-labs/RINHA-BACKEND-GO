#####################
#      builder      #
#####################

FROM golang:1.22.0-alpine AS builder

RUN apk update && apk add --no-cache gcc libc-dev make

# Set the Current Working Directory inside the container
WORKDIR /_builder

# Copy all the source code to the Working Directory
COPY ./app .

# Build the Go app
RUN make build-server



#####################
#       FINAL       #
#####################

FROM alpine:3.12.0 AS final

# Set the Current Working Directory inside the container
WORKDIR /_app
RUN apk update && apk add --no-cache curl

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /_builder/bin/server /_app/_server

# Command to run the executable
CMD ["/_app/_server"]
