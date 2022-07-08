package main

import (
	"awesomeProject/something"
	"fmt"
	"strings"
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

	fmt.Println(multiply(3, 2))

	// _ is ignore varaible
	totalLength, upperName := lenAndUpper("pedro")
	length, _ := lenAndUpper("pedro")

	fmt.Println(totalLength, upperName, length)

	repeatMe("nico", "pedro", "marco", "hello")

	fmt.Println(lenAndUpperNaked("pedro"))
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func lenAndUpperNaked(name string) (length int, upperName string) {
	defer fmt.Println("I'm done")

	length = len(name)
	upperName = strings.ToUpper(name)
	return
}

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}
