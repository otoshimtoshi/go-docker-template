package infrastructure

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf(".envが読み込み出来ませんでした: %v", err)
	}
}
