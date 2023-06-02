package main

import (
	"os"

	"todolist/pkg/handler"
	"todolist/pkg/repository"
	"todolist/pkg/service"

	"todolist"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := InitConfig(); err != nil {
		logrus.Printf("initConfig err: %v", err)
		return
	}
	if err := gotenv.Load(); err != nil {
		logrus.Printf("error in gotenv.Load: %v", err)
		return
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		UserName: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBname:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Printf("error in db: %v", err)
		return
	}
	repos := repository.NewRepository(db)
	service := service.NewService(*repos)
	handlers := handler.NewHandler(service)
	srv := new(todolist.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %v", err)
	}
}

func InitConfig() error {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
