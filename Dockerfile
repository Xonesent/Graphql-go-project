FROM golang:latest AS builder

WORKDIR /app

COPY . .

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o main ./cmd/main.go

FROM alpine:latest

WORKDIR /app

RUN apk update

COPY --from=builder /app/main /app
COPY --from=builder /app/.env /app
COPY --from=builder /app/migrations /app/migrations

ENV PATH="/app:${PATH}"

CMD ["main"]