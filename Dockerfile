FROM golang:1.26.4-alpine3.24 AS builder

WORKDIR /app

COPY . .

RUN go build -o main ./cmd/api/main.go

RUN apk add --no-cache curl tar

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.19.1/migrate.linux-amd64.tar.gz \
    | tar xvz && \
    mv migrate /migrate

FROM alpine:3.24

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /migrate ./migrate
COPY db/migration ./migration
COPY app.env .
COPY start.sh .
RUN chmod +x /app/start.sh
EXPOSE 8080

CMD ["/app/main"]
ENTRYPOINT ["/app/start.sh"]