package main

import "fmt"

func PrintMaps() {
	prices := map[string]float64{
		"testing": 0.123,
		"bread":   2.40,
	}

	fmt.Println("Printing an entire map:")
	fmt.Println(prices)
	fmt.Println("Printing a single item:")
	fmt.Println(prices["testing"])
}

func LoopMaps() {
	alphabet := map[string]int{
		"a": 0,
		"b": 1,
		"c": 2,
		"d": 3,
	}

	for key, value := range alphabet {
		fmt.Printf("Key: %s, Value: %d \n", key, value)
	}
}
