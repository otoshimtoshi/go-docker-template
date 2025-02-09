#　v1.23.5
FROM golang:1.23-alpine

WORKDIR /app

# v1.61.7
RUN go install github.com/air-verse/air@latest
# v1.24.0
RUN go install github.com/go-delve/delve/cmd/dlv@latest

COPY go.mod go.sum ./
RUN go mod download

# v1.10.0
RUN go get github.com/gin-gonic/gin \
# v1.9.16
    github.com/jinzhu/gorm \
    github.com/jinzhu/gorm/dialects/mysql \
# v1.5.1
    github.com/joho/godotenv

# v4.18.2
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

CMD ["air", "-c", ".air.toml"]