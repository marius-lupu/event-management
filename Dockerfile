# Build the event-management binary
# https://hub.docker.com/_/golang?tab=tags
FROM golang:1.18 as builder

# Add Maintainer Info
LABEL maintainer="Marius Lupu"

# Set the Current Working Directory inside the container
WORKDIR /workspace

# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum

# Download all dependencies
RUN go mod download

# Copy the go source
COPY *.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o event-management .

# Use distroless as minimal base image to package the event-management binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/event-management .
USER 65532:65532

ENTRYPOINT ["/event-management"]