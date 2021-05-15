package sql_connector

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"                  // use GORM as ORM
	_ "github.com/jinzhu/gorm/dialects/mysql" //dialect for mysql
)

// OpenConnection => open connection to mysql db
func OpenConnection(mysqlUri string) *gorm.DB {
	log.Print("starting GORM mysql connection ...")

	db, err := gorm.Open("mysql", mysqlUri)

	if err != nil {
		log.Print(err)
		panic("failed to initialize connection to mysql database")
	}

	db.DB().SetMaxIdleConns(32)
	db.DB().SetMaxIdleConns(64)
	db.DB().SetConnMaxLifetime(30 * time.Second)
	db.LogMode(true)

	return db
}
