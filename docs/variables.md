# Variables

Variables are used to store values in Go.

A variable has a name and a type.

## Variable Declaration

Variables can be declared using the `var` keyword:

```go
var name string
var age int
```

In this example, name is a string and age is an integer.

## Variable Initialization

You can declare and initialize a variable at the same time:

```go
var name string = "Amir"
var age int = 35
```

Go can also infer the type from the assigned value:

```go
var name = "Amir"
var age = 35
```

## Short Variable Declaration

Inside functions, you can use the := operator:

```go
name := "Amir"
age := 35
```

The short variable declaration is commonly used in Go functions.

## Multiple Variables

You can declare multiple variables together:

```go
var name, city string
name = "Amir"
city = "Karaj"
```

You can also initialize multiple variables:

```go
var name, city = "Amir", "Karaj"
```

## Constants

Constants are declared using the const keyword:

```go
const pi = 3.14
const appName = "GoDoc AI"
```

Constants cannot be changed after they are declared.

## Variable Types

Go provides several basic variable types:

```go
var name string
var age int
var price float64
var active bool
```

The variables above store a string, integer, floating-point number, and boolean value.

## Zero Values

When a variable is declared without an initial value, Go assigns its zero value:

```go
var name string
var age int
var price float64
var active bool
```

```
The zero values are:
string  -> ""
int     -> 0
float64 -> 0
bool    -> false
```

