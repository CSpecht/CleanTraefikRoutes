package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBA *gorm.DB

func truncateTable(tablename, dsn string) {
	DBA, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Println(err)
		return
	}

	stmt := "TRUNCATE TABLE " + tablename + ";"

	if _, err := DBA.Raw(stmt).Rows(); err != nil {
		log.Println(err)
	}
	log.Println("Database Cleaned")
}
