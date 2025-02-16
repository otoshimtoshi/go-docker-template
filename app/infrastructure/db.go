package infrastructure

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // MySQL dialect for gorm
)

type DB struct {
	Host       string
	Username   string
	Password   string
	DBName     string
	Connection *gorm.DB
}

func NewDB() *DB {
	c := NewConfig()

	return newDB(&DB{
		Host:       c.DB.Host,
		Username:   c.DB.Username,
		Password:   c.DB.Password,
		DBName:     c.DB.DBName,
		Connection: nil,
	})
}

//nolint:lll
func newDB(d *DB) *DB {
	db, err := gorm.Open("mysql", d.Username+":"+d.Password+"@tcp("+d.Host+")/"+d.DBName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}

	d.Connection = db

	return d
}

// Begin begins a transaction.
func (db *DB) Begin() *gorm.DB {
	return db.Connection.Begin()
}

func (db *DB) Connect() *gorm.DB {
	return db.Connection
}
