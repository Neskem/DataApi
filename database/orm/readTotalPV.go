package orm

import (
	"DataApi.Go/database/models/PV"
	"DataApi.Go/lib/common"
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
