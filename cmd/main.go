package main

import (
	"log"

	"todolist/pkg/handler"

	"todolist"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todolist.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %v", err)
	}
}
