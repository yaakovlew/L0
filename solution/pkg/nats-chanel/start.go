package nats_chanel

import (
	"fmt"
	stan "github.com/nats-io/stan.go"
)

const (
	port    = "4222"
	cluster = "firstname"
	client  = "secondname"
)

func Message() {
	sc, err := stan.Connect(cluster, client, stan.NatsURL("nats://localhost:"+port))
	if err != nil {
		panic(err)
	}
	defer sc.Close()
	sc.Publish("foo", []byte("Hello World")) // does not return until an ack has been received from NATS Streaming
	sub, _ := sc.Subscribe("foo", func(m *stan.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})
	sub.Unsubscribe()
}
