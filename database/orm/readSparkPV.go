package orm

import (
	"DataApi.Go/database/models/PV"
	"DataApi.Go/lib/common"
	"fmt"
	"github.com/jinzhu/gorm"
	"strconv"
)

type StatPagePV = PV.StatPagePV
type PVPageIdSum = PV.PVPageIdSum
type PVMonthly = PV.PVMonthly
type SumPV = PV.SumPV

func SelectPvList(db *gorm.DB, date string, pageIds []string) map[string]int {
	table := common.GetSparkPVTableName(date)
	rows, err := db.Table(table).Model(&StatPagePV{}).Select("SUM(pv_valid) AS pv_valid, page_url").Where(
		"page_id IN (?)", pageIds).Group("page_url").Order("page_url").Rows()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer rows.Close()
	rowsMap := make(map[string]int)
	for rows.Next() {
		var statPagePV StatPagePV
		err := db.ScanRows(rows, &statPagePV)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		rowsMap[statPagePV.PageUrl] = statPagePV.PvValid
	}
	return rowsMap
}

var GlobalRows map[string]map[string]map[string]string
func SelectPvDataByPageId(db *gorm.DB, date string, pageIds []string) map[string]map[string]string{
	table := common.GetSparkPVTableName(date)
	rows, err := db.Table(table).Model(&StatPagePV{}).Select("*").Where(
		"page_id IN (?)", pageIds).Rows()
	if err != nil{
		fmt.Println(err)
		return nil
	}
	defer rows.Close()
	rowsMap := make(map[string]map[string]string)
	for rows.Next() {
		var statPagePV StatPagePV
		err := db.ScanRows(rows, &statPagePV)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		inner, ok := rowsMap[statPagePV.PageId]
		if !ok {
			inner = make(map[string]string)
			inner["page_url"] = statPagePV.PageUrl
			inner["pv"] = strconv.Itoa(statPagePV.Pv)
			inner["pv_valid"] = strconv.Itoa(statPagePV.PvValid)
			inner["pv_invalid"] = strconv.Itoa(statPagePV.PvInvalid)
			inner["pv_count"] = strconv.Itoa(statPagePV.PvCount)
			inner["stay_0_count"] = strconv.Itoa(statPagePV.Stay0Count)
			inner["stay_1_count"] = strconv.Itoa(statPagePV.Stay1Count)
			rowsMap[statPagePV.PageId] = inner
		}
	}
	if len(rowsMap) > 0 {
		GlobalRows[date] = rowsMap
	}
	return rowsMap
}

func SelectPvByHost(db *gorm.DB, date string, hostName string) int {
	table := common.GetSparkPVTableName(date)
	var sum SumPV
	db.Table(table).Select("SUM(pv_valid) AS total").Where("page_hostname= ?", hostName).Scan(&sum)
	return sum.Total
}

func SelectPvByAuthorAndHost(db *gorm.DB, date string, author string, hostName string) int {
	table := common.GetSparkPVTableName(date)
	var sum SumPV
	db.Table(table).Select("SUM(pv_valid) AS total").Where("page_author= ? AND page_hostname LIKE ?", author, "%" + hostName).Scan(&sum)
	return sum.Total
}

func SelectTotalPvList(db *gorm.DB, urls []string) []common.JSON {
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

func SelectMonthlyPVList(db *gorm.DB, month string, urls []string) []common.JSON {
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
