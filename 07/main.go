package main

import "log"

func main() {
	person := make(map[string]interface{}, 0)

	person["name"] = "Alice"
	person["age"] = 21
	person["height"] = 167.64

	age, ok := person["age"].(int)
	if !ok {
		log.Fatal("could not assert value to int")
		return
	}

	person["age"] = age + 1

	log.Printf("%+v", person)
}

// 2024/04/01 11:40:07 map[age:22 height:167.64 name:Alice]
