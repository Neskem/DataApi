package orm

import (
	"DataApi.Go/database/models/PV"
	"DataApi.Go/lib/common"
	"fmt"
	"github.com/jinzhu/gorm"
	"sync"
)

type StatPagePV = PV.StatPagePV
type SumPV = PV.SumPV

func QueryHostByAuthor(db *gorm.DB, date string, author string) string {
	table := common.GetSparkPVTableName(date)
	var stat StatPagePV
	db.Table(table).Where("page_author = ?", author).First(&stat)
	return stat.PageHostname
}

func QuerySparkPVByUrl(db *gorm.DB, betweenDate []string, url string) int {
	b := make(chan int, len(betweenDate))
	go func() {
		for _, s := range betweenDate {
			table := common.GetSparkPVTableName(s)
			var stat StatPagePV
			db.Table(table).Where("page_url = ?", url).First(&stat)
			b <- stat.Pv
		}
		close(b)
	}()
	pv := 0

	for n := range b {
		pv += n
	}
	return pv
}

func QueryPVByUrl(db *gorm.DB, betweenDate []string, url string) int {
	result := make(chan int, len(betweenDate))
	wg := sync.WaitGroup{}
	wg.Add(len(betweenDate))
	for _, date := range betweenDate {
		go func(date string) {
			defer wg.Done()
			table := common.GetSparkPVTableName(date)
			var stat StatPagePV
			db.Table(table).Where("page_url = ?", url).First(&stat)
			result <- stat.Pv
		}(date)
	}
	go func() {
		wg.Wait()
		close(result)
	}()
	pv := 0

	for n := range result {
		pv += n
	}
	return pv
}

