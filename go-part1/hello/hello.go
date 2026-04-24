package main

import (
	"fmt"
	"day.happy365/greetings"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

  message, err := greetings.Hello("Gladys")
  // message, err := greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)

	messages, err := greetings.Hellos([]string{"Alice", "Bob", "Lucy"})
	if err != nil {
		log.Fatal(err)
	}
	for _, message := range messages {
		fmt.Println(message)
	}
}

