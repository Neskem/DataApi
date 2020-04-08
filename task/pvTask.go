package task

import (
	"DataApi.Go/database/orm"
	"DataApi.Go/lib/common"
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
	var mutex sync.Mutex
	var pageIdsArray []string
	var globalRows map[string]map[string]map[string]int
	globalRows = make(map[string]map[string]map[string]int)
	var mappingMap map[string]string
	mappingMap = make(map[string]string)
	for _, mapping := range mappings {
		newUrl := common.GetZiUrl(author, mapping["new"])
		newPageId := common.GetPageID(newUrl)
		oldUrl := common.GetZiUrl(author, mapping["old"])
		oldPageId := common.GetPageID(oldUrl)
		pageIdMapping := make(map[string]string)
		pageIdMapping["old"] = oldPageId
		pageIdMapping["new"] = newPageId
		mappingMap[newPageId] = oldPageId
		pageIdsArray = append(pageIdsArray, oldPageId)
		pageIdsArray = append(pageIdsArray, newPageId)
	}
	wg := sync.WaitGroup{}
	wg.Add(len(betweenDate))
	for _, date := range betweenDate {
		go func(date string) {
			defer wg.Done()
			mutex.Lock()
			globalRows = orm.SelectPvDataByPageId(db, date, pageIdsArray, globalRows)
			mutex.Unlock()
		}(date)
	}

	go func() {
		wg.Wait()
		globalRows = MoveDailyPvFunc(globalRows, mappingMap)
		updateDailyPvFunc(db, globalRows)
	}()
	return true
}

func MoveDailyPvFunc(rows map[string]map[string]map[string]int, mappingMap map[string]string)map[string]map[string]map[string]int{
	for date, pageIdMap := range rows {
		for pageId := range pageIdMap {
			oldPageId, ok := mappingMap[pageId]
			if ok {
				_, ok := rows[date][oldPageId]
				if ok {
					rows[date][pageId]["pv"] = rows[date][pageId]["pv"] + rows[date][oldPageId]["pv"]
					rows[date][pageId]["pv_valid"] = rows[date][pageId]["pv_valid"] + rows[date][oldPageId]["pv_valid"]
					rows[date][pageId]["pv_invalid"] = rows[date][pageId]["pv_invalid"] + rows[date][oldPageId]["pv_invalid"]
					rows[date][pageId]["pv_count"] = rows[date][pageId]["pv_count"] + rows[date][oldPageId]["pv_count"]
					rows[date][pageId]["stay_0_count"] = rows[date][pageId]["stay_0_count"] + rows[date][oldPageId]["stay_0_count"]
					rows[date][pageId]["stay_1_count"] = rows[date][pageId]["stay_1_count"] + rows[date][oldPageId]["stay_1_count"]

					rows[date][oldPageId]["pv"] = 0
					rows[date][oldPageId]["pv_valid"] = 0
					rows[date][oldPageId]["pv_invalid"] = 0
					rows[date][oldPageId]["pv_count"] = 0
					rows[date][oldPageId]["stay_0_count"] = 0
					rows[date][oldPageId]["stay_1_count"]= 0
				}
			}
		}
	}
	return rows
}

func updateDailyPvFunc(db *gorm.DB, rows map[string]map[string]map[string]int) {
	for date, pageIdMap := range rows {
		wg := sync.WaitGroup{}
		wg.Add(len(pageIdMap))
		for pageId, dataMap := range pageIdMap {
			go func(db *gorm.DB, date string, pageId string, dataMap map[string]int) {
				defer wg.Done()
				orm.UpdatePvDataByPageId(db, date, pageId, dataMap)
			}(db, date, pageId, dataMap)
		}
		wg.Wait()
	}
}
