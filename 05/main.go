package main

import "fmt"

func main() {
	person := make(map[string]interface{}, 0)

	person["name"] = "Alice"
	person["age"] = 21
	person["height"] = 167.64

	fmt.Printf("%+v", person)
}

// map[age:21 height:167.64 name:Alice]
