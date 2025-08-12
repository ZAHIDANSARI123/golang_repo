package main

import "fmt"

func main() {
	age := 34
	name := "Alice"
	height := 4.89090

	fmt.Println("age: ", age, "height: ", height, "name: ", name)

	fmt.Println("hello")

	// fmt.Printf("age is %d\n", age)
	// fmt.Printf("height is %.2f\n", height)
	// fmt.Printf("Type of age is %T\n", age)
	// fmt.Printf("Type of height is %T\n", height)

	fmt.Printf("age is %d\n", age)
	fmt.Printf("Name: %s, Age: %d, Height: %.2f\n", name, age, height)
}
