package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/patrickdevbr-portfolio/cms/apps/content-service/internal/application/services"
	"github.com/patrickdevbr-portfolio/cms/apps/content-service/internal/infra/db/mongodb"
	"github.com/patrickdevbr-portfolio/cms/apps/content-service/internal/infra/rest"
	"github.com/patrickdevbr-portfolio/cms/libs/go-common/auth"
	"github.com/patrickdevbr-portfolio/cms/libs/go-common/mongodatabase"
)

func main() {
	engine := gin.Default()

	oidcProvider, err := auth.NewOIDCProvider()
	if err != nil {
		fmt.Println(err)
	}
	engine.Use(auth.NewMiddleware(oidcProvider))

	mongoClient, err := mongodatabase.Connect(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	v1 := engine.Group("/v1")

	pages := v1.Group("/pages")
	rest.NewPageRest(pages, &services.PageServiceImpl{
		PageRepository: &mongodb.PageRepository{
			Client: mongoClient,
		},
	})

	engine.Run(":8080")

}
