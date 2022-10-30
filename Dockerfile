# Build stage
FROM golang:latest-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine:3.16

RUN apk --no-cache add ca-certificates

WORKDIR /app
COPY --from=builder /app/main .
COPY cfg.env .
COPY db/migration ./db/migration

EXPOSE 8080
ENV CGO_ENABLED=0

ENTRYPOINT ["/app/main" ]