package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/patrickdevbr-portfolio/cms/apps/content-service/internal/application/services"
	"github.com/patrickdevbr-portfolio/cms/apps/content-service/internal/infra/amqpevent"
	"github.com/patrickdevbr-portfolio/cms/apps/content-service/internal/infra/db/mongodb"
	"github.com/patrickdevbr-portfolio/cms/apps/content-service/internal/infra/rest"
	"github.com/patrickdevbr-portfolio/cms/libs/go-common/mongodatabase"
	"github.com/patrickdevbr-portfolio/cms/libs/go-common/rabbitmq"
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

	ctx := context.Background()
	mongoClient, err := mongodatabase.Connect(ctx)
	if err != nil {
		fmt.Println(err)
	}
	defer mongoClient.Disconnect(ctx)

	rabbitMQPublisher, err := rabbitmq.NewRabbitMQPublisher()
	if err != nil {
		fmt.Println(err)
	}
	defer rabbitMQPublisher.Close()

	eventPublisher := amqpevent.NewRabbitMQEventPublisher(rabbitMQPublisher)
	pageRepo := mongodb.NewPageRepository(mongoClient)
	pageSvc := services.NewPageService(&pageRepo, eventPublisher)

	rest.NewPageRest(mux, pageSvc)

	srv := &http.Server{
		Addr:    app.config.addr,
		Handler: mux,
	}

	return srv.ListenAndServe()
}
