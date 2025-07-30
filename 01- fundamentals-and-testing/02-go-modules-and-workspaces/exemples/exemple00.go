package main

import "fmt"

func PrintValue[T any](value T) {
	fmt.Println(value)
}

func main() {
	//print string
	PrintValue("String")
	//print integer
	PrintValue(55)
	//print double
	PrintValue(1.1)
}
