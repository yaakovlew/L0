package nats_chanel

import (
	"fmt"
	stan "github.com/nats-io/stan.go"
)

func Message() {
	sc, err := stan.Connect("prod", "store", stan.NatsURL("nats://localhost:4222"))
	if err != nil {
		panic(err)
	}
	defer sc.Close()
	sc.Publish("foo", []byte("Hello World"))
	sub, _ := sc.Subscribe("foo", func(m *stan.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})
	sub.Unsubscribe()
}
