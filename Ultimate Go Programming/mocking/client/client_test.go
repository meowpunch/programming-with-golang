package main

// mock implements PubSub interface
type mock struct{}

func (m *mock) Publish(key string, value interface{}) error {
	return nil
}

func (m *mock) Subscribe(key string) error {
	return nil
}
