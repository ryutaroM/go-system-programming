package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var count int
	pool := sync.Pool{
		New: func() interface{} {
			count++
			return fmt.Sprintf("created %d", count)
		},
	}

	pool.Put("manual put 1")
	pool.Put("manual put 2")
	// pool.Put("manual put 3")
	// fmt.Println(pool.Get())
	// fmt.Println(pool.Get())
	runtime.GC() //good bye
	fmt.Println(pool.Get())
}
