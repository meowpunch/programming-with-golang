package main

import "fmt"

// ServerI Immediately smells very very
// Interface should describe behaviours but ServerI is not a behaviour but a thing
type ServerI interface {
	Start() error
	Stop() error
	Wait() error
}

type server struct {
	host string
}

// NewServer should return concrete type as a factory function.
// If caller should deal with decoupling if it's needed
func NewServer(host string) ServerI {
	// Smelly as unexported type pointer in interface
	return &server{host}
}

func (s *server) Start() error {
	fmt.Println("start server")
	return nil
}

func (s *server) Stop() error {
	fmt.Println("stop server")
	return nil
}

func (s *server) Wait() error {
	fmt.Println("wait server")
	return nil
}

func main() {
	server := NewServer("localhost")

	server.Start()
	server.Stop()
	server.Wait()
}
