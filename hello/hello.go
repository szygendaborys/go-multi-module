package main

import (
	"fmt"
	"sort"
	"strings"

	"example.com/initials"
	"golang.org/x/example/hello/reverse"
)

func main() {
	fmt.Println("I have been called")
	fmt.Println(reverse.String("Hello"))
	fmt.Println(reverse.Int(123))

	test := "TEST"
	var test2 int = 123
	fmt.Println(test, test2)

	// Formatting strings
	fmt.Printf("Testing the string %q test lol, %T \n", test, test2)

	// Save formatted string
	output := fmt.Sprintf("Testing the string %q test lol, %T \n", test, test2)
	fmt.Println(output)

	StandardLibrary()
	Mathematical()
	Loops()
	PrintInitials()
	PrintVariableFromAnotherModule()
	PrintMaps()
	LoopMaps()
}

func StandardLibrary() {
	greeting := "Hello there!"

	fmt.Println(strings.Contains(greeting, "ed"))
	fmt.Println(greeting)
	fmt.Println(strings.ToUpper(greeting))
}

func Mathematical() {
	ages := []int{12323, 3213213, 123123, 3, 1, 231, 4324}

	sort.Ints(ages)

	fmt.Println(ages)

	index := sort.SearchInts(ages, 99999999999)
	index2 := sort.SearchInts(ages, 3)
	fmt.Println(index)
	fmt.Println(index2)
}

func Loops() {
	iterator := 0
	numbers := []int{12323, 3213213, 123123, 3, 1, 231, 4324}

	for iterator < 5 {
		fmt.Printf("The value for iterator is %d \n", iterator)
		iterator++
	}

	sort.Ints(numbers)
	fmt.Println("Iteration over an array of numbers")

	for i := 0; i < len(numbers); i++ {
		fmt.Printf("Iterated over num %d \n", numbers[i])
	}

	fmt.Println("Another way to iterate over array elements")

	// Another way to iterate over a slice of numbers
	for j, value := range numbers {
		fmt.Printf("Index: %d, and the value: %d \n", j, value)
	}
}

func PrintInitials() {
	name := "Borys Szygenda"
	nameInitial, surnameInitial, err := initials.GetInitials(name)

	if err != nil {
		fmt.Printf("An error occurred: %s \n", err)
		return
	}

	fmt.Printf("Successfully generated initials: %s%s \n", nameInitial, surnameInitial)
}

func PrintVariableFromAnotherModule() {
	fmt.Println("Printing a variable outside of the file scope: ")
	fmt.Println(testing)
}
