###################################
# STEP 1: build executable binary #
###################################
FROM golang:alpine AS build-stage

# Add maintainer info
LABEL maintainer="Rodericus Ifo Krista"

# Set build env
ENV GOOS=linux
ENV CGO_ENABLED=0

# Install linux packages
RUN apk update
RUN apk add --no-cache git bash
RUN apk add build-base make

# Setup working directory
WORKDIR /app

# Copy Go dependency files first
COPY go.mod go.sum ./

# Download and install all the dependencies
RUN go mod download
RUN go mod tidy
RUN go mod verify
RUN go install github.com/google/wire/cmd/wire@latest

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Build Application
RUN make build

###############################
# STEP 2: build a small image #
###############################
FROM alpine:latest AS build-release-stage

# Set build arg
ARG ENV
ARG PORT

# Validate build arg
RUN if [ -z "$ENV" ] || [ -z "$PORT" ]; then echo "ERROR: ENV and PORT must be provided. Use --build-arg ENV=dev --build-arg PORT=8080"; exit 1; fi

# Set build env
ENV ENV=${ENV}
ENV PORT=${PORT}

# Setup working directory
WORKDIR /app

# Copy the binary
COPY --from=build-stage /dist/main /dist/main

# Copy environment files
COPY env /app/env

# Add a non-root user
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

# Expose Port
EXPOSE ${PORT}

# Use shell form for CMD
CMD ["sh", "-c", "/dist/main -env \"$ENV\""]