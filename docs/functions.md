# Functions in Go

## Functions are reusable blocks of code.

A function can be declared using the `func` keyword.

## Declaring a Function

The following example defines a function named `add`:

```go
func add(a int, b int) int {
    return a + b
}
```
The add function receives two integers and returns their sum.
## Calling a Function
You can call a function by using its name followed by parentheses:
```go
result := add(10, 20)
fmt.Println(result)
```
Functions help you organize code into small and reusable pieces.
## Multiple Return Values
Go functions can return multiple values:

```go
func divide(a int, b int) (int, int) {
    quotient := a / b
    remainder := a % b

    return quotient, remainder
}
```

You can receive both return values:
```go
quotient, remainder := divide(10, 3)

fmt.Println(quotient)
fmt.Println(remainder)
```
## Function Returning an Error
A common pattern in Go is returning an error from a function:
```go
func findUser(id int) (string, error) {
    if id <= 0 {
        return "", errors.New("invalid user id")
    }

    return "Amir", nil
}
```
The caller can check the error:

```go
name, err := findUser(10)

if err != nil {
    fmt.Println(err)
    return
}

fmt.Println(name)
```
Functions make Go code easier to organize, test, reuse, and maintain.
