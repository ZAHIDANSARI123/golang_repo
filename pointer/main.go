package main 


import "fmt"

func modifyValueByReferenceBy(num *int) {
	*num = *num + 40
}

func main() {
	// var num int
	num := 2

	// pointer and address of operator
	// var ptr *int
	// ptr = &num
	ptr := &num

	fmt.Println("Num has value:", num)
	fmt.Println("Ptr has value:", ptr)

	var pointer *int  // default value of pointer is nil
	if pointer == nil {
		fmt.Println("Pointer is not assigned to any memory address")
	}

	value := 20
	modifyValueByReferenceBy(&value)
	fmt.Println("Value after modification:", value)
}
