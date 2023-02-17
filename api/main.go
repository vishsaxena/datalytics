package main

import "fmt"

func someFunc(str string) {
	fmt.Println(str)
}

func main() {
	go someFunc("First Method Call")
	go someFunc("Second Method Call")
	go someFunc("Third Method Call")
	fmt.Println("Hello Datalytics API...")
}