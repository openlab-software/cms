package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/patrickdevbr-portfolio/cms/libs/go-common/auth"
	"github.com/patrickdevbr-portfolio/cms/libs/go-common/logger"
)

func Init() error {
	engine := gin.Default()
	oidcProvider, err := auth.NewOIDCProvider()
	if err != nil {
		return err
	}

	engine.Use(auth.NewAuthMiddleware(oidcProvider))

	v1 := engine.Group("/v1")

	v1.GET("/", func(ctx *gin.Context) {
		logger.Debug()
		user := ctx.MustGet("user").(string)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
			"user":    user,
		})
	})

	engine.Run(":8080")
	return nil
}
