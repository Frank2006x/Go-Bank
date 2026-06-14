FROM golang:1.26.4-alpine3.24 AS builder
WORKDIR /app
COPY . .
RUN go build -o main ./cmd/api/main.go

FROM alpine:3.24
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .

EXPOSE 8080

CMD ["/app/main"]
