FROM golang:alpine as builder
WORKDIR /app
COPY . .
EXPOSE 7009
ENTRYPOINT ["go", "run", "main.go"]