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

	total := superAdd(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	fmt.Println(total)

	fmt.Println(canIDrink(16))

	fmt.Println(canIDrinkSwitch(16))
}

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

/**
switch use if, else if, else if, else if, else if
*/
func canIDrinkSwitch(age int) bool {
	defer fmt.Println("I'm Done!")
	switch koreanAge := age + 2; {
	case koreanAge < 18:
		return false
	case koreanAge == 18:
		return true
	case koreanAge > 50:
		return false
	}
	return false
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

// range give array index
func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	//for i := 0; i < len(numbers); i++ {
	//	fmt.Println(numbers[i])
	//}
	return total
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}
