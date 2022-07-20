package main

import (
	"github.com/dsoloview/gorestapi/internal/app/handler"
	"github.com/dsoloview/gorestapi/internal/app/server"
	"github.com/dsoloview/gorestapi/internal/app/service"
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

	services := service.NewService()
	handlers := handler.NewHandler(services)

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