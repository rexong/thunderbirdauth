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
ENV PORT=8080
ENV BASIC_USERNAME=home
ENV BASIC_PASSWORD=1234
ENV DB_PATH=db/app.db
ENV LDAP_SERVER_IP_ADDRESS=0.0.0.0
ENV LDAP_SERVER_PORT=10389
ENV LDAP_STORE_PATH=db/ldap-data
ENV ADMIN_DN="cn=admin,dc=example,dc=com"
ENV ADMIN_PASSWORD=adminPassword

CMD ["/thunderbirdauth"]

