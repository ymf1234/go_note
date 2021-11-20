package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()

	logger.Warn("warning test")

	engine := gin.New()

	fmt.Println(engine)
}
