package ch2

import (
	"fmt"
	"sync"
)

func worker(ports chan int, wg *sync.WaitGroup) {

	for port := range ports {
		fmt.Println(port)
	}
	wg.Done()
}

func AsyncTCPScanWithWorkers() {
	ports := make(chan int, 100)
	var wg sync.WaitGroup

	// Start a goroutine for each item in the cap(ports)
	for range cap(ports) {
		go worker(ports, &wg)
	}

	// Start sending items from 1 to 1024 to the ports channel
	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		ports <- i
	}

	// Close the ports channel
	close(ports)
}
