### `docs/goroutines.md`

# Goroutines in Go

A goroutine is a lightweight thread managed by the Go runtime.

Goroutines allow functions to run concurrently.

## Starting a Goroutine

You can start a goroutine by using the `go` keyword:

```go
go doSomething()
```

The function doSomething runs concurrently with the rest of the program.

## Simple Example

Here is a simple example of starting a goroutine:

```go
package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello from goroutine")
}

func main() {
	go sayHello()

	time.Sleep(time.Second)
}
```

The go keyword starts sayHello as a goroutine.

## Multiple Goroutines

A program can start multiple goroutines:

```go
go taskOne()
go taskTwo()
go taskThree()
```

These functions can execute concurrently.

## Goroutines and Channels

Goroutines are often used together with channels. A channel allows goroutines to communicate with each other:

```go
messages := make(chan string)

go func () {
messages <- "Hello from goroutine"
}()

message := <-messages

fmt.Println(message)
```
The goroutine sends a message through the channel, and the main goroutine receives it.
## Important Considerations
Goroutines are lightweight, but concurrent programs still need synchronization.
When multiple goroutines access shared data, you must prevent race conditions.
Go provides channels, mutexes, and the sync package for synchronization.
Goroutines are one of the most important features of Go for building concurrent applications.