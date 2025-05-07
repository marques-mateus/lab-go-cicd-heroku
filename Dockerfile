FROM golang:1.24-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download 

COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@latest && \
    swag init

RUN go build -o main

FROM alpine:latest
WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/docs ./docs

EXPOSE 8080


CMD ["./main"]