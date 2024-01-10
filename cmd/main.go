package main

import (
	"l0_ms/internal/app/server"
	"l0_ms/internal/config"
	"l0_ms/internal/dao"
	"l0_ms/internal/service"
	"log"
	"os"
)


func main() {

	configPath := os.Getenv("CONFIG_PATH")

	configure, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatalf("failed to Config %v", err)
	}

	// Открытие соединения с postgres
	dbClient, err := config.DatabaseConnect(configure.Database)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	log.Printf("connecting to the database: success")

	defer func() {
		dbInstance, _ := dbClient.DB()
		if err = dbInstance.Close(); err != nil {
			log.Fatalf("failed to close database: %v", err)
		}
	}()

	// Клиент для реализации бизнес-логики
	client := dao.NewOrderClient(dbClient)
	err = client.Start()
	if err != nil {
		log.Fatalf("failed Start: %v", err)
	}

	natsClient, err := config.NatsStreamingConnect(configure.NatsStreaming)
	if err != nil {
		log.Fatalf("failed to connect to nats-streaming: %v", err)
	}

	log.Printf("connecting to the nats-streaming: success")

	defer func() {
		if err = natsClient.Close(); err != nil {
			log.Fatalf("failed to close nats-streaming: %v", err)
		}
	}()

	// Клиент для nats-streaming
	natsStreaming := service.NewSubscriber(configure.NatsStreaming.Channel, natsClient, client)

	natsStreaming.NatsSteamingSubscribe()

	// Запуск http сервера
	httpRouter := server.NewHttpRouter(configure.Server, client)
	log.Printf("start http-server")
	err = httpRouter.Start()
	if err != nil {
		log.Fatalf("failed to listen the port: %d %v", configure.Server.Port, err)
	}
}