package main

import (
	"template-gin/router"
	"template-gin/server"
)

func main() {
	srv := server.NewServer()

	router.SetupRouter(srv)
	srv.Run()
}
