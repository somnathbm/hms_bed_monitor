package api

import (
	"context"
	"errors"
	"fmt"
	"hospi_bed_stats/db"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

const appName = "hms-bm"

var (
	tracer = otel.Tracer(appName)
	logger = otelslog.NewLogger(appName)
)

func Api() {
	// Handle SIGINT (CTRL + C) properly
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Setup Opentelemetry
	otelShutdown, err := SetupOTelSDK(ctx)
	if err != nil {
		return
	}
	// Handle shutdown properly so that nothing leaks
	defer func() {
		err = errors.Join(err, otelShutdown(context.Background()))
	}()

	server := gin.Default()
	server.Use(otelgin.Middleware("hms-bm-svc"))

	// initialize the custom metrics
	// allCounterMetrics := metrics.GetAllCounterMetrics()

	server.GET("/bm/healthy", func(c *gin.Context) {
		// fire off the tracer
		ctx, span := tracer.Start(c.Request.Context(), "/bm/healthy", trace.WithAttributes(attribute.String("message", "OK!!")))
		defer span.End()

		// set logger
		logger.InfoContext(ctx, "service is healthy", "bm-logger", true)

		// no metrics
		c.JSON(http.StatusOK, gin.H{
			"message": "OK!!",
		})
	})

	server.GET("/bm/beds", func(c *gin.Context) {
		result, err := db.GetAllBeds()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"data":  "not found",
				"error": err.Error(),
			})
			return
		}
		// patientNum := len(result)

		// // fire off the tracer
		// ctx, span := tracer.Start(c.Request.Context(), "/pm/patients")
		// defer span.End()

		// // set log
		// logger.InfoContext(ctx, "patient count", "pm-logger", patientNum)

		// // set metrics
		// patientCountAttr := attribute.Int("patient.total", patientNum)
		// span.SetAttributes(patientCountAttr)
		// allMetrics["PatientCountMetric"].Record(ctx, 1, metric.WithAttributes(patientCountAttr))

		c.JSON(http.StatusOK, gin.H{
			"data": result,
		})
		// fire off the tracer
		// ctx, span := tracer.Start(c.Request.Context(), c.Request.RequestURI)
		// defer span.End()

		// // set logger
		// logger.InfoContext(ctx, "beds.total", "bm-logger", 1)
	})

	server.GET("/bm/beds/:id", func(c *gin.Context) {
		// ctx, span := tracer.Start(c.Request.Context(), c.Request.RequestURI)
		// defer span.End()

		bedTypeId := c.Param("id")
		result, err := db.GetBedDetails(bedTypeId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"data":  "not found",
				"error": err.Error(),
			})
			return
		}

		// set log if the DB operation succeeds
		// logger.InfoContext(ctx, "patient.info", "pm-logger", result)
		// metrics - no needed

		c.JSON(http.StatusOK, gin.H{
			"data": result,
		})

		// // fire off the tracer
		// ctx, span := tracer.Start(c.Request.Context(), c.Request.RequestURI)
		// defer span.End()

		// config, error := config.LoadDefaultConfig(context.TODO())
		// if error != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{
		// 		"error": "config not found",
		// 	})
		// 	return
		// }
		// bedTypeId := c.Param("id")
		// tableInfo := db.TableInfo{DBClient: dynamodb.NewFromConfig(config), TableName: "hms_bed_stat_svc"}
		// result, err := tableInfo.GetBedDetails(bedTypeId)
		// if err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{
		// 		"error": "bed_type_id mismatch",
		// 	})
		// 	return
		// }
		// // set logger
		// logger.InfoContext(ctx, "beds.type", "bm-logger", 1)

		// c.JSON(http.StatusOK, gin.H{
		// 	"data": result,
		// })
	})

	server.Run()

	// Wait for interruption
	select {
	case <-ctx.Done():
		stop()
		fmt.Println("signal received")
	}
}
