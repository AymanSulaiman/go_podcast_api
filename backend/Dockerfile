ARG APP_NAME=main

# Build stage
FROM golang:1.21.4 as build
ARG APP_NAME
ENV APP_NAME=$APP_NAME
WORKDIR /app
COPY . ./app
RUN go mod download
RUN go build -o /$APP_NAME

# Production stage
FROM alpine:latest as production
ARG APP_NAME
ENV APP_NAME=$APP_NAME
WORKDIR /root/
COPY --from=build /$APP_NAME ./
CMD ./$APP_NAME




# FROM golang:1.21.4-alpine

# WORKDIR /app

# COPY . .

# RUN go mod download && go mod verify

# EXPOSE 8000

# CMD ["go","run", "main.go"]
