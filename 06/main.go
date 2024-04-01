package main

import "fmt"

func main() {
	person := make(map[string]interface{}, 0)

	person["name"] = "Alice"
	person["age"] = 21
	person["height"] = 167.64

	person["age"] = person["age"] + 1
	// invalid operation: person["age"] + 1 (mismatched types interface{} and int)compilerMismatchedTypes

	fmt.Printf("%+v", person)
}
