package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type fork struct{ sync.Mutex }

type philosopher struct {
	leftFork, rightFork *fork
	number              int
}

const (
	count  = 5
	hunger = 3
	eat    = time.Second / 100
	think  = time.Second / 100
)

func (p philosopher) eat(wg *sync.WaitGroup, host chan bool) {
	sleep := func(delay time.Duration) {
		n := rand.Intn(10)
		time.Sleep(time.Duration(n) * delay)
	}

	for j := hunger; j > 0; j-- {
		host <- true

		p.leftFork.Lock()
		p.rightFork.Lock()

		fmt.Printf("starting to eat %d\n", p.number)
		sleep(eat)

		p.rightFork.Unlock()
		p.leftFork.Unlock()

		<-host

		fmt.Printf("finishing eating %d\n", p.number)
		sleep(think)
	}

	wg.Done()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var wg sync.WaitGroup
	var host = make(chan bool, 2)

	forks := make([]*fork, count)
	for i := 0; i < count; i++ {
		forks[i] = new(fork)
	}

	philosophers := make([]*philosopher, count)
	for i := 0; i < count; i++ {
		philosophers[i] = &philosopher{number: i, leftFork: forks[i], rightFork: forks[(i+1)%count]}

		wg.Add(1)
		go philosophers[i].eat(&wg, host)

	}

	wg.Wait()
}
