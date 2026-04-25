package main

import "fmt"

func main() {
	var ptr1 *int

	var num = 42
	ptr1 = &num

	fmt.Println(ptr1)

	
	ptr3 := new(int)
	*ptr3 = 100

	fmt.Printf("ptr1 指向的值: %d\n", *ptr1)
	fmt.Printf("ptr3 指向的值: %d\n", *ptr3)
}

