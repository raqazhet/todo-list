package main

import (
	"log"

	"todolist/pkg/handler"
	"todolist/pkg/repository"
	"todolist/pkg/service"

	"todolist"
)

func main() {
	repos := repository.NewRepository()
	service := service.NewService(*repos)
	handlers := handler.NewHandler(service)
	srv := new(todolist.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %v", err)
	}
}
