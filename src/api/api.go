package api

import (
	"context"
	"fmt"
	"hospi_bed_stats/db"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/config"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gin-gonic/gin"
)

func Api() {
	server := gin.Default()
	server.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "PONG!",
		})
	})

	server.GET("/check-table", func(c *gin.Context) {
		config, error := config.LoadDefaultConfig(context.TODO())
		if error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "config not found",
			})
		}
		tableInfo := db.TableInfo{DBClient: dynamodb.NewFromConfig(config), TableName: "test-my-table"}
		result, err := tableInfo.CheckTableExists(tableInfo.TableName)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "table not found!",
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"message": result.Table.TableName,
		})
	})

	server.GET("/beds", func(c *gin.Context) {
		fmt.Print("Hello....")
		config, error := config.LoadDefaultConfig(context.TODO())
		fmt.Print("config\n")
		fmt.Printf("%v", config)
		if error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "config not found",
			})
		}
		tableInfo := db.TableInfo{DBClient: dynamodb.NewFromConfig(config), TableName: "hms_bed_stat_svc"}
		result, err := tableInfo.GetAllBeds()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "data not found!",
			})
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
		}
		bedTypeId := c.Param("id")
		tableInfo := db.TableInfo{DBClient: dynamodb.NewFromConfig(config), TableName: "hms_bed_stat_svc"}
		result, err := tableInfo.GetBedDetails(bedTypeId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "bed_type_id mismatch",
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"data": result,
		})
	})
	server.Run()
}
