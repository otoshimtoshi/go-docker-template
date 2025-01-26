#　v1.23.5
FROM golang:latest

WORKDIR /app

# v1.61.7
RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

CMD ["air", "-c", ".air.toml"]