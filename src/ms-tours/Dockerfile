FROM golang:alpine AS builder
WORKDIR /app
COPY . .
EXPOSE 7007
ENTRYPOINT ["go", "run", "main.go"]