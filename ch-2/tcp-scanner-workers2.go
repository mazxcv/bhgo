package ch2

import (
	"fmt"
	"net"
	"sort"
)

// worker2 runs as a goroutine and takes a channel of ports to scan and a result channel to send open ports to.
// It attempts to make a TCP connection to each port and sends the port number to the results channel if successful.
// If an error occurs, it sends 0 to the results channel and continues to the next port.
func worker2(ports chan int, result chan int) {
	for p := range ports {
		address := net.JoinHostPort(address, fmt.Sprintf("%d", p))
		conn, err := net.Dial("tcp", address)
		if err != nil {
			fmt.Printf("Error connecting to %d: %v\n", p, err)
			result <- 0
			continue
		}
		conn.Close()
		fmt.Printf("Port %d is open\n", p)
		result <- p
	}
}

const (
	workerCount = 500
	portCount   = 65535
	address     = "localhost"
)

// AsyncTCPScanWithWorkers2 performs a TCP port scan on a range of ports using goroutines and channels.
// It uses a worker pool to concurrently check each port on the target host.
// Open ports are collected and sorted before being printed to the console.
func AsyncTCPScanWithWorkers2() {
	// Create a buffered channel for ports and an unbuffered channel for results
	ports := make(chan int, workerCount)
	results := make(chan int)
	var openPorts []int

	fmt.Println("Starting goroutines")
	// Launch worker goroutines
	for range cap(ports) {
		go worker2(ports, results)
	}

	fmt.Println("Starting port sending goroutine")
	// Goroutine to send port numbers to the ports channel
	go func() {
		for i := 1; i <= portCount; i++ {
			ports <- i
		}
	}()

	fmt.Println("Starting result receiver loop")
	// Receive results from the results channel
	for range portCount {
		port := <-results
		if port != 0 {
			openPorts = append(openPorts, port)
		}
	}

	// Close channels
	close(ports)
	close(results)

	// Sort open ports and print them
	sort.Ints(openPorts)
	fmt.Println("Printing open ports")
	fmt.Println(openPorts)
}
