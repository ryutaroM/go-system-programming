package main

import (
	"fmt"
	"time"
)

func main() {
	tasks := []string{
		"cmake ..",
		"cmake . --build Release",
		"cpack",
	}
	for _, t := range tasks {
		go func() {
			fmt.Println(t)
		}()
	}
	time.Sleep(time.Second)
}
