package ch2

import (
	"fmt"
	"net"
	"sync"
)

// main performs a TCP port scan on scanme.nmap.org.
// It is a basic example of how to use the net and sync packages
// to concurrently scan TCP ports.
// The function takes no arguments and returns nothing.
func AsyncTCPScan() {

	var wg sync.WaitGroup
	for port := 1; port <= 1024; port++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()

			// Construct the address string.
			address := fmt.Sprintf("scanme.nmap.org:%d", port)

			// Attempt to make a TCP connection.
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}

			// Close the connection.
			conn.Close()
			fmt.Printf("%d open\n", port)
		}(port)
	}

	// Wait for all goroutines to finish.
	wg.Wait()
}
