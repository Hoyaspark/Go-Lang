package main

import "fmt"

type person struct {
	name         string
	age          int
	favoriteFood []string
}

func main() {
	foods := []string{"kimchi", "meat"}
	nico := person{name: "pedro", age: 25, favoriteFood: foods}

	fmt.Println(nico)
}
