FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download && \
    go build -o betcrafter-cli ./cmd/cli && \
    go build -o betcrafter-web ./cmd/web


FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/betcrafter-cli .
COPY --from=builder /app/betcrafter-web .
COPY .env .

CMD ["./betcrafter-cli"]