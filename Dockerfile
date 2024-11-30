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
FROM gcr.io/distroless/base-debian11 AS build-release-stage

# Setup working directory
WORKDIR /app

# Copy the static executable
COPY --from=build-stage /dist/main /dist/main

# Use user nonroot
USER nonroot:nonroot

# Run the static executable
CMD [ "/dist/main", "-env", "docker" ]