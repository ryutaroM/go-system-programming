package main

import (
	"fmt"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("Accept Ctrl+C for 10 seconds")
	time.Sleep(10 * time.Second)

	signal.Ignore(syscall.SIGINT, syscall.SIGHUP)

	fmt.Println("Ignore Ctrl+C for 10 seconds")
	time.Sleep(10 * time.Second)
}
