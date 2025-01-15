package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {
	fmt.Println("I have been called")
	fmt.Println(reverse.String("Hello"))
	fmt.Println(reverse.Int(123))
}
