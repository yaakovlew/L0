package nats_chanel

import (
	stan "github.com/nats-io/stan.go"
	"solution/pkg/db"
	"sync"
)

func Take() {
	sc, err := stan.Connect("prod", "sub-1", stan.NatsURL("nats://localhost:4223"))
	if err != nil {
		panic(err)
	}
	defer sc.Close()

	sc.Subscribe("bestellungen", func(m *stan.Msg) {
		db.AddInDB(m.Data)
	})
	wait := func() {
		w := sync.WaitGroup{}
		w.Add(1)
		w.Wait()
	}
	wait()
}
