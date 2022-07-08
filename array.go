package main

import "fmt"

func main() {
	// array
	names := [5]string{"pedro", "marco", "nico"}
	names[3] = "max"
	names[4] = "jean"
	fmt.Println(names)

	// slice
	var namesSlice []string

	resultSlice := append(namesSlice, "pedro")

	fmt.Println(resultSlice)

}
