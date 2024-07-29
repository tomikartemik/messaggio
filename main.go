package main

import (
	"fmt"
	"messaggio/pkg/handler"
	"messaggio/pkg/repository"
	"messaggio/pkg/service"
	"messaggio/pkg/utils"
	"time"
)

func main() {

	cfg := utils.GetConfig()
	db, err := utils.NewPostgresDB(cfg)

	if err != nil {
		fmt.Println("failed to init DB: %s", err.Error())
	}

	time.Sleep(10 * time.Second)
	kafkaWriter := utils.InitKafkaWriter()
	//kafkaReader := utils.InitKafkaReader()
	utils.InitKafkaTopic()

	repos := repository.NewRepository(db)
	services := service.NewService(repos, kafkaWriter)
	handlers := handler.NewHandler(services)

	srv := new(utils.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		fmt.Println("error occured while running server %s", err.Error())
	}
}
