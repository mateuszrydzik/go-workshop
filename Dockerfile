# Create build stage based on buster image
FROM golang:1.16-buster AS builder
COPY ./workshop /app
WORKDIR /app
ENTRYPOINT ["tail", "-f", "/dev/null"]