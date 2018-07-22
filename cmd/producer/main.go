package main

import (
	"flag"
	"github.com/augurysys/go-franz"
	"time"
)

func main() {

	brokers := flag.String("brokers", "localhost:29092", "Broker url")
	topic := flag.String("topic", "", "Topic to produce to")

	message := flag.String("message","", "Send this one message and exit")

	flag.Parse()

	cfg := franz.NewConfig()

	cfg.Brokers = *brokers

	prod, err := franz.NewProducer(cfg)
	if err != nil {
		panic(err)
	}

	defer prod.Close()

	prod.Produce([]byte("key"), []byte(*message), time.Now(), *topic)

}
