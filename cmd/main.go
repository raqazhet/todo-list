package main

import (
	"log"

	"todolist/pkg/handler"
	"todolist/pkg/repository"
	"todolist/pkg/service"

	"todolist"

	"github.com/spf13/viper"
)

func main() {
	if err := InitConfig(); err != nil {
		log.Printf("initConfig err: %v", err)
		return
	}
	repos := repository.NewRepository()
	service := service.NewService(*repos)
	handlers := handler.NewHandler(service)
	srv := new(todolist.Server)
	if err := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %v", err)
	}
}

func InitConfig() error {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
