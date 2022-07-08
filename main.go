package main

import (
	"awesomeProject/something"
	"fmt"
)

/**
1. look for package main if you build .go
2. look for func main
*/
func main() {
	fmt.Println("Hello, World!")
	something.SayHello()

	const name string = "park sang ho"
	//name = "sang ho" --> error
	fmt.Println(name)

	var varName string = "parksangho"
	varName = "sangho"
	fmt.Println(varName)

	insideName := "psh"

	fmt.Println(insideName)
}
