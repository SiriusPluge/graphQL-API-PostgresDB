# STEP 1: Build executable binary
FROM golang:latest AS builder
# Copy project into image
COPY . $GOPATH/src/graphQL-API-PostgresDB
# Set working directory to /go-graphql-api which contains main.go
WORKDIR $GOPATH/src/graphQL-API-PostgresDB

# STEP 2: Build a small image start from scratch
FROM scratch
# Expose port 4000
EXPOSE 8080