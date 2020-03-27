package orm

import (
	"DataApi.Go/database/models/PV"
	"DataApi.Go/lib/common"
	"fmt"
	"github.com/jinzhu/gorm"
)

type PVPageIdSum = PV.PVPageIdSum

func QueryTotalPV(db *gorm.DB, url string) int {
	table := "pv_pageid_sum"
	pageID := common.GetPageID(url)
	var pvPageIdSum PVPageIdSum
	db.Table(table).Where("page_id = ?", pageID).First(&pvPageIdSum)
	return pvPageIdSum.Pv
}

func QueryTotalPVList(db *gorm.DB, urls []string) []common.JSON {
	table := "pv_pageid_sum"
	var pageIds []string
	pageIdMap := make(map[string]string)
	for _, url := range urls {
		pageId := common.GetPageID(url)
		pageIdMap[pageId] = url
		pageIds = append(pageIds, pageId)
	}
	var pvPageIdSum PVPageIdSum
	rows, err := db.Table(table).Model(pvPageIdSum).Where("page_id IN (?)", pageIds).Rows()
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	var rowsList []common.JSON
	for rows.Next() {
		var pvPageIdSum PVPageIdSum
		err := db.ScanRows(rows, &pvPageIdSum)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		rowsList = append(rowsList, common.JSON{
			"url": pageIdMap[pvPageIdSum.PageId],
			"page_id": pvPageIdSum.PageId,
			"pv_valid": pvPageIdSum.PvValid,
		})
	}
	return rowsList
}

func GetTotalList(db *gorm.DB, urls []string) []common.JSON{
	result := make(chan int)
	go func() {
		for _, url := range urls {
			pv := QueryTotalPV(db, url)
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

func GetTotalPVList(db *gorm.DB, urls []string) []common.JSON{
	response := QueryTotalPVList(db, urls)
	return response
}