package service

import (
	"l0_ms/internal/models"

	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
	"math/rand"

	"github.com/nats-io/stan.go"
)

type Publisher struct {
	natsClient stan.Conn
	channel    string
}

func NewPublisher(channel string, natsClient stan.Conn) *Publisher {

	natsSteamingClient := Publisher{
		natsClient: natsClient,
		channel:    channel,
	}
	return &natsSteamingClient
}

func hex(i int, prefix bool) string {
	i64 := int64(i)

	if prefix {
			return "0x" + strconv.FormatInt(i64, 16) // base 16 for hexadecimal
	} else {
			return strconv.FormatInt(i64, 16) // base 16 for hexadecimal
	}
}

func (p *Publisher) NatsSteamingSubscribe() {
	file, err := os.ReadFile("./model.json")
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}

	var order models.Order

	err = json.Unmarshal(file, &order)
	if err != nil {
		log.Fatalf("failed Unmarshal: %v", err)
	}
	for {
		order.SmID = rand.Intn(1000 - 100) + 100
		order.DateCreated = time.Now()
		
		// Генерируем трек номер из 12 символов для возможности генерировать штрих код
		trackNumber := uint64(rand.Int63n(900_000_000_000)) + 100_000_000_000
		order.TrackNumber = strconv.FormatUint(trackNumber, 10)

		OrderUID := fmt.Sprintf("%s%stest", hex(order.SmID*23, false), hex(int(time.Now().Unix()), false))
		order.OrderUID = OrderUID

		fmt.Println("Send Order UID: ", order.OrderUID)

		bytes, err := json.Marshal(order)
		if err != nil {
			log.Println("ERROR: json.Marshal:", err)
			continue
		}

		err = p.natsClient.Publish(p.channel, bytes)
		if err != nil {
			log.Println("ERROR: conn.Publisher:", err)
			continue
		}
		time.Sleep(time.Second * 5)
	}
}
