package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"time"
)

func DB() *gorm.DB {
	sConn := "host=localhost port=5432 user=postgres dbname=kp_express password=postgres sslmode=disable"
	db, err := gorm.Open("postgres", sConn)
	if err != nil {
		log.Fatal(err)
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(10)
	db.DB().SetConnMaxLifetime(time.Hour)
	db.LogMode(true)
	return db
}
