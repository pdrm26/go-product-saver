FROM golang:1.23-alpine AS builder

WORKDIR /app

ARG BINARY_NAME=go-post-saver

COPY go.mod go.sum ./
# RUN go get

# Download dependencies
# This step is cached if go.mod and go.sum don't change
RUN go mod download


# Copy the rest of the source code
COPY . .

# Build the application
RUN go build -o /${BINARY_NAME}


# Use a minimal alpine image for the final stage
FROM alpine:3.20

ARG BINARY_NAME


COPY --from=builder /${BINARY_NAME} /${BINARY_NAME}


EXPOSE 8080

CMD [ "/${BINARY_NAME}" ]