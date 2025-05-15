package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/patrickdevbr-portfolio/cms/apps/content-service/internal/application/services"
	"github.com/patrickdevbr-portfolio/cms/apps/content-service/internal/infra/db/mongodb"
	"github.com/patrickdevbr-portfolio/cms/apps/content-service/internal/infra/rest"
	"github.com/patrickdevbr-portfolio/cms/libs/go-common/mongodatabase"
)

func main() {
	router := http.NewServeMux()

	// oidcProvider, err := auth.NewOIDCProvider()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// authMiddleware := auth.NewMiddleware(oidcProvider)

	mongoClient, err := mongodatabase.Connect(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	rest.NewPageRest(router, &services.PageServiceImpl{
		PageRepository: &mongodb.PageRepository{
			Client: mongoClient,
		},
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	server.ListenAndServe()
}
