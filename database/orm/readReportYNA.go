package orm

import (
	"DataApi.Go/database/models/YNA"
	"DataApi.Go/lib/common"
	"fmt"
	"github.com/jinzhu/gorm"
)

type YnaReport = YNA.YnaReprot

func SelectYnaReportList(db *gorm.DB, adUnitId []int, startDate int, endDate int) []common.JSON {
	table := "yna_report"
	var ynaReport YnaReport
	rows, err := db.Table(table).Model(&ynaReport).Where("adunit_id IN (?) AND DATE BETWEEN ? AND ?", adUnitId, startDate, endDate).Rows()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer rows.Close()
	var rowsList []common.JSON
	for rows.Next() {
		var ynaReport YnaReport
		err := db.ScanRows(rows, &ynaReport)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		rowsList = append(rowsList, common.JSON{
			"date": ynaReport.Date,
			"adunit_id": ynaReport.AdUnitId,
			"impressions": ynaReport.Impressions,
			"clicks": ynaReport.Clicks,
			"revenueInTWD": ynaReport.Revenueintwd,
			"customerRevenueInTWD": ynaReport.Customerrevenueintwd,
		})
	}
	return rowsList
}
