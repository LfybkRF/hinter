package config


import (
	"fmt"
	"github.com/nats-io/stan.go"
)

type NatsStreamingConfig struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	ClusterID string `yaml:"cluster_id"`
	ClientID  string `yaml:"client_id"`
	Channel   string `yaml:"channel"`
}

func NatsStreamingConnect(config NatsStreamingConfig) (stan.Conn, error) {
	host := fmt.Sprintf("%s:%d", config.Host, config.Port)

	sc, err := stan.Connect(config.ClusterID+"", config.ClientID, stan.NatsURL(host))
	if err != nil {
		return nil, err
	}
	return sc, nil
}