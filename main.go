package main

import (
	"fmt"
	"log"
	"template-gin/config"
	"template-gin/db"
	"template-gin/handler"
	"template-gin/repo"
	"template-gin/router"
	"template-gin/server"
	"template-gin/usecase"
)

func main() {
	cfg, err := config.Setup(".", "config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	srv := server.NewServer()

	mongodb, err := db.ConnectMongo(
		cfg.Mongo.URI,
		cfg.Mongo.Database,
	)
	if err != nil {
		log.Fatal(err)
	}

	repo := repo.NewMongoRepo(mongodb)
	uc := usecase.NewUsecase(repo)
	handler := handler.NewHandler(uc)

	router.SetupRouter(srv, handler)

	fmt.Printf("server running on port %v\n", cfg.App.Port)
	if err := srv.Run(fmt.Sprintf(":%s", cfg.App.Port)); err != nil {
		log.Fatal(err)
	}
}
