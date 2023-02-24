package main

import (
	_ "github.com/lib/pq"
	nats_chanel "solution/pkg/nats-chanel"
)

func main() {
	//user := model.User{"Makar", "Chudra", 11, "mymail@yandex.ru"}
	//db.RefreshDb()
	nats_chanel.Produce()
	//go nats_chanel.Take()
	//handler.Handle()
}
