package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/patrickdevbr-portfolio/cms/apps/content-service/internal/application/services"
	"github.com/patrickdevbr-portfolio/cms/apps/content-service/internal/infra/db/mongodb"
	"github.com/patrickdevbr-portfolio/cms/apps/content-service/internal/infra/rest"
	"github.com/patrickdevbr-portfolio/cms/libs/go-common/mongodatabase"
)

type application struct {
	config config
}

type config struct {
	addr string
}

func (app *application) run() error {
	godotenv.Load(".env.dev")

	mux := http.NewServeMux()

	// oidcProvider, err := auth.NewOIDCProvider()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// authMiddleware := auth.NewMiddleware(oidcProvider)

	mongoClient, err := mongodatabase.Connect(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	rest.NewPageRest(mux, &services.PageServiceImpl{
		PageRepository: &mongodb.PageRepository{
			Client: mongoClient,
		},
	})

	srv := &http.Server{
		Addr:    app.config.addr,
		Handler: mux,
	}

	return srv.ListenAndServe()
}
