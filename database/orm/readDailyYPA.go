package orm

import (
	"DataApi.Go/database/models/YPA"
	"DataApi.Go/lib/common"
	"fmt"
	"github.com/jinzhu/gorm"
)

type YpaReportDaily = YPA.YpaReportDaily

func QueryDailyYpa(db *gorm.DB, date string) float64 {
	table := "ypa_report_daily"
	var ypaReportDaily YpaReportDaily
	db.Table(table).Where("date = ?", date).First(&ypaReportDaily)
	return ypaReportDaily.Revenue
}

func QueryBetweenDailyYpa(db *gorm.DB, startDate int, endDate int) []common.JSON {
	table := "ypa_report_daily"
	var ypaReportDaily YpaReportDaily
	rows, err := db.Table(table).Model(&ypaReportDaily).Where("date BETWEEN ? AND ?", startDate, endDate).Rows()
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	var rowsList []common.JSON
	for rows.Next() {
		var ypaReportDaily YpaReportDaily
		err := db.ScanRows(rows, &ypaReportDaily)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		rowsList = append(rowsList, common.JSON{
			"date": ypaReportDaily.Date,
			"revenue": ypaReportDaily.Revenue,
		})
	}
	return rowsList
}

func QueryDailyYpaList(db *gorm.DB, startDate int, endDate int) []common.JSON {
	response := QueryBetweenDailyYpa(db, startDate, endDate)
	return response
}
