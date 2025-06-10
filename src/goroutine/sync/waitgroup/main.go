package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		fmt.Println("work1")
		wg.Done()
	}()

	go func() {
		fmt.Println("work2")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("all work done")
}
