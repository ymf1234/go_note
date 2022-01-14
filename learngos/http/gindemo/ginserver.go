package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"math/rand"
	"time"
)

func main() {
	r := gin.Default()

	logger, err := zap.NewProduction()

	if err != nil {
		panic(err)
	}
	// 中间件
	r.Use(func(context *gin.Context) {
		s := time.Now()

		context.Next()

		logger.Info("incoming request",
			zap.String("path", context.Request.URL.Path),
			zap.Int("status", context.Writer.Status()),
			zap.Duration("elapsed", time.Now().Sub(s)),
		)
	}, func(context *gin.Context) {
		context.Set("requestId", rand.Int())
		context.Next()
	})


	r.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "hello",
		})
	})
	err = r.Run()
}
