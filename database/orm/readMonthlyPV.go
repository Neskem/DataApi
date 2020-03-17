package orm

import (
	"DataApi.Go/database/models/PV"
	"DataApi.Go/lib/common"
	"github.com/jinzhu/gorm"
)

type PVMonthly = PV.PVMonthly


func QueryMonthlyPV(db *gorm.DB, month string, url string) int {
	table := common.GetPVMonthlyTableName(month)
	pageID := common.GetPageID(url)
	var monthlyPv PVMonthly
	db.Table(table).Where("page_id = ?", pageID).First(&monthlyPv)
	return monthlyPv.Pv
}

func GetMonthlyList(db *gorm.DB, month string, urls []string) []common.JSON{
	result := make(chan int)
	go func() {
		for _, url := range urls {
			pv := QueryMonthlyPV(db, month, url)
			result <- pv
		}
		close(result)
	}()
	var response []common.JSON

	index := 0
	for n := range result {
		response = append(response, common.JSON{
			"url": urls[index],
			"pv_valid": n,
		})
		index = index + 1
	}
	return response
}
