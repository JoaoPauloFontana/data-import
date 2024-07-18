FROM golang:1.18-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o api ./cmd

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/api .

EXPOSE 8080

CMD ["./api"]
