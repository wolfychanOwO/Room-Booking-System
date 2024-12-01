FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN go build -o booking-service ./cmd/booking-svc

RUN ls -la /app/

FROM alpine:latest

COPY --from=builder /app/booking-service /usr/local/bin/booking-service

ENV DATABASE_URL=postgres://user:password@postgres-booking:5432/booking_db?sslmode=disable

CMD ["booking-service"]
