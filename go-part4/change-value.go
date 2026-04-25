package main

import "fmt"

func changeValue(x int) {
	x += 100
}

func changePointer(ptr *int) {
	*ptr += 100
}

func main() {
	num := 10
	fmt.Printf("原始值: %d\n", num)
	changeValue(num)
	fmt.Printf("值传递: %d\n", num)
	changePointer(&num)
	fmt.Printf("指针传递: %d\n", num)
}

