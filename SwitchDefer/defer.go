package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
	defer fmt.Println("world2")
	defer fmt.Println("world3")
}
