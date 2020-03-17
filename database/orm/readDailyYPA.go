package orm

import (
	"DataApi.Go/database/models/YPA"
	"DataApi.Go/lib/common"
	"fmt"
	"github.com/jinzhu/gorm"
	"strings"
)

type YpaReportDaily = YPA.YpaReportDaily

func QueryDailyYpa(db *gorm.DB, date string) float64 {
	table := "ypa_report_daily"
	var ypaReportDaily YpaReportDaily
	db.Table(table).Where("date = ?", date).First(&ypaReportDaily)
	return ypaReportDaily.Revenue
}

func QueryDailyYpaList(db *gorm.DB, betweenDate []string) []common.JSON {
	result := make(chan float64)
	go func() {
		for _, date := range betweenDate {
			revenue := QueryDailyYpa(db, date)
			fmt.Println("revenue: ", revenue)
			result <- revenue
		}
		close(result)
	}()
	var response []common.JSON

	index := 0
	for n := range result {
		response = append(response, common.JSON{
			"date": strings.Trim(betweenDate[index], "-"),
			"revenue": n,
		})
		index = index + 1
	}
	return response
}
