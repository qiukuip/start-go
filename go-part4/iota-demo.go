package main

import "fmt"

const (
	READ int = 1 << iota
	WRITE
	EXECUTE
)

func main() {
	fmt.Println("READ: ", READ)
	fmt.Println("WRITE: ", WRITE)
	fmt.Println("EXECUTE: ", EXECUTE)
}

