package task

import (
	"DataApi.Go/database/orm"
	"DataApi.Go/lib/common"
	"fmt"
	"github.com/jinzhu/gorm"
	"sync"
)

func QueryTotalPvList(db *gorm.DB, urls []string) []common.JSON{
	response := orm.SelectTotalPvList(db, urls)
	return response
}

func QueryMonthlyPvList(db *gorm.DB, month string, urls []string) []common.JSON{
	response := orm.SelectMonthlyPVList(db, month, urls)
	return response
}

func QueryPvListByHost(db *gorm.DB, betweenDate []string, hostName string) int{
	result := make(chan int, len(betweenDate))
	wg := sync.WaitGroup{}
	wg.Add(len(betweenDate))
	for _, date := range betweenDate {
		go func(date string) {
			defer wg.Done()
			pv := orm.SelectPvByHost(db, date, hostName)
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

func QueryPvListByAuthorAndHost(db *gorm.DB, betweenDate []string, author string, hostName string) int{
	result := make(chan int, len(betweenDate))
	wg := sync.WaitGroup{}
	wg.Add(len(betweenDate))
	for _, date := range betweenDate {
		go func(date string) {
			defer wg.Done()
			pv := orm.SelectPvByAuthorAndHost(db, date, author, hostName)
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

func QueryDailyPvList(db *gorm.DB, betweenDate []string, urls []string) map[string]int{
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
			rowsMap := orm.SelectPvList(db, date, pageIds)
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

func MoveDailyPv(db *gorm.DB, author string, betweenDate []string, mappings []map[string]string) bool {
	//result := make(chan map[string]common.JSON, len(betweenDate))
	var mutex sync.Mutex
	var pageIdMappings []map[string]string
	var pageIdsArray []string
	orm.GlobalRows = make(map[string]map[string]map[string]string)
	for _, mapping := range mappings {
		newUrl := "https://zi.media/" + "@" + author + "/" + "post" + "/" + mapping["new"]
		newPageId := common.GetPageID(newUrl)
		oldUrl := "https://zi.media/" + "@" + author + "/" + "post" + "/" + mapping["old"]
		oldPageId := common.GetPageID(oldUrl)
		pageIdMapping := make(map[string]string)
		pageIdMapping["old"] = oldPageId
		pageIdMapping["new"] = newPageId
		pageIdsArray = append(pageIdsArray, oldPageId)
		pageIdsArray = append(pageIdsArray, newPageId)
		pageIdMappings = append(pageIdMappings, pageIdMapping)
	}
	wg := sync.WaitGroup{}
	wg.Add(len(betweenDate))
	for _, date := range betweenDate {
		go func(date string) {
			defer wg.Done()
			mutex.Lock()
			orm.SelectPvDataByPageId(db, date, pageIdsArray)
			//result <- pvData
			mutex.Unlock()
		}(date)
	}
	go func() {
		wg.Wait()
		//close(result)
		for date, pageIdMap := range orm.GlobalRows {
			fmt.Println(date)
			for pageId, dataMap := range pageIdMap {
				fmt.Println(pageId )
				for key, value := range dataMap {
					fmt.Println(key, ": ", value)
				}
			}
		}
	}()

	return true
}
