package main

import "fmt"

func printClass(grade int) {
	switch {
	case grade > 5:
		fmt.Println("Great!")
	case grade <= 5 && grade > 3:
		fmt.Println("Well")
	default:
		fmt.Println("Sad!")
	}
}

func main() {
	var myGrade = 5
	printClass(myGrade)
}

