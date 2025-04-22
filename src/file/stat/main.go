package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("%s [exec file name]", os.Args[0])
		os.Exit(1)
	}
	info, err := os.Stat(os.Args[1])
	if err == os.ErrNotExist {
		fmt.Printf("file not found: %s\n", os.Args[1])
	} else if err != nil {
		panic(err)
	}
	fmt.Println("FileInfo:")
	fmt.Printf("Name: %s\n", info.Name())
	fmt.Printf("Size: %d\n", info.Size())
	fmt.Printf("last modified: %s\n", info.ModTime())
	fmt.Printf("IsDir: %t\n", info.IsDir())
	fmt.Printf("Normal File?: %t\n", info.Mode().IsRegular())
	fmt.Printf("Unix access bits: %o\n", info.Mode().Perm())
	fmt.Printf("Mode text: %s\n", info.Mode().String())
}
