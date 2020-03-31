package task

import (
	"DataApi.Go/database/orm"
	"DataApi.Go/lib/common"
	"github.com/jinzhu/gorm"
)

func QueryYnaReportList(db *gorm.DB, StartDate int, EndDate int, adUnitIds []int) []common.JSON {
	result := orm.QueryYnaReportFix(db, adUnitIds, StartDate, EndDate)
	return result
}
