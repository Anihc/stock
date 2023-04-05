package main

import "fmt"

func main() {

	const (
		a1 = iota
		a2 = 999
		a3 = iota
	)
	// var user string = "tom"
	var name, age, b = "tt", 20, true
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", name)
	fmt.Printf("%v\n", age)

	const PI float64 = 3.14

	fmt.Printf("age: %v\n", age)

	fmt.Printf("%v\n", PI)

	fmt.Printf("%v\n", a1)
	fmt.Printf("%v\n", a2)
	fmt.Printf("%v\n", a3)
}
