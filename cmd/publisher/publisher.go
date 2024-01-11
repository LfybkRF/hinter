package main

import (
	"l0_ms/internal/service"
	"l0_ms/internal/config"
	"log"
)


func main() {
	configPath := "./config/publisher.config.yaml"
	configure, err := config.NewConfig(configPath)

	if err != nil {
		log.Fatalf("failed to Config %v", err)
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

	natsStreaming := service.NewPublisher(configure.NatsStreaming.Channel, natsClient)
	natsStreaming.NatsSteamingSubscribe()
}