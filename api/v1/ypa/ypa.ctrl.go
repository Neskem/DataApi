package ypa

import (
	"DataApi.Go/lib/common"
	"DataApi.Go/task"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
)

func GetDailyYPA(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	startDate, _ := strconv.Atoi(c.Query("start_date"))
	endDate, _ := strconv.Atoi(c.Query("end_date"))

	result := task.QueryDailyYpaList(db, startDate, endDate)
	response := common.JSON{
		"status": true,
		"data": result,
	}
	c.JSON(200, response)
}

func PostAllotting(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	type RequestBody struct {
		Username []string `json:"username" binding:"required"`
		StartDate string `json:"month" start_date:"required"`
		EndDate string `json:"month" end_date:"required"`
	}
	var requestBody RequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		fmt.Println(err)
		c.AbortWithStatus(400)
		return
	}

	startDate, _ := strconv.Atoi(requestBody.StartDate)
	endDate, _ := strconv.Atoi(requestBody.EndDate)

	result := task.QueryDailyYpaList(db, startDate, endDate)
	response := common.JSON{
		"status": true,
		"data": result,
	}
	c.JSON(200, response)
}
