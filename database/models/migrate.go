package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB)  {
	db.AutoMigrate(&StatPagePV{})
	db.AutoMigrate(&Post{})
	fmt.Println("Auto Migration has been processed")

}
