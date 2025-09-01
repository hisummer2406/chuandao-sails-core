package main

import "sync"

var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done()
	println("hello goroutine", i)
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go hello(i)
	}
	wg.Wait()
}
