package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"todolist/pkg/handler"
	"todolist/pkg/redisC"
	"todolist/pkg/repository"
	"todolist/pkg/service"

	"todolist"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
)

// @title  QR
// @version 1.0
// @description API server for todolist Application

// @host localhost:8000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

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
	casheRedis := redisC.NewCasheRedis()
	repos := repository.NewRepository(db)

	service := service.NewService(*repos, casheRedis)
	handlers := handler.NewHandler(service)
	srv := new(todolist.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %v", err)
		}
	}()
	logrus.Print("todoApp started")
	quite := make(chan os.Signal, 1)
	signal.Notify(quite, syscall.SIGINT, syscall.SIGINT)
	<-quite
	if err := srv.Shutdown(context.Background()); err != nil {
		return
	}
	logrus.Print("TodoApp shutting down")
}

func InitConfig() error {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
