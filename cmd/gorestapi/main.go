package main

import (
	"github.com/dsoloview/gorestapi/internal/app/handler"
	"github.com/dsoloview/gorestapi/internal/app/http/controller"
	"github.com/dsoloview/gorestapi/internal/app/repository"
	"github.com/dsoloview/gorestapi/internal/app/service"
	"github.com/dsoloview/gorestapi/internal/database"
	"github.com/dsoloview/gorestapi/internal/server"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatalf("error init config %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env %s", err.Error())
	}

	db, err := database.NewPostgresDB()
	if err != nil {
		logrus.Fatal(err)
	}

	repositories := repository.NewRepository(db)
	services := service.NewService(repositories)
	controllers := controller.NewController(services)
	handlers := handler.NewHandler(controllers)

	serv := new(server.Server)
	if err := serv.Run(viper.GetString("port"), handlers.Handle()); err != nil {
		logrus.Fatal(err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
