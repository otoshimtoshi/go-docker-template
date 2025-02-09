package infrastructure

import "os"

type Config struct {
	DB struct {
		Host     string
		Username string
		Password string
		DBName   string
	}
	Routing struct {
		Port string
	}
}

func NewConfig() *Config {
	c := new(Config)

	c.DB.Host = os.Getenv("DB_HOST")
	c.DB.Username = os.Getenv("DB_USERNAME")
	c.DB.Password = os.Getenv("DB_PASSWORD")
	c.DB.DBName = os.Getenv("DB_DBNAME")

	c.Routing.Port = ":" + os.Getenv("PORT")

	return c
}
