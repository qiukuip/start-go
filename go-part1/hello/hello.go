package main

import (
	"fmt"
	"day.happy365/greetings"
)

func main() {
	fmt.Println("Hellow world!")

	message := greetings.Hello("Gladys")
	fmt.Println(message)
}

