FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download 

COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN go build -o notification-svc ./cmd/notification-svc

RUN ls -la /app/

FROM alpine:latest

COPY --from=builder /app/notification-svc /usr/local/bin/notification-svc

CMD ["notification-svc"]
