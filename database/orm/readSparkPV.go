package orm

import (
	"DataApi.Go/database/models/PV"
	"DataApi.Go/lib/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type StatPagePV = PV.StatPagePV
type SumPV = PV.SumPV

func readDailyPV(c *gin.Context, tableName string) common.JSON {
	db := c.MustGet("db").(*gorm.DB)

	type RequestBody struct {
		Url string `json:"url" binding:"required"`
	}
	var requestBody RequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		fmt.Println(err)
		c.AbortWithStatus(400)
		return nil
	}

	var stat StatPagePV
	db.Table(tableName).Where("page_url = ?", requestBody.Url).First(&stat)

	return stat.Serialize()
}

func QuerySparkPVByUrl(db *gorm.DB, betweenDate []string, url string) int {
	b := make(chan int)
	go func() {
		for _, s := range betweenDate {
			table := common.GetSparkPVTableName(s)
			var stat StatPagePV
			db.Table(table).Where("page_url = ?", url).First(&stat)
			b <- stat.Pv
		}
		close(b)
	}()
	result := 0

	for n := range b {
		result += n
	}
	return result
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
