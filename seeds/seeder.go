package main

import (
	"log"

	"go-docker-template/app/infrastructure"
	"go-docker-template/seeds/seeds"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	infrastructure.LoadEnv()
	c := infrastructure.NewConfig()

	dbConn := openConnection(c)
	defer dbConn.Close()

	for _, seed := range seeds.All() {
		if err := seed.Run(dbConn); err != nil {
			log.Printf("Running seed '%s', failed with error: %s", seed.Name, err)
		}
	}
}

//nolint:lll
func openConnection(c *infrastructure.Config) *gorm.DB {
	db, err := gorm.Open(
		"mysql",
		c.DB.Main.Username+":"+c.DB.Main.Password+"@tcp("+c.DB.Main.Host+")/"+c.DB.Main.DBName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Printf("Couldn't establish database connection: %s", err)
	}

	return db
}
