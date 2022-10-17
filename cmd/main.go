package main

import (
	"log"
	"restapi"
	"restapi/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(restapi.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("running error: %s", err.Error())
	}

}
