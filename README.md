# Golang Interfaces Explained

1. Provide a plain-English explanation of what `interfaces` are;
2. Explain why they are useful and how you might want to use them in your code;
3. Talk about what `interface{}` (the empty `interface`) and any are;
4. And run through some of the helpful `interface` `type`s that you'll find in the standard library.

## What is an interface in Go?

An `interface` `type` in Go is kind of like a **definition**.
It defines and describes the exact methods that some other `type` must have.

One example of an `interface` `type` from the standard library is the `fmt.Stringer` `interface`, which looks like this:

```go
`type` Stringer interface {
    String() string
}
```

We say that something **satisfies this interface** (or **implements this `interface`**) if it has a method with the exact signature `String() string`.

For example, the following `Book` `type` satisifies the `interface` because it has a `String() string` method:

```go
`type` Book struct {
  Title string
  Author string
}

func (b Book) String() string {
  return fmt.Sprintf("Book: %s - %s", b.Title, b.Author)
}
```

It's not really important what this `Book` `type` is or does. The only thing that matters is that is has a method called `String()` which returns a `string` value.

Or, as another example, the following `Count` `type` also satisfies the `fmt.Stringer` `interface` — again because it has a method with the exact signature `String() string`.

```go
`type` Count int

func (c Count) String() string {
    return strconv.Itoa(int(c))
}
```

The important thing to grasp is that we have two different `type`s, `Book` and `Count`, which do different things. But the thing they have in common is that they both satisfy the `fmt.Stringer` `interface`.

You can think of this the other way around too. If you know that an object satisfies the `fmt.Stringer` interface, you can rely on it having a method with the exact signature `String() string` that you can call.

Now for the important part.

**Wherever you see declaration in Go (such as a variable, function parameter or `struct` field) which has an `interface` `type`, you can use an object of any `type` _so long as it satisfies the `interface`_.**

For example, let's say that you have the following function:

```go
func WriteLog(s fmt.Stringer) {
    log.Print(s.String())
}
```

Because this `WriteLog()` function uses the `fmt.Stringer` `interface` `type` in its parameter declaration, we can pass in any object that satisfies the `fmt.Stringer` `interface`. For example, we could pass either of the `Book` and `Count` `type`s that we made earlier to the `WriteLog()` method, and the code would work OK.

Additionally, because the object being passed in satisfies the `fmt.Stringer` `interface`, we know that it has a `String() string` method that the `WriteLog()` function can safely call.

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

This is pretty cool. In the `main` function we've created different `Book` and `Count` `type`s, but passed both of them to the same `WriteLog()` function. In turn, that calls their relevant `String()` functions and logs the result.

If you run the code, you should get some output which looks like this:

```
2024/03/19 19:00:00 Book: Alice in Wonderland - Lewis Carrol
2024/03/19 19:00:00 3
```

I don't want to labor the point here too much. But the key thing to take away is that by using a `interface` `type` in our `WriteLog()` function declaration, we have made the function agnostic (or flexible) about the exact **`type`** of object it receives. All that matters is **what methods it has**.

---

## Why are they useful?

There are all sorts of reasons that you might end up using a `interface` in Go, but in my experience the three most common are:

1. To help reduce duplication or boilerplate code.
2. To make it easier to use mocks instead of real objects in unit tests.
3. As an architectural tool, to help enforce decoupling between parts of your codebase.

Let's step through these three use-cases and explore them in a bit more detail.

### Reducing boilerplate code

OK, imagine that we have a `Customer` `struct` containing some data about a customer. In one part of our codebase we want to write the customer information to a `bytes.Buffer`, and in another part of our codebase we want to write the customer information to an `os.File` on disk. But in both cases, we want to serialize the customer `struct` to JSON first.

This is a scenario where we can use Go's interfaces to help reduce boilerplate code.

The first thing you need to know is that Go has an `io.Writer` `interface` `type` which looks like this:

```go
`type` Writer interface {
  Write(p []byte) (n int, err error)
}
```

And we can leverage the fact that both `bytes.Buffer` and the `os.File` `type` satisfy this interface, due to them having the `bytes.Buffer.Write()` and `os.File.Write()` methods respectively.

Let's take a look at a simple implementation:

```go
package main

import (
    "bytes"
    "encoding/json"
    "io"
    "log"
    "os"
)

// Create a Customer `type`
type Customer struct {
    Name string
    Age  int
}

// Implement a WriteJSON method that takes an io.Writer as the parameter.
// It marshals the customer struct to JSON, and if the marshal worked
// successfully, then calls the relevant io.Writer's Write() method.
func (c *Customer) WriteJSON(w io.Writer) error {
    js, err := json.Marshal(c)
    if err != nil {
        return err
    }

    _, err = w.Write(js)
    return err
}

func main() {
    // Initialize a customer struct.
    c := &Customer{Name: "Alice", Age: 21}

    // We can then call the WriteJSON method using a buffer...
    var buf bytes.Buffer
    err := c.WriteJSON(&buf)
    if err != nil {
        log.Fatal(err)
    }

    // Or using a file.
    f, err := os.Create("/tmp/customer")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()


    err = c.WriteJSON(f)
    if err != nil {
        log.Fatal(err)
    }
}
```

Of course, this is just a toy example (and there are other ways we could structure the code to achieve the same end result). But it nicely illustrates the benefit of using an `interface` — we can create the `Customer.WriteJSON()` method once, and we can call that method any time that we want to write to something that satisfies the `io.Writer` `interface`.

But if you're new to Go, this still begs a couple of questions: How do you know that the `io.Writer` `interface` even exists? And how do you know in advance that `bytes.Buffer` and `os.File` both satisfy it?

There's no easy shortcut here I'm afraid — you simply need to build up experience and familiarity with the `interfaces` and different `type`s in the standard library. Spending time thoroughly reading the standard library documentation, and looking at other people's code will help here. But as a quick-start I've included a list of some of the most useful interface `type`s at the end of this post.

But even if you don't use the interfaces from the standard library, there's nothing to stop you from creating and using your own interface `type`s. We'll cover how to do that next.
