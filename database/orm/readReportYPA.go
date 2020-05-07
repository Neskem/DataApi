package orm

import (
	"DataApi.Go/database/models/YPA"
	"DataApi.Go/lib/common"
	"fmt"
	"github.com/jinzhu/gorm"
	"strings"
)

type YpaSourceReportDaily = YPA.YpaSourceReportDaily

func SelectBetweenDailyYpa(db *gorm.DB, startDate int, endDate int) []common.JSON {
	table := "ypa_source_report_daily"
	var ypaSourceReportDaily YpaSourceReportDaily
	rows, err := db.Table(table).Model(&ypaSourceReportDaily).Where("date BETWEEN ? AND ?", startDate, endDate).Rows()
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	var rowsList []common.JSON
	for rows.Next() {
		var ypaSourceReportDaily YpaSourceReportDaily
		err := db.ScanRows(rows, &ypaSourceReportDaily)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		rowsList = append(rowsList, common.JSON{
			"date":strings.Replace(ypaSourceReportDaily.Date, "-", "", 2),
			"revenue": ypaSourceReportDaily.Revenue,
		})
	}
	return rowsList
}
