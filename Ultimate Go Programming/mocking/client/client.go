package main

import "pubsub"

type PubSub interface {
	Publish(key string, value interface{}) error
	Subscribe(key string) error
}

type Data struct {
	Line string
}

func main() {
	ps := pubsub.New("localhost")

	ps.Publish("key", Data{"data1"})
	ps.Subscribe("key")
}
