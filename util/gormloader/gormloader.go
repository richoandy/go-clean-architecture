package gormloader

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //dialect for mysql
)

// OpenConnection => open connection to mysql db
func OpenConnection(mysqlURI string) (*gorm.DB, error) {
	log.Print("starting GORM mysql connection ...")

	db, err := gorm.Open("mysql", mysqlURI)

	if err != nil {
		return nil, err
	}

	db.DB().SetMaxIdleConns(32)
	db.DB().SetMaxIdleConns(64)
	db.DB().SetConnMaxLifetime(30 * time.Second)
	db.LogMode(true)

	return db, nil
}
