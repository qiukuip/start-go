package main

import "fmt"

var globalVar = "全局变量"

func init() {
	fmt.Println("init1")
}

func init() {
	fmt.Println("init2")
}

func init3() {
	fmt.Println("init3")
}

func main() {
	fmt.Println("main函数执行")
	fmt.Printf("globaleVar: %+v\n", globalVar)
}

