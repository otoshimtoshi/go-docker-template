#　v1.23.5
FROM golang:1.23-alpine

WORKDIR /app

# v1.61.7
RUN go install github.com/air-verse/air@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest

COPY go.mod go.sum ./
RUN go mod download

RUN go get github.com/gin-gonic/gin \
    github.com/jinzhu/gorm \
    github.com/jinzhu/gorm/dialects/mysql \
    github.com/joho/godotenv
    
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

CMD ["air", "-c", ".air.toml"]