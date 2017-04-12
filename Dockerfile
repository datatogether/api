FROM golang:1.8.1-alpine

# Copy the local package files to the container’s workspace.
ADD . /go/src/github.com/archivers-space/archivers-api

# Install api binary globally within container 
RUN go install github.com/archivers-space/archivers-api

# Set binary as entrypoint
ENTRYPOINT /go/bin/archivers-api

# Expose default port
EXPOSE 8080
