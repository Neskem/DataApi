package task

import (
	"DataApi.Go/database/orm"
	"DataApi.Go/lib/common"
	"github.com/jinzhu/gorm"
)

func QueryDailyYpaList(db *gorm.DB, startDate int, endDate int) []common.JSON {
	response := orm.QueryBetweenDailyYpa(db, startDate, endDate)
	return response
}
