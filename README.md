# go-docker-template

## Library

- Web フレームワーク：[gin-gonic](https://github.com/gin-gonic/gin)
- ORM：[go-gorm](https://github.com/go-gorm/gorm)
- ホットリロード：[air-verse/air](https://github.com/air-verse/air)
- デバッグ：[go-delve/delve](https://github.com/go-delve/delve)

## Docker

```sh
# バックグラウンド実行の場合
% docker compose up -d
# ログ出力あり実行の場合
% docker compose up
```

## Go Modules

```sh
% go mod tidy
```

## Database Migration

```sh
% docker exec -it go-api sh
% migrate -database 'mysql://root:password@tcp(go-db:3306)/sample' -path migrations up
% migrate -database 'mysql://root:password@tcp(go-db:3306)/sample' -path migrations down
% go run seeds/seeder.go
```

## Linting

```sh
% golangci-lint run
% golangci-lint run --fix
```
