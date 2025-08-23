package main

import "fmt"

func main() {

	// numbers := []int{1, 2, 3, 4, 5}
	// numbers = append(numbers, 3, 10,11,12,14,15,16,17)  // add the two array data
	// fmt.Println("Number is :", numbers)
	// fmt.Printf("Number has data type : %T\n", numbers)
	// fmt.Println("Length : ", len(numbers))


	// name := [] int {}
	// fmt.Println("name : ", name)


	numbers := make([]int,  3,  5)

	numbers = append(numbers, 4)
	numbers = append(numbers, 88)
	numbers = append(numbers, 6)

	fmt.Println("Slice : ", numbers)
	fmt.Println("Length:", len(numbers))
	fmt.Println("Capacity:", cap(numbers))


	stock := make([]int, 0)
	fmt.Println("Slice:", stock)
	fmt.Println("Length:", len(stock))
	fmt.Println("Capacity:", cap(stock))
}