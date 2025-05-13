package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/patrickdevbr-portfolio/cms/libs/go-common/logger"
)

func Init() {
	engine := gin.Default()

	v1 := engine.Group("/v1")

	v1.GET("/", func(ctx *gin.Context) {
		logger.Debug()
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	engine.Run(":8080")
}
