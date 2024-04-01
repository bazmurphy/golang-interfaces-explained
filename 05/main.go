package main

import (
	"fmt"
	"log"
)

func main() {
	// we are using the empty interface type here
	person := make(map[string]interface{}, 0)

	person["name"] = "Alice"
	person["age"] = 21
	person["height"] = 167.64

	// if we try to add 1 age it won't work
	// because its now an empty interface type and no longer an int

	// person["age"] = person["age"] + 1
	// invalid operation: person["age"] + 1 (mismatched types interface{} and int)

	// we have to cast it back to an int to be able to add to it
	age, ok := person["age"].(int)
	if !ok {
		log.Fatal("could not assert value to int")
		return
	}

	person["age"] = age + 1

	fmt.Printf("%+v", person)
}

// but in this case it is better to define a Person struct with relevant typed fields
type Person struct {
	Name   string
	Age    int
	Height float32
}
