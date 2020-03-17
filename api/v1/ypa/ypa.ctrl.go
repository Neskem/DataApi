package ypa

import (
	"DataApi.Go/database/orm"
	"DataApi.Go/lib/common"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
)

func getDailyYPA(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	startDate, _ := strconv.Atoi(c.Query("start_date"))
	endDate, _ := strconv.Atoi(c.Query("end_date"))

	betweenDates := common.GetBetweenDays(startDate, endDate, true)

	result := orm.QueryDailyYpaList(db, betweenDates)
	response := common.JSON{
		"status": true,
		"data": result,
	}
	c.JSON(200, response)
}
