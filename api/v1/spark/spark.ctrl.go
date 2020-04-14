package spark

import (
	"DataApi.Go/database/models/PV"
	"DataApi.Go/lib/common"
	"DataApi.Go/task"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
)

type StatPagePV = PV.StatPagePV

func PostDailyPV(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	type RequestBody struct {
		Urls []string `json:"urls" binding:"required"`
		StartDate string `json:"start_date" binding:"required"`
		EndDate string `json:"end_date" binding:"required"`
	}
	var requestBody RequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		fmt.Println(err)
		c.AbortWithStatus(400)
		return
	}

	urls := common.Unique(requestBody.Urls)
	startDate, _ := strconv.Atoi(requestBody.StartDate)
	endDate, _ := strconv.Atoi(requestBody.EndDate)
	betweenDates := common.GetBetweenDays(startDate, endDate, false)
	result := task.QueryDailyPvList(db, betweenDates, urls)
	var response []common.JSON
	for _, url := range requestBody.Urls {
		data := common.JSON{
			"start_date": requestBody.StartDate,
			"end_date": requestBody.EndDate,
			"url": url,
			"page_id": common.GetPageID(url),
			"pv_valid": 0 + result[url],
		}
		response = append(response, data)
	}
	c.JSON(200, common.JSON{
		"status": true,
		"data": response,
	})
}

func PostMonthlyPV(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	type RequestBody struct {
		Urls []string `json:"urls" binding:"required"`
		Month string `json:"month" binding:"required"`
	}
	var requestBody RequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		fmt.Println(err)
		c.AbortWithStatus(400)
		return
	}

	urls := common.Unique(requestBody.Urls)
	result := task.QueryMonthlyPvList(db, requestBody.Month, urls)
	c.JSON(200, common.JSON{
		"status": true,
		"data": result,
	})
}

func PostTotalPV(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	type RequestBody struct {
		Urls []string `json:"urls" binding:"required"`
	}
	var requestBody RequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		fmt.Println(err)
		c.AbortWithStatus(400)
		return
	}

	urls := common.Unique(requestBody.Urls)
	result := task.QueryTotalPvList(db, urls)
	c.JSON(200, common.JSON{
		"status": true,
		"data": result,
	})
}

func PostMovePV(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	type RequestBody struct {
		Mappings []map[string]string `json:"mappings" binding:"required"`
		StartDate string `json:"start_date" binding:"required"`
		EndDate string `json:"end_date" binding:"required"`
		Author string `json:"author" binding:"required"`
	}
	var requestBody RequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		fmt.Println(err)
		c.AbortWithStatus(400)
		return
	}
	startDate, _ := strconv.Atoi(requestBody.StartDate)
	endDate, _ := strconv.Atoi(requestBody.EndDate)
	betweenDates := common.GetBetweenDays(startDate, endDate, false)
	result := task.MoveDailyPv(db, requestBody.Author, betweenDates, requestBody.Mappings)
	c.JSON(200, common.JSON{
		"status": result,
	})
}

func GetAuthorPV(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	startDate, _ := strconv.Atoi(c.Query("start_date"))
	endDate, _ := strconv.Atoi(c.Query("end_date"))
	author := c.Query("author")
	hostName := c.Query("hostname")
	if hostName == ""{
		hostName = "zi.media"
	}

	betweenDates := common.GetBetweenDays(startDate, endDate, false)
	pv := task.QueryPvListByAuthorAndHost(db, betweenDates, author, hostName)
	result := common.JSON{
		"status": true,
		"data": common.JSON{
			"start_date": startDate,
			"end_date":   endDate,
			"author":   author,
			"hostname": hostName,
			"pv_valid":   pv,
		},
	}
	c.JSON(200, result)
}

func GetHostPV(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	startDate, _ := strconv.Atoi(c.Query("start_date"))
	endDate, _ := strconv.Atoi(c.Query("end_date"))
	hostName := c.Query("hostname")

	betweenDates := common.GetBetweenDays(startDate, endDate, false)
	pv := task.QueryPvListByHost(db, betweenDates, hostName)
	result := common.JSON{
		"status": true,
		"data": common.JSON{
			"start_date": startDate,
			"end_date":   endDate,
			"hostname":   hostName,
			"pv_valid":   pv,
		},
	}
	c.JSON(200, result)
}