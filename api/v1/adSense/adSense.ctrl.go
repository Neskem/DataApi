package adSense

import (
	"DataApi.Go/database/orm"
	"DataApi.Go/lib/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func ReadDailyAdSense(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	type RequestBody struct {
		AccountIds []string `json:"account_ids" binding:"required"`
		StartDate int `json:"start_date" binding:"required"`
		EndDate int `json:"end_date" binding:"required"`
	}
	var requestBody RequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		fmt.Println(err)
		c.AbortWithStatus(400)
		return
	}
	accountIds := common.Unique(requestBody.AccountIds)
	result := orm.QueryAdSenseReportList(db, accountIds, requestBody.StartDate, requestBody.EndDate)
	c.JSON(200, common.JSON{"status": true, "data": result})
}

func ReadDailyAdSenseRevenue(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	type RequestBody struct {
		AccountIds []string `json:"account_ids" binding:"required"`
		StartDate int `json:"start_date" binding:"required"`
		EndDate int `json:"end_date" binding:"required"`
	}
	var requestBody RequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		fmt.Println(err)
		c.AbortWithStatus(400)
		return
	}
	accountIds := common.Unique(requestBody.AccountIds)
	result := orm.QueryAdSenseRevenueList(db, accountIds, requestBody.StartDate, requestBody.EndDate)
	c.JSON(200, common.JSON{"status": true, "data": result})
}

func GetAdSenseDomains(c *gin.Context) {

}