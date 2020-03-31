package task

import (
	"DataApi.Go/database/orm"
	"DataApi.Go/lib/common"
	"github.com/jinzhu/gorm"
	"sync"
)

type queryPv func(db *gorm.DB, date string, queryTarget string) int

func GetTotalPvList(db *gorm.DB, urls []string) []common.JSON{
	response := orm.QueryTotalPVList(db, urls)
	return response
}

func GetMonthlyPvList(db *gorm.DB, month string, urls []string) []common.JSON{
	response := orm.QueryMonthlyPVList(db, month, urls)
	return response
}

func GetPvListByFunction(db *gorm.DB, betweenDate []string, target string, function queryPv) int{
	result := make(chan int, len(betweenDate))
	wg := sync.WaitGroup{}
	wg.Add(len(betweenDate))
	for _, date := range betweenDate {
		go func(date string) {
			defer wg.Done()
			pv := function(db, date, target)
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

func GetPvListByAuthor(db *gorm.DB, betweenDate []string, author string, hostName string) int{
	result := make(chan int, len(betweenDate))
	wg := sync.WaitGroup{}
	wg.Add(len(betweenDate))
	for _, date := range betweenDate {
		go func(date string) {
			defer wg.Done()
			pv := orm.QueryPvValidByAuthorAndHost(db, date, author, hostName)
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

func GetDailyPvList(db *gorm.DB, betweenDate []string, urls []string) []common.JSON{
	result := make(chan int, len(urls))
	go func() {
		for _, url := range urls {
			pv := orm.QuerySparkPVByUrl(db, betweenDate, url)
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
