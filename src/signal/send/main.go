package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s [pid]\n", os.Args[0])
		return
	}

	pid, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	process, err := os.FindProcess(pid)
	if err != nil {
		panic(err)
	}

	process.Signal(os.Kill)
	// you can also use
	// process.Kill()
}
