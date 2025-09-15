package main

import "fmt"

func main() {
	// map creation using make function

	studentGrades := make(map[string]int)

	studentGrades["Alice"] = 100
	studentGrades["Bob"] = 90
	studentGrades["Charlie"] = 92
	studentGrades["Diana"] = 78

	fmt.Println("Mark of Alice is :", studentGrades["Alice"])
	studentGrades["Alice"] = 88
	fmt.Println("Updated Mark of Alice is :", studentGrades["Alice"])

	delete(studentGrades, "Alice")
	fmt.Println("Mark of Alice is :", studentGrades["Alice"])

	// checking	 if a key exists
	Grades, Exists := studentGrades["Alice"]
	fmt.Println("Mark of Alice is :", Grades)
	fmt.Println("Is Alice's mark present in the map?", Exists)

	grades, exists := studentGrades["Charlie"]
	fmt.Println("Mark of Charlie is :", grades)
	fmt.Println("Is Charlie's mark present in the map?", exists)

	// range loop to iterate over map unordered collection
	for index, value := range studentGrades {
		fmt.Printf("key is %s and marks is %d\n", index, value)
	}

	// map creation using map literal
	person := map[string]int{
		"John":  25,
		"Jane":  30,
		"Emily": 22,
	}

	for index, value := range person {
		fmt.Printf("------key is %s and marks is %d\n", index, value)
	}

}
