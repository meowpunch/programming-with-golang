package main

import "fmt"

type Server struct {
	host string
}

func NewConcreteServer(host string) Server {
	// Smelly as unexported type pointer in interface
	return Server{host}
}

func (s *Server) Start() error {
	fmt.Println("start server")
	return nil
}

func (s *Server) Stop() error {
	fmt.Println("stop server")
	return nil
}

func (s *Server) Wait() error {
	fmt.Println("wait server")
	return nil
}

func main() {
	server := NewConcreteServer("localhost")

	server.Start()
	server.Stop()
	server.Wait()
}
