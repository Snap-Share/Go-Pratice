# Expanded GoLang Teaching Roadmap with Theory and Examples

## 1. Basics

### Go installation and setup
- Install Go from golang.org
- Set up GOPATH and workspace

### Hello World program
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

### Variables, data types, and constants
```go
var name string = "John"
age := 30 // Short variable declaration
const Pi = 3.14159
```

### Control structures
```go
// If statement
if age >= 18 {
    fmt.Println("Adult")
} else {
    fmt.Println("Minor")
}

// For loop
for i := 0; i < 5; i++ {
    fmt.Println(i)
}

// Switch
switch day {
case "Monday":
    fmt.Println("Start of the week")
default:
    fmt.Println("Another day")
}
```

## 2. Functions and packages

### Function declaration and usage
```go
func add(a, b int) int {
    return a + b
}

result := add(5, 3)
```

### Variadic functions
```go
func sum(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}

fmt.Println(sum(1, 2, 3, 4)) // Output: 10
```

### Packages and imports
```go
package main

import (
    "fmt"
    "math"
)

func main() {
    fmt.Println(math.Pi)
}
```

## 3. Data structures

### Arrays and slices
```go
// Array
var numbers [5]int = [5]int{1, 2, 3, 4, 5}

// Slice
fruits := []string{"apple", "banana", "orange"}
fruits = append(fruits, "grape")
```

### Maps
```go
ages := map[string]int{
    "Alice": 30,
    "Bob":   25,
}
ages["Charlie"] = 35
```

### Structs
```go
type Person struct {
    Name string
    Age  int
}

p := Person{Name: "Alice", Age: 30}
fmt.Println(p.Name) // Output: Alice
```

## 4. Object-Oriented Programming in Go

Go doesn't have traditional class-based OOP, but it provides mechanisms to achieve OOP principles.

### Methods

Theory: Methods are functions associated with a particular type. They allow you to add behavior to custom types.

Example:
```go
type Rectangle struct {
    Width, Height float64
}

// Method with a receiver of type Rectangle
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func main() {
    r := Rectangle{Width: 10, Height: 5}
    fmt.Printf("Area of rectangle: %.2f\n", r.Area())
}
```

### Interfaces

Theory: Interfaces define a set of method signatures. Types implicitly implement interfaces by implementing all the methods in the interface.

Example:
```go
type Shape interface {
    Area() float64
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func PrintArea(s Shape) {
    fmt.Printf("Area: %.2f\n", s.Area())
}

func main() {
    c := Circle{Radius: 5}
    PrintArea(c)
}
```

### Embedding

Theory: Go supports composition over inheritance through embedding, allowing you to include one struct within another.

Example:
```go
type Person struct {
    Name string
    Age  int
}

type Employee struct {
    Person  // Embedding Person struct
    JobTitle string
}

func main() {
    e := Employee{
        Person:   Person{Name: "Alice", Age: 30},
        JobTitle: "Developer",
    }
    fmt.Printf("%s is a %s\n", e.Name, e.JobTitle)
}
```

## 5. Concurrency

Go's concurrency model is based on goroutines and channels, making it easier to write concurrent programs.

### Goroutines

Theory: Goroutines are lightweight threads managed by the Go runtime. They allow functions to run concurrently.

Example:
```go
func printNumbers(name string) {
    for i := 1; i <= 3; i++ {
        fmt.Printf("%s: %d\n", name, i)
        time.Sleep(100 * time.Millisecond)
    }
}

func main() {
    go printNumbers("Goroutine 1")
    go printNumbers("Goroutine 2")
    time.Sleep(time.Second)
}
```

### Channels

Theory: Channels are the pipes that connect concurrent goroutines, allowing them to communicate and synchronize.

Example:
```go
func sum(s []int, c chan int) {
    sum := 0
    for _, v := range s {
        sum += v
    }
    c <- sum // Send sum to channel
}

func main() {
    s := []int{7, 2, 8, -9, 4, 0}
    c := make(chan int)
    go sum(s[:len(s)/2], c)
    go sum(s[len(s)/2:], c)
    x, y := <-c, <-c // Receive from channel
    fmt.Println(x, y, x+y)
}
```

