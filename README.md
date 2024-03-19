# Golang Interfaces Explained

1. Provide a plain-English explanation of what interfaces are;
2. Explain why they are useful and how you might want to use them in your code;
3. Talk about what interface{} (the empty interface) and any are;
4. And run through some of the helpful interface types that you'll find in the standard library.

## What is an interface in Go?

An interface type in Go is kind of like a **definition**.
It defines and describes the exact methods that some other type must have.

One example of an interface type from the standard library is the `fmt.Stringer` interface, which looks like this:

```go
type Stringer interface {
    String() string
}
```

We say that something **satisfies this interface** (or **implements this interface**) if it has a method with the exact signature `String() string`.

For example, the following `Book` type satisifies the interface because it has a String() `string` method:

```go
type Book struct {
  Title string
  Author string
}

func (b Book) String() string {
  return fmt.Sprintf("Book: %s - %s", b.Title, b.Author)
}
```

It's not really important what this `Book` type is or does. The only thing that matters is that is has a method called `String()` which returns a `string` value.

Or, as another example, the following `Count` type also satisfies the `fmt.Stringer` interface — again because it has a method with the exact signature `String() string`.

```go
type Count int

func (c Count) String() string {
    return strconv.Itoa(int(c))
}
```

The important thing to grasp is that we have two different types, `Book` and `Count`, which do different things. But the thing they have in common is that they both satisfy the `fmt.Stringer` interface.

You can think of this the other way around too. If you know that an object satisfies the `fmt.Stringer` interface, you can rely on it having a method with the exact signature `String() string` that you can call.

Now for the important part.

**Wherever you see declaration in Go (such as a variable, function parameter or struct field) which has an interface type, you can use an object of any type _so long as it satisfies the interface_.**

For example, let's say that you have the following function:

```go
func WriteLog(s fmt.Stringer) {
    log.Print(s.String())
}
```

Because this `WriteLog()` function uses the `fmt.Stringer` interface type in its parameter declaration, we can pass in any object that satisfies the `fmt.Stringer` interface. For example, we could pass either of the `Book` and `Count` types that we made earlier to the `WriteLog()` method, and the code would work OK.

Additionally, because the object being passed in satisfies the `fmt.Stringer` interface, we know that it has a `String() string` method that the `WriteLog()` function can safely call.

```go
package main

import (
    "fmt"
    "strconv"
    "log"
)

// Declare a Book type which satisfies the fmt.Stringer interface.
type Book struct {
    Title  string
    Author string
}

func (b Book) String() string {
    return fmt.Sprintf("Book: %s - %s", b.Title, b.Author)
}

// Declare a Count type which satisfies the fmt.Stringer interface.
type Count int

func (c Count) String() string {
    return strconv.Itoa(int(c))
}

// Declare a WriteLog() function which takes any object that satisfies
// the fmt.Stringer interface as a parameter.
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
```

This is pretty cool. In the `main` function we've created different `Book` and `Count` types, but passed both of them to the same `WriteLog()` function. In turn, that calls their relevant `String()` functions and logs the result.

If you run the code, you should get some output which looks like this:

```
2024/03/19 19:00:00 Book: Alice in Wonderland - Lewis Carrol
2024/03/19 19:00:00 3
```

I don't want to labor the point here too much. But the key thing to take away is that by using a interface type in our `WriteLog()` function declaration, we have made the function agnostic (or flexible) about the exact **type** of object it receives. All that matters is **what methods it has**.

---
