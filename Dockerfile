FROM golang:1.24-alpine AS builder
RUN apk add --no-cache gcc musl-dev sqlite

WORKDIR /app
COPY . .
RUN go build -o thunderbirdauth

FROM alpine:latest
COPY --from=builder /app/thunderbirdauth /thunderbirdauth

RUN mkdir -p /db
CMD ["/thunderbirdauth"]