func QueryPVByUrl2(db *gorm.DB, date string, pageIds []string) map[string]int {
	table := common.GetSparkPVTableName(date)
	rows, err := db.Table(table).Model(&StatPagePV{}).Select("SUM(pv_valid) AS pv_valid, page_url").Where("page_id IN (?)", pageIds).Group("page_url").Order("page_url").Rows()
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

func QuerySparkPVByAuthor(db *gorm.DB, betweenDate []string, author string) int {
	pvChan := make(chan int)
	go func() {
		for _, date := range betweenDate {
			table := common.GetSparkPVTableName(date)
			var sum SumPV
			db.Table(table).Select("sum(pv) as total").Where("page_author = ?", author).Scan(&sum)
			pvChan <- sum.Total
		}
		close(pvChan)
	}()
	result := 0

	for n := range pvChan {
		result += n
	}
	return result
}

func QuerySparkPVByHost(db *gorm.DB, betweenDate []string, hostName string) int {
	pvChan := make(chan int)
	go func() {
		for _, date := range betweenDate {
			table := common.GetSparkPVTableName(date)
			var sum SumPV
			db.Table(table).Select("sum(pv) as total").Where("page_hostname= ?", hostName).Scan(&sum)
			pvChan <- sum.Total
		}
		close(pvChan)
	}()

	result := 0

	for n := range pvChan {
		result += n
	}
	return result
}

func QueryPvValidByHost(db *gorm.DB, date string, hostName string) int {
	table := common.GetSparkPVTableName(date)
	var sum SumPV
	db.Table(table).Select("sum(pv_valid) as total").Where("page_hostname= ?", hostName).Scan(&sum)
	return sum.Total
}

func QueryPvValidByAuthor(db *gorm.DB, date string, author string) int {
	table := common.GetSparkPVTableName(date)
	var sum SumPV
	db.Table(table).Select("sum(pv_valid) as total").Where("page_author= ?", author).Scan(&sum)
	return sum.Total
}

func QueryPvValidByAuthorAndHost(db *gorm.DB, date string, author string, hostName string) int {
	table := common.GetSparkPVTableName(date)
	var sum SumPV
	db.Table(table).Select("sum(pv_valid) as total").Where("page_author= ? AND page_hostname LIKE ?", author, "%" + hostName).Scan(&sum)
	return sum.Total
}

func QueryUrlList(db *gorm.DB, betweenDate []string, urls []string) []common.JSON{
	result := make(chan int)
	go func() {
		for _, url := range urls {
			pv := QuerySparkPVByUrl(db, betweenDate, url)
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

func QueryAuthor(db *gorm.DB, betweenDate []string, author string) common.JSON{
	pv := QuerySparkPVByAuthor(db, betweenDate, author)
	response := common.JSON{
		"status": true,
		"data": common.JSON{
			"start_date": betweenDate[0],
			"end_date":   betweenDate[len(betweenDate)-1],
			"author":   author,
			"pv_valid":   pv,
		},
	}
	return response
}

func QueryHost(db *gorm.DB, betweenDate []string, hostName string) common.JSON{
	pv := QuerySparkPVByHost(db, betweenDate, hostName)
	response := common.JSON{
		"status": true,
		"data": common.JSON{
			"start_date": betweenDate[0],
			"end_date":   betweenDate[len(betweenDate)-1],
			"hostname":   hostName,
			"pv_valid":   pv,
		},
	}
	return response
}


func QueryAllPvByHost(db *gorm.DB, betweenDate []string, hostName string) int{
	result := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(len(betweenDate))
	for _, date := range betweenDate {
		fmt.Println(date)
		go func(date string) {
			defer wg.Done()
			pv := QueryPvValidByHost(db, date, hostName)
			result <- pv
		}(date)
	}
	go func() {
		wg.Wait()
		close(result)
	}()
	var totalPv int
	for pv := range result {
		totalPv = totalPv + pv
	}
	return totalPv
}

type queryPv func(db *gorm.DB, date string, queryTarget string) int

func QueryAllPvByAuthor(db *gorm.DB, betweenDate []string, author string, function queryPv) int{
	result := make(chan int, len(betweenDate))
	wg := sync.WaitGroup{}
	wg.Add(len(betweenDate))
	for _, date := range betweenDate {
		fmt.Println(date)
		go func(date string) {
			defer wg.Done()
			pv := function(db, date, author)
			result <- pv
		}(date)
	}
	go func() {
		wg.Wait()
		close(result)
	}()
	var totalPv int
	for pv := range result {
		totalPv = totalPv + pv
	}
	return totalPv
}

func QueryUrlPvList(db *gorm.DB, betweenDate []string, urls []string) []common.JSON{
	result := make(chan int, len(urls))
	wg := sync.WaitGroup{}
	wg.Add(len(urls))
	for _, url := range urls{
		go func(url string) {
			defer wg.Done()
			pv := QueryPVByUrl(db, betweenDate, url)
			result <- pv
		}(url)
	}
	go func() {
		wg.Wait()
		close(result)
	}()
	var response []common.JSON

	index := 0
	for n := range result {
		fmt.Println("url:", n)
		response = append(response, common.JSON{
			"url": urls[index],
			"pv_valid": n,
			"start_date": betweenDate[0],
			"end_date": betweenDate[len(betweenDate)-1],
			"page_id": common.GetPageID(urls[index]),
		})
		index = index + 1
	}
	return response
}

func QueryUrlPvList2(db *gorm.DB, betweenDate []string, urls []string) map[string]int{
	result := make(chan map[string]int, len(betweenDate))
	var pageIds []string
	for _, url := range urls {
		pageId := common.GetPageID(url)
		pageIds = append(pageIds, pageId)
	}
	wg := sync.WaitGroup{}
	wg.Add(len(betweenDate))
	for _, date := range betweenDate{
		go func(date string) {
			defer wg.Done()
			rowsMap := QueryPVByUrl2(db, date, pageIds)
			result <- rowsMap
		}(date)
	}
	go func() {
		wg.Wait()
		close(result)
	}()
	keyVal := make(map[string]int)
	for _, url := range urls {
		keyVal[url] = 0
	}

	index := 0
	for n := range result {
		for _, url := range urls {
			keyVal[url] = keyVal[url] + n[url]
		}
		index = index + 1
	}
	return keyVal
}
