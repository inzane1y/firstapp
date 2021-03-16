package main

import "fmt"

func main() {
	// Sending channel
	c := make(chan int)
	// Quit channel
	q := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			// Get values from the channel 5 times
			fmt.Println(<-c)
		}

		// Break the channel
		q <- 0
	}()

	fib(c, q)
}

// Fibonacci sequence (sending int channel, quit int channel)
func fib(c, q chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x: // Send x into the channel
			x, y = y, x+y // Reevaluate variables
		case <-q: // Close function if recieved a value from quit-channel
			return
		}
	}
}
