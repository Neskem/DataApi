package orm

import (
	"DataApi.Go/database/models/YNA"
	"DataApi.Go/lib/common"
	"fmt"
	"github.com/jinzhu/gorm"
	"strconv"
	"sync"
)

type YnaReport = YNA.YnaReprot

func QueryYnaReport(db *gorm.DB, adUnitId int, date int) common.JSON {
	table := "yna_report"
	var ynaReport YnaReport
	db.Table(table).Where("adunit_id = ? and date = ?", adUnitId, date).First(&ynaReport)
	fmt.Println(ynaReport)
	return common.JSON{
		"date": date,
		"adunit_id": adUnitId,
		"impressions": ynaReport.Impressions,
		"clicks": ynaReport.Clicks,
		"revenueInTWD": ynaReport.Revenueintwd,
		"customerRevenueInTWD": ynaReport.Customerrevenueintwd,
	}
}

func QueryYnaReportList(db *gorm.DB, betweenDate []string, adUnitIds []int) []common.JSON {
	result := make(chan common.JSON)
	wg := sync.WaitGroup{}
	wg.Add(len(adUnitIds))
	for _, id := range adUnitIds {
		go func(result chan<- common.JSON) {
			defer wg.Done()
			for _, date := range betweenDate {
				dateTime, _ := strconv.Atoi(date)
				revenue := QueryYnaReport(db, id, dateTime)
				result <- revenue
			}
		}(result)
	}
	go func(){
		wg.Wait()
		close(result)
	}()
	var response []common.JSON

	index := 0
	for n := range result {
		response = append(response, n)
		index = index + 1
	}
	return response
}
