package api

import (
	"context"
	"hospi_bed_stats/db"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/config"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gin-gonic/gin"
)

func Api() {
	server := gin.Default()

	server.GET("/healthy", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK!!",
		})
	})

	server.GET("/beds", func(c *gin.Context) {
		config, error := config.LoadDefaultConfig(context.TODO())
		if error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "config not found",
			})
			return
		}
		tableInfo := db.TableInfo{DBClient: dynamodb.NewFromConfig(config), TableName: "hms_bed_stat_svc"}
		result, err := tableInfo.GetAllBeds()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "data not found!",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": result,
		})
	})

	server.GET("/beds/:id", func(c *gin.Context) {
		config, error := config.LoadDefaultConfig(context.TODO())
		if error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "config not found",
			})
			return
		}
		bedTypeId := c.Param("id")
		tableInfo := db.TableInfo{DBClient: dynamodb.NewFromConfig(config), TableName: "hms_bed_stat_svc"}
		result, err := tableInfo.GetBedDetails(bedTypeId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "bed_type_id mismatch",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": result,
		})
	})

	server.Run()
}
