package main

import (
	"fmt"
	"log"
	"template-gin/config"
	"template-gin/router"
	"template-gin/server"
)

func main() {
	cfg, err := config.Setup(".", "config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	srv := server.NewServer()

	router.SetupRouter(srv)

	fmt.Printf("server running on port %v\n", cfg.Port)
	if err := srv.Run(fmt.Sprintf(":%s", cfg.Port)); err != nil {
		log.Fatal(err)
	}
}
