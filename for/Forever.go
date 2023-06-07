package main

import "fmt"
import "math/rand"

func main() {
	for i:=5; i<10; i++ {
	fmt.Println("My favorite number is", rand.Intn(i))
	}
}