package main

import (
	"fmt"
	"log"
	"strconv"
)

type Stringer interface {
	String() string
}

// Declare a Book type which satisfies the fmt.Stringer interface.
// Book satisfies/implements the "Stringer interface"
type Book struct {
	Title  string
	Author string
}

// because it has a Method with the exact signature of Stringer "String() string"
func (b Book) String() string {
	return fmt.Sprintf("Book: %s - %s", b.Title, b.Author)
}

// Declare a Count type which satisfies the fmt.Stringer interface.
// Count also satisfies/implements the "Stringer interface"
type Count int

// because it has a Method with the exact signature of Stringer "String() string"
func (c Count) String() string {
	return strconv.Itoa(int(c))
}

// Declare a WriteLog() function which takes any object that satisfies
// the fmt.Stringer interface as a parameter.

// because this function is using the fmt.Stringer interface type in its parameter declaration
// we can pass in any object that satisfies the fmt.Stringer interface
// we can pass either the Book or Count types to the WriteLog() method and the code would work OK
// additionally because the object being passed in satisfies the fmt.Stringer interface
// we know that is has a String() string method that the WriteLog() function can safely call
func WriteLog(s fmt.Stringer) {
	log.Print(s.String())
}

func main() {
	// Initialize a Count object and pass it to WriteLog().
	book := Book{"Alice in Wonderland", "Lewis Carrol"}
	WriteLog(book)

	// Initialize a Count object and pass it to WriteLog().
	count := Count(3)
	WriteLog(count)
}

// 2024/03/19 19:00:00 Book: Alice in Wonderland - Lewis Carrol
// 2024/03/19 19:00:00 3
