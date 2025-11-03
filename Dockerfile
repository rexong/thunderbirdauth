FROM golang:1.24-alpine AS builder
RUN apk add --no-cache gcc musl-dev sqlite

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY main.go ./
COPY server/ ./server/
COPY db/ ./db/

RUN go build -o thunderbirdauth

FROM alpine:latest
RUN apk add --no-cache sqlite
COPY --from=builder /app/thunderbirdauth /thunderbirdauth

RUN mkdir -p /db
ENV ENV=dev
ENV BASIC_USERNAME=home
ENV BASIC_PASSWORD=1234

CMD ["/thunderbirdauth"]

