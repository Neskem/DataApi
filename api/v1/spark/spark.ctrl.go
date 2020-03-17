package spark

import (
	"DataApi.Go/database/models"
	"DataApi.Go/database/models/PV"
	"DataApi.Go/database/orm"
	"DataApi.Go/lib/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
)

type StatPagePV = PV.StatPagePV
type Post = models.Post

func create(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	type RequestBody struct {
		Url string `json:"url" binding:"required"`
		Pv int `json:"pv" binding:"required"`
	}
	var requestBody RequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		fmt.Println(err)
		c.AbortWithStatus(400)
		return
	}

	stat := StatPagePV{Page_url: requestBody.Url, Pv: requestBody.Pv}
	db.NewRecord(stat)
	db.Create(&stat)
	c.JSON(200, stat.Serialize())
}

func read(c *gin.Context)  {
	db := c.MustGet("db").(*gorm.DB)
	type RequestBody struct {
		Url string `json:"url" binding:"required"`
	}
	var requestBody RequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		fmt.Println(err)
		c.AbortWithStatus(400)
		return
	}
	var stat StatPagePV
	if err := db.Set("gorm:auto_preload", true).Where("page_url = ?", requestBody.Url).First(&stat).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}

	c.JSON(200, stat.Serialize())
}

func readDailyPV(c *gin.Context) {
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

func readMonthlyPV(c *gin.Context) {
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

func readTotalPV(c *gin.Context) {
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

func getAuthorPV(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	startDate, _ := strconv.Atoi(c.Query("start_date"))
	endDate, _ := strconv.Atoi(c.Query("end_date"))
	author := c.Query("author")

	betweenDates := common.GetBetweenDays(startDate, endDate, false)
	result := orm.QueryAuthor(db, betweenDates, author)
	c.JSON(200, result)

}

func getHostPV(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	startDate, _ := strconv.Atoi(c.Query("start_date"))
	endDate, _ := strconv.Atoi(c.Query("end_date"))
	hostName := c.Query("hostname")

	betweenDates := common.GetBetweenDays(startDate, endDate, false)
	result := orm.QueryHost(db, betweenDates, hostName)
	c.JSON(200, result)
}
