package main

import (
	"fmt"
	"sync"
	"time"
)

type Chopstick struct {
	sync.Mutex
}

type Philosopher struct {
	number                        int
	leftChopstick, rightChopstick *Chopstick
}

func (p Philosopher) eat(host chan struct{}, wg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		<-host

		p.leftChopstick.Lock()
		p.rightChopstick.Lock()
		fmt.Println("starting to eat", p.number)

		time.Sleep(500)

		p.rightChopstick.Unlock()
		p.leftChopstick.Unlock()
		fmt.Println("finishing eating", p.number)

		host <- struct{}{}
	}

	wg.Done()
}

func main() {
	chopsticks := make([]*Chopstick, 5)
	for i := range chopsticks {
		chopsticks[i] = &Chopstick{}
	}

	philosophers := make([]*Philosopher, 5)
	for i := range philosophers {
		philosophers[i] = &Philosopher{
			number:         i + 1,
			leftChopstick:  chopsticks[i],
			rightChopstick: chopsticks[(i+1)%5],
		}
	}

	host := make(chan struct{}, 2)

	var wg sync.WaitGroup
	wg.Add(5)

	for i := 0; i < 2; i++ {
		host <- struct{}{}
	}

	// Start the philosophers
	for _, p := range philosophers {
		go p.eat(host, &wg)
	}

	wg.Wait()
}
