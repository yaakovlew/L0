package nats_chanel

import (
	"encoding/json"
	"github.com/nats-io/stan.go"
	"solution/pkg/model"
	"time"
)

func Produce() {

	sc, err := stan.Connect("prod", "simple-pub", stan.NatsURL("nats://localhost:4223"))
	if err != nil {
		panic(err)
	}
	defer sc.Close()
	time.Sleep(3 * time.Second)
	user := model.User{"Andre", "Pirlo", "Apirlo@mail.ru"}
	b, _ := json.Marshal(user)
	sc.Publish("bestellungen", b)
	testCount := make(map[string]string)
	testCount["FirstName"] = "Pedro"
	testCount["SecondName"] = "Bonucci"
	b, _ = json.Marshal(testCount)
	time.Sleep(3 * time.Second)
	sc.Publish("bestellungen", b)
	/*for i := 1; ; i++ {
		time.Sleep(3 * time.Second)
		sc.Publish("bestellungen", []byte("Bestellung "+strconv.Itoa(i)))
	}*/
}
