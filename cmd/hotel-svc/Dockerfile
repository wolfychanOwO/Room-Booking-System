FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN go build -o hotel-service ./cmd/hotel-svc

RUN ls -la /app/

FROM alpine:latest

COPY --from=builder /app/hotel-service /usr/local/bin/hotel-service

ENV DATABASE_URL=postgres://user:password@postgres-hotel:5432/hotel_db?sslmode=disable

CMD ["hotel-service"]
