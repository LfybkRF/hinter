package main


import (
	stan "github.com/nats-io/stan.go"
	"log"
)


func main() {
	log.Printf("Starting Gateway\n")
	s, err := stan.Connect("demo", "demo")
	if err!= nil {
        log.Fatalf("Error: %s\n", err)
    }
	defer s.Close()
	log.Printf("Gateway started\n")
	s.Subscribe("demo.>", func(msg *stan.Msg) {
		log.Printf("Received message: %s\n", msg.Data)
    })
	
}