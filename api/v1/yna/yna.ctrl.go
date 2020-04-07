package yna

import (
	"DataApi.Go/lib/common"
	"DataApi.Go/task"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func PostYNAReport(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	type RequestBody struct {
		AdUnitIds []int `json:"adunit_ids" binding:"required"`
		StartDate int `json:"start_date" binding:"required"`
		EndDate int `json:"end_date" binding:"required"`
	}
	var requestBody RequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		fmt.Println(err)
		c.AbortWithStatus(400)
		return
	}

	result := task.QueryYnaReportList(db, requestBody.StartDate, requestBody.EndDate, requestBody.AdUnitIds)
	response := common.JSON{
		"status": true,
		"data": result,
	}
	c.JSON(200, response)
}
