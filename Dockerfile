FROM golang:1.26 AS builder
# FROM golang:1.26
WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY ./cmd ./cmd
COPY ./internal ./internal

# FROM alpine:3.1
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o /usr/local/bin/app ./cmd/sisima

FROM alpine:3.1
COPY --from=builder /usr/local/bin/app /app
ENV \
    SERVER_HOST=127.0.0.1 \
    SERVER_PORT=8888 \
    POSTGRES_HOST=127.0.0.1 \
    POSTGRES_PORT=5432 \
    POSTGRES_USER=sisima_user \
    POSTGRES_PASSWORD=sisima_password \
    POSTGRES_DATABASE=sisima

CMD ["./app", "&&", "ls", "-la"]
# CMD ["env"]
