package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"os"
)

func Initialize() (*gorm.DB, error) {
	dbConfig := os.Getenv("DB_CONFIG")
	db, err := gorm.Open("mysql", dbConfig)
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	fmt.Println("Connected to database")
	//models.Migrate(db)
	return db, err
}
