package orm

import (
	"DataApi.Go/database/models/PV"
	"DataApi.Go/lib/common"
	"fmt"
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

func QueryMonthlyPVList(db *gorm.DB, month string, urls []string) []common.JSON {
	table := common.GetPVMonthlyTableName(month)
	var pageIds []string
	pageIdMap := make(map[string]string)
	for _, url := range urls {
		pageId := common.GetPageID(url)
		pageIdMap[pageId] = url
		pageIds = append(pageIds, pageId)
	}
	var monthlyPv PVMonthly
	rows, err := db.Table(table).Model(monthlyPv).Where("page_id IN (?)", pageIds).Rows()
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	var rowsList []common.JSON
	for rows.Next() {
		var monthlyPv PVMonthly
		err := db.ScanRows(rows, &monthlyPv)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		rowsList = append(rowsList, common.JSON{
			"url": pageIdMap[monthlyPv.PageId],
			"page_id": monthlyPv.PageId,
			"pv_valid": monthlyPv.PvValid,
		})
	}
	return rowsList
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

func GetMonthlyPVList(db *gorm.DB, month string, urls []string) []common.JSON{
	response := QueryMonthlyPVList(db, month, urls)
	return response
}