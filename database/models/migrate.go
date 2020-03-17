package models

import (
	"DataApi.Go/database/models/PV"
	"fmt"
	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB)  {
	db.AutoMigrate(&PV.StatPagePV{})
	db.AutoMigrate(&Post{})
	fmt.Println("Auto Migration has been processed")

}
