package service

import (
	"l0_ms/internal/dao"
	"log"

	"github.com/nats-io/stan.go"
)


type Subscriber struct {
	natsClient  stan.Conn
	channel		string
	client		*dao.Client
}


func NewSubscriber(channel string, natsClient stan.Conn, client *dao.Client) *Subscriber {

	natsSteamingClient := Subscriber{
		natsClient: natsClient,
		channel:    channel,
		client:     client,
	}
	return &natsSteamingClient
}

func (s *Subscriber) NatsSteamingSubscribe() {
	_, err := s.natsClient.Subscribe(
		s.channel, func(m *stan.Msg) {
			err := s.client.AddOrder(m.Data)
			if err != nil {
				log.Printf("Error order adding %v", err)
			}
		})

	if err != nil {
		log.Fatal(err)
	}
}