### Select Statement

Theory: The select statement lets a goroutine wait on multiple communication operations, choosing which one to proceed with if multiple are ready.

Example:
```go
func fibonacci(c, quit chan int) {
    x, y := 0, 1
    for {
        select {
        case c <- x:
            x, y = y, x+y
        case <-quit:
            fmt.Println("quit")
            return
        }
    }
}

func main() {
    c := make(chan int)
    quit := make(chan int)
    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println(<-c)
        }
        quit <- 0
    }()
    fibonacci(c, quit)
}
```

## 6. Error Handling

Go uses explicit error handling instead of exceptions.

### Error Interface

Theory: Go's built-in error interface is used for error handling. Functions often return an error value along with their result.

Example:
```go
type error interface {
    Error() string
}

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

func main() {
    result, err := divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result)
    }
}
```

### Panic and Recover

Theory: Panic is used for unexpected errors, while recover is used to catch and handle panics.

Example:
```go
func mayPanic() {
    panic("a problem")
}

func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered. Error:\n", r)
        }
    }()
    mayPanic()
    fmt.Println("After mayPanic()") // This line won't be reached
}
```

## 7. Testing

Go has a built-in testing framework that simplifies the process of writing and running tests.

### Unit Testing

Theory: The testing package provides support for automated testing of Go packages.

Example:
```go
// main.go
func Add(a, b int) int {
    return a + b
}

// main_test.go
import "testing"

func TestAdd(t *testing.T) {
    sum := Add(2, 3)
    if sum != 5 {
        t.Errorf("Add(2, 3) = %d; want 5", sum)
    }
}
```

### Benchmarking

Theory: Go's testing framework also supports writing benchmarks to measure the performance of your code.

Example:
```go
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(2, 3)
    }
}
```

## 8. Standard Library

Go's standard library is comprehensive and well-designed, providing many useful packages.

### File I/O

Theory: The `os` and `io/ioutil` packages provide functions for file operations.

Example:
```go
import (
    "fmt"
    "io/ioutil"
    "log"
)

func main() {
    content, err := ioutil.ReadFile("file.txt")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("File contents: %s", content)
}
```

### HTTP Client and Server

Theory: The `net/http` package provides HTTP client and server implementations.

Example:
```go
import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
```

### JSON Handling

Theory: The `encoding/json` package provides encoding and decoding of JSON data.

Example:
```go
import (
    "encoding/json"
    "fmt"
)

type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func main() {
    p := Person{"Alice", 30}
    jsonData, _ := json.Marshal(p)
    fmt.Println(string(jsonData))

    var decoded Person
    json.Unmarshal(jsonData, &decoded)
    fmt.Printf("%+v\n", decoded)
}
```

## 9. Advanced Topics

### Reflection

Theory: Reflection allows inspection of structs, interfaces, and other types at runtime.

Example:
```go
import (
    "fmt"
    "reflect"
)

func main() {
    x := 3.4
    fmt.Println("type:", reflect.TypeOf(x))
    v := reflect.ValueOf(x)
    fmt.Println("value:", v)
    fmt.Println("kind:", v.Kind())
}
```

### Context Package

Theory: The `context` package is used for carrying deadlines, cancellation signals, and request-scoped values across API boundaries and between processes.

Example:
```go
import (
    "context"
    "fmt"
    "time"
)

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
    defer cancel()

    select {
    case <-time.After(2 * time.Second):
        fmt.Println("overslept")
    case <-ctx.Done():
        fmt.Println(ctx.Err()) // prints "context deadline exceeded"
    }
}
```

### Generics (Go 1.18+)

Theory: Generics allow you to write functions and data structures that work with any type that satisfies certain constraints.

Example:
```go
func PrintSlice[T any](s []T) {
    for _, v := range s {
        fmt.Printf("%v ", v)
    }
    fmt.Println()
}

func main() {
    PrintSlice([]int{1, 2, 3})
    PrintSlice([]string{"a", "b", "c"})
}
```

This expanded roadmap covers the theory and provides examples for the more advanced topics in Go programming. Each section builds on the previous ones, providing a comprehensive learning path for mastering Go.