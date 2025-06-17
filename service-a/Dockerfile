# Ganti versi dari golang:1.21-alpine → golang:1.22.3-alpine
FROM golang:1.22.3-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o main .

CMD ["./main"]