package nats_chanel

import (
	"fmt"
	stan "github.com/nats-io/stan.go"
)

func Take() {
	sc, err := stan.Connect("prod", "sub-1", stan.NatsURL("nats://localhost:4223"))
	if err != nil {
		panic(err)
	}
	fmt.Println("hi")
	defer sc.Close()

	sc.Subscribe("bestellungen", func(m *stan.Msg) {
		fmt.Printf("Got: %s\n", string(m.Data))
	})

}
