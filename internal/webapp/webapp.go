package webapp

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/leogsouza/running-table/internal/db"
)

func StartServer(database *db.Database) {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/results", func(c *gin.Context) {
		results := database.GetRecords()
		c.JSON(http.StatusOK, gin.H{
			"results": results,
		})
	})
	r.POST("/results", func(c *gin.Context) {
		var json db.Record
		if err := c.BindJSON(&json); err == nil {
			database.AddRecord(json)
			c.JSON(http.StatusCreated, json)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{})
		}
	})
	r.Run()
}
