package main

import "sync"

func main() {
	var wg sync.WaitGroup
	for i := 0;; i++{
		wg.Add(1)
		go func(i int) {
			println(i)
			wg.Done()
		} (i)
	}
	wg.Wait()
}
