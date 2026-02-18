
FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o sisima ./cmd/sisima

FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache tzdata

ENV TZ=Asia/Jakarta

COPY --from=builder /app/sisima .

EXPOSE 8888

CMD ["./sisima"]

