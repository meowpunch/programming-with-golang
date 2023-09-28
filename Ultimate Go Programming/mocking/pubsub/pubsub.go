package pubsub

import "fmt"

// PubSub assume that there is test in the package to spin up docker container with pub sub image
// In the package there is no more concrete implementation for PubSub. Interface is not needed.
// If client of API wanted to create unit test and mock PubSub struct, they can create interface in their package.
type PubSub struct {
	host string
}

func New(host string) *PubSub {
	return &PubSub{host}
}

func (ps PubSub) Publish(key string, value interface{}) error {
	fmt.Println("publish " + key)
	return nil
}

func (ps PubSub) Subscribe(key string) error {
	fmt.Println("subscribe " + key)
	return nil
}
