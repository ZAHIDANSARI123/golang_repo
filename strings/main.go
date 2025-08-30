package main

import ("fmt"
	   "strings"
)

func main() {
	data := "apple, orange, banana, grape"
	parts := strings.Split(data, ", ")
	fmt.Println(parts)

	// count occurrences of a substring
	str := "Hello, World!"
	count := strings.Count(str, "Hellos")
	fmt.Println("Count of :", count)

	// trim spaces from both sides	
	str = "   Hello GoLang   "
	Trimmed := strings.TrimSpace(str)	
	fmt.Println("Trimmed string:", Trimmed)

	// join strings
	str1 := "Jahid"
	str2 := "Ansari"
	result := strings.Join([]string{str1, "ahmad", str2}, " ")
	fmt.Println("Result is:", result)
}
