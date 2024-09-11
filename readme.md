# Day 1
- How to create a Go project using
```
go mod init "project name"
```
- How to run a Go project with 2 ways
```
go run main.go 
//this will run the code without creating the executable file
```
```
go build main.go && ./build_name 
//this will create the executable file
```
- Standard package called "fmt". Format Package. which gives methods for formatting, printing etc
- Basic Math Operations
- Variable Declartions
    - Declaring a variable with showing the type
    ```
    var num int = 10
    ```
    - Declaring a variable without showing the type
    ```
    num := 10
    ```
- Int to String converision.

Note: cannot use "Sum as string: %s\n" (constant of type string) as io.Writer value in argument to fmt.Fprintf: string does not implement io.Writer (missing method Write)
```
sumStr := fmt.Sprintf("%d", sum)
fmt.Fprintf(os.Stdout, "string: %s\n", sumStr)
```

# Day 2
- Basic Data Types
    - int, int8, int16, int32, int64
    - uinteger
    - float32, float64
    - bool
    - string
    - byte
    - complex64, complex128
- Taking Values from CLI
```
var input string
fmt.Print("Enter your input: ")
fmt.Scanln(&input)
fmt.Println("Your input is:", input)
```
- Conditional Statement - IF, IF else, switch
```
var input string
fmt.Print("Enter your input: ")
fmt.Scanln(&input)

switch input {
case "a":
    fmt.Println("You entered a")
case "b":
    fmt.Println("You entered b")
default:
    fmt.Println("You entered something else")
}
```
- Iterational Statement - only For
```
//while, do-while loop are not available in Go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

d := 3
for d < 10 {
    fmt.Println(d)
    d++
}

b := false
for !b {
    fmt.Println("Infinite loop")
    b = true
}

n := 0
for {
    n++
    fmt.Println(n)
    if n > 10 {
        break
    }
}
```
- Arrays
```
var b [3]string
b[0] = "sd"
b[1] = "ss"
b[2] = "fd"
for i := 0; i < len(b); i++ {
    fmt.Println(b[i])
}
for i, v := range b {
    fmt.Println(i, v)
}
```
```
//Multiple Dimensions arrays
var b [3][3]int
var i, j int
for i = 0; i < 3; i++ {
    for j = 0; j < 3; j++ {
        b[i][j] = i + j
    }
}
fmt.Println(b)
```
```
//Declaration and Assigning
arr := make([]int, 5)
b := []int{1, 2, 3, 4, 5}
```
```
b := []int{1, 2, 3, 4, 5}
fmt.Println(b)
// var ele int = 6
b = append([]int{6},b...)
fmt.Println(b)
// To appeand method to the slice
// First msut be a slice(array) and then the elements to be appended
```
```
b := []int{1, 2, 3, 4, 5}
var ele int = 6
var arr [] int
arr = append(arr,ele)
b = append(arr,b...)
fmt.Println(b)
fmt.Println(arr)
```
- Slices - slicing syntax
```
arr := []int{1, 2, 3, 4, 5}
fmt.Println("Array: ", arr)
up := arr[:len(arr)-1]
up1 := arr[len(arr)-1:]
up2 := arr[0:3]
fmt.Println("Updated Arrays: ", up, up1, up2)
```
```
arr := []int{1, 2, 3, 4, 5}
fmt.Println("Array: ", arr)
// Reverse the slice arr
up := make([]int, len(arr))
for i := len(arr) - 1; i >= 0; i-- {
    up = append(up, arr[i])
}
up = up[len(arr):]
fmt.Println("Updated Arrays: ", up)
```
- Map - nothing but a dictionary in python
```
// Example using map
student := make(map[string]int)
student["John"] = 90
student["Alice"] = 85
student["Bob"] = 92

fmt.Println("Student:", student)
fmt.Println("Length of student map:", len(student))

// Accessing values from the map
fmt.Println("John's score:", student["John"])
fmt.Println("Alice's score:", student["Alice"])

// Updating values in the map
student["Bob"] = 95
fmt.Println("Updated Bob's score:", student["Bob"])

// Deleting a key-value pair from the map
delete(student, "Alice")
fmt.Println("Student after deleting Alice:", student)
```

# Day 3
- Structs
- Interfaces
- Type Interface : [Reference](https://go.dev/tour/basics/14)
- Type conversions : [Reference](https://go.dev/tour/basics/13)
- Quiz : [Link](https://www.codequizzes.com/golang)