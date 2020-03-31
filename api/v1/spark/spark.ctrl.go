package spark

import (
	"DataApi.Go/database/models/PV"
	"DataApi.Go/database/orm"
	"DataApi.Go/lib/common"
	"DataApi.Go/task"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
)

type StatPagePV = PV.StatPagePV

func ReadDailyPV(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	type RequestBody struct {
		Urls []string `json:"urls" binding:"required"`
		StartDate int `json:"start_date" binding:"required"`
		EndDate int `json:"end_date" binding:"required"`
	}
	var requestBody RequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		fmt.Println(err)
		c.AbortWithStatus(400)
		return
	}

	betweenDates := common.GetBetweenDays(requestBody.StartDate, requestBody.EndDate, false)

	result := orm.QueryUrlPvList2(db, betweenDates, requestBody.Urls)
	c.JSON(200, common.JSON{
		"status": true,
		"data": result,
	})

}

func ReadMonthlyPV(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	type RequestBody struct {
		Urls []string `json:"urls" binding:"required"`
		Month int `json:"month" binding:"required"`
	}
	var requestBody RequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		fmt.Println(err)
		c.AbortWithStatus(400)
		return
	}

	result := orm.GetMonthlyPVList(db, strconv.Itoa(requestBody.Month), requestBody.Urls)
	c.JSON(200, common.JSON{
		"status": true,
		"data": result,
	})

}

func ReadTotalPV(c *gin.Context) {
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
	result := orm.GetTotalPVList(db, urls)
	c.JSON(200, common.JSON{
		"status": true,
		"data": result,
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
	pv := task.GetPvListByAuthor(db, betweenDates, author, hostName)
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

//func GetHostPV(c *gin.Context) {
//	db := c.MustGet("db").(*gorm.DB)
//	startDate, _ := strconv.Atoi(c.Query("start_date"))
//	endDate, _ := strconv.Atoi(c.Query("end_date"))
//	hostName := c.Query("hostname")
//
//	betweenDates := common.GetBetweenDays(startDate, endDate, false)
//	result := orm.QueryHost(db, betweenDates, hostName)
//	c.JSON(200, result)
//}

func GetHostPV(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	startDate, _ := strconv.Atoi(c.Query("start_date"))
	endDate, _ := strconv.Atoi(c.Query("end_date"))
	hostName := c.Query("hostname")

	betweenDates := common.GetBetweenDays(startDate, endDate, false)
	pv := task.GetPvListByFunction(db, betweenDates, hostName, orm.QueryPvValidByHost)
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