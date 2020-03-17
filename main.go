package main

import (
	"DataApi.Go/api"
	"DataApi.Go/database"
	"DataApi.Go/middleware"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	dbConfig := os.Getenv("DB_CONFIG")
	db, _ := database.Initialize(dbConfig)
	port := os.Getenv("PORT")
	app := gin.Default()
	app.Use(middleware.LoggerToFile())
	app.Use(database.Inject(db))
	app.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{
			"message": "hello " + name,
		})
	})
	api.ApplyRoutes(app) // apply api router
	err2 := app.Run(":" + port)
	if err2 != nil {
		panic(err2)
	}
}
