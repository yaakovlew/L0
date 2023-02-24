package nats_chanel

import (
	"github.com/nats-io/stan.go"
	"strconv"
	"time"
)

func Produce() {

	sc, err := stan.Connect("prod", "simple-pub", stan.NatsURL("nats://localhost:4223"))
	if err != nil {
		panic(err)
	}
	defer sc.Close()
	for i := 1; ; i++ {
		sc.Publish("bestellungen", []byte("Bestellung "+strconv.Itoa(i)))
		time.Sleep(2 * time.Second)
	}
}
