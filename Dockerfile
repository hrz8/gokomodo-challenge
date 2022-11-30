# -- build stage --
FROM golang:1.19.3-alpine3.16 AS build

# setup
RUN apk update && apk add git --no-cache 
RUN apk add build-base

# set working directory
WORKDIR /app

# install dependencies
ADD go.mod go.sum ./
RUN go mod download

# build source
ADD cmd ./cmd
ADD internal ./internal
ADD pkg ./pkg
RUN CGO_ENABLED=1 GOOS=linux go build -a -ldflags '-linkmode external -extldflags "-static"' -o bin/server cmd/server/main.go

# -- run stage --
FROM alpine:3.16.2

# setup
RUN apk update && apk add ca-certificates --no-cache

# set working directory
WORKDIR /app

# copy app
COPY --from=build /app/bin/server ./server

EXPOSE 3000

# container run command
CMD ["./server"]
