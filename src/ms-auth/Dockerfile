FROM golang:alpine as builder
WORKDIR /app
COPY . .
EXPOSE 8040
ENTRYPOINT ["go", "run", "main.go"]