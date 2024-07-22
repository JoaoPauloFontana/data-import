FROM golang:1.18-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o importer ./cmd

FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache postgresql-client

COPY --from=builder /app/importer .
COPY ./Reconfile_fornecedores.xlsx ./

CMD ["./importer"]