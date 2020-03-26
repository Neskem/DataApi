package spark

import (
	"DataApi.Go/database/models/PV"
	"DataApi.Go/database/orm"
	"DataApi.Go/lib/common"
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

	result := orm.QueryUrlList(db, betweenDates, requestBody.Urls)
	c.JSON(200, result)

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

	result := orm.GetMonthlyList(db, strconv.Itoa(requestBody.Month), requestBody.Urls)
	c.JSON(200, result)

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

	result := orm.GetTotalList(db, requestBody.Urls)
	c.JSON(200, result)

}

func GetAuthorPV(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	startDate, _ := strconv.Atoi(c.Query("start_date"))
	endDate, _ := strconv.Atoi(c.Query("end_date"))
	author := c.Query("author")

	betweenDates := common.GetBetweenDays(startDate, endDate, false)
	result := orm.QueryAuthor(db, betweenDates, author)
	c.JSON(200, result)

}

func GetHostPV(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	startDate, _ := strconv.Atoi(c.Query("start_date"))
	endDate, _ := strconv.Atoi(c.Query("end_date"))
	hostName := c.Query("hostname")

	betweenDates := common.GetBetweenDays(startDate, endDate, false)
	result := orm.QueryHost(db, betweenDates, hostName)
	c.JSON(200, result)
}
