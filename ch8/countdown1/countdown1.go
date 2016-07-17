package main

import (
	"fmt"
	"time"
	"os"
)

func main() {
	fmt.Println("Commencing countdown.")
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()

	ticker := time.NewTicker(1 * time.Second)

	for countdown := 1; countdown <= 10; countdown++ {
		fmt.Println(countdown)
		select {
		case <- ticker.C:
		case <-abort:
			fmt.Println("Aborting...")
			ticker.Stop()
			return
		}
	}

	fmt.Println("Launching...")
}


