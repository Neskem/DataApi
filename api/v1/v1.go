package v1

import (
	"DataApi.Go/api/v1/adSense"
	"DataApi.Go/api/v1/spark"
	"DataApi.Go/api/v1/yna"
	"DataApi.Go/api/v1/ypa"
	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.GET("/ping", ping)
		spark.ApplyRoutes(v1)
		ypa.ApplyRoutes(v1)
		yna.ApplyRoutes(v1)
		adSense.ApplyRoutes(v1)
	}
}
