package adSense

import (
	"DataApi.Go/lib/common"
	"DataApi.Go/task"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
)

func PostDailyAdSense(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	type RequestBody struct {
		AccountIds []string `json:"account_ids" binding:"required"`
		StartDate string `json:"start_date" binding:"required"`
		EndDate string `json:"end_date" binding:"required"`
	}
	var requestBody RequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		fmt.Println(err)
		c.AbortWithStatus(400)
		return
	}
	accountIds := common.Unique(requestBody.AccountIds)
	startDate, _ := strconv.Atoi(requestBody.StartDate)
	endDate, _ := strconv.Atoi(requestBody.EndDate)
	result := task.QueryAdSenseReportList(db, accountIds, startDate, endDate)
	c.JSON(200, common.JSON{
		"status": true,
		"data": result,
	})
}

func PostDailyAdSenseRevenue(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	type RequestBody struct {
		AccountIds []string `json:"account_ids" binding:"required"`
		StartDate string `json:"start_date" binding:"required"`
		EndDate string `json:"end_date" binding:"required"`
	}
	var requestBody RequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		fmt.Println(err)
		c.AbortWithStatus(400)
		return
	}
	accountIds := common.Unique(requestBody.AccountIds)
	startDate, _ := strconv.Atoi(requestBody.StartDate)
	endDate, _ := strconv.Atoi(requestBody.EndDate)
	result := task.QueryAdSenseRevenueList(db, accountIds, startDate, endDate)
	c.JSON(200, common.JSON{
		"status": true,
		"data": result,
	})
}

func GetAdSenseDomains(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	result := task.QueryAdSenseDomainList(db)
	c.JSON(200, common.JSON{
		"status": true,
		"data": result,
	})
}