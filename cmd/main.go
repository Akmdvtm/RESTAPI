package main

import (
	todoApp "github.com/akmdvtm/todo-app"
	"github.com/akmdvtm/todo-app/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todoApp.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("running error: %s", err.Error())
	}

}
