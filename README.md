# Go-lang Walkthrough

- Go is strongly typed language, means it strictly enforces data types and prevents operations b/w incompatible types.
- making it a reliable choice for catching potential errors early in the development process.
- Example: . A `string` variable like "hello world" can not be changed to an `int`, such as the number 3.

# Compile & Run Source File:

- We can compile go lang source file using `go build file_name.go` command.
- We can run go lang source file using `go run file_name.go` command.

# Data-Types:

- Some builtin `datatype` in 'golang'

```go
bool    // holds comparison cases like true/false

string  // store group of characters

int  int8  int16  int32  int64  // signed integers, 
uint uint8 uint16 uint32 uint64 uintptr // unsigned integers

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128    // used to store complex number like: "i + 2j"
```

# Variables:

- We can declare variable by using `var` keyword, name of that variable & then the data-type

```go
var variable_name int

// Example of declaring pi:

var pi float64 = 3.14159
```

## Short-hand Declaration of variable:

- We can also use shorthand declaration `:=` to declare variable in place of `var`

```go
var name string

// Same as the above statement: 
name := ""

// Some other example:
cars := 10  // inferred as integer

temperature := 0.0  // inferred as floating point value, because it has decimal value

var isFunny = true  // inferred as boolean type
```

## Type Inference:

- To declare a variable without specifying an explicit type (either by using the `:=` syntax or `var = expression` syntax), the variable's type is inferred from the value on the right hand side.

```go
var i int
j := i  // j is also an int

// However, when the right hand side is a literal value (an untyped numeric constant like 42 or 3.14), 
// the new variable will be an int, float64, or complex128 depending on its precision:
i := 42 // int
f := 3.14   // float64
g := 0.837 + 0.5i   // complex128
```

## Same line Decelerations:

- We can declare multiple variables on same line.

```go
milage, company := 911, "porsche"

// is same as:
milage := 911
company := "porsche"
```

## Constants:

- Constants are declared like variable but the use the `const` keyword.
- Constants can't use the `:=` shorthand declaration syntax.
- The value of a constants can't be changed once it has been declared.


# How to Format Strings in Go

- Go follows the printf tradition from the C language. 
    - `fmt.Printf` – Prints a formatted string to standard out.
    - `fmt.Sprintf()` – Returns the formatted string.


## `%v` - Interpolate the default representation

- The `%v` variant prints the Go syntax representation of a value. 
- We can usually use this if you're unsure what else to use. That said, it's better to use the type-specific variant if we can.


```go
s := fmt.Sprintf("I am %v years old", 10)
// I am 10 years old

s := fmt.Sprintf("I am %v years old", "way too many")
// I am way too many years old


// %s - Interpolate a string
s := fmt.Sprintf("I am %s years old", "way too many")
// I am way too many years old


// %d - Interpolate an integer in decimal form
s := fmt.Sprintf("I am %d years old", 10)
// I am 10 years old


// %f - Interpolate a decimal

s := fmt.Sprintf("I am %f years old", 10.523)
// I am 10.523000 years old

// The ".2" rounds the number to 2 decimal places
s := fmt.Sprintf("I am %.2f years old", 10.523)
// I am 10.53 years old
```

# Conditionals:

- `if` statements in Go don't use parentheses around the condition.
- `else if` and else are supported as you would expect

```go
if height > 6 {
    fmt.Println("You are super tall!")
} else if height > 4 {
    fmt.Println("You are tall enough!")
} else {
    fmt.Println("You are not tall enough!")
}
```

## The initial statement of an if block

- An `if` conditional can have an "initial" statement. The variable(s) created in the initial statement are only defined within the scope of the `if` body.

```go
// Syntax:
if INITIAL_STATEMENT; CONDITION {
}


// Example:
length := getLength(email)
if length < 1 {
    fmt.Println("Email is invalid")
}

// We can do:
if length := getLength(email); length < 1 {
    fmt.Println("Email is invalid")
}
```

# Functions in Go:

- `func` is used to declare function

```go
// Syntax:
func functionName (dataType parameter1, dataType parameter2,...) returnType {
    // scope
}

// Example:
func sub(x int, y int) int {
    return x-y
}
```


