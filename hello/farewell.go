package main

import "fmt"

const testing string = "Hey, I am a variable outside of the module scope"

func SayGoodbye(name string) {
	fmt.Printf("Goodbye %s! \n", name)
}
