FROM golang:alpine AS builder
WORKDIR /app
COPY . .
EXPOSE 7008
ENTRYPOINT ["go", "run", "main.go"]