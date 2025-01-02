# Go Concurrency

Concurrency is a fundamental feature of Go that allows you to efficiently perform multiple tasks simultaneously. This guide provides a comprehensive overview of Go concurrency, covering its key concepts and usage patterns.

---

## Table of Contents
1. [Introduction to Concurrency](#introduction-to-concurrency)
2. [Goroutines](#goroutines)
3. [Channels](#channels)
4. [Buffered vs. Unbuffered Channels](#buffered-vs-unbuffered-channels)
5. [Select Statement](#select-statement)
6. [WaitGroups](#waitgroups)
7. [Mutex and Synchronization](#mutex-and-synchronization)
8. [Context Package](#context-package)
9. [Concurrency Best Practices](#concurrency-best-practices)
10. [Resources](#resources)

---

## Introduction to Concurrency

Concurrency enables a program to perform multiple tasks at the same time. In Go, concurrency is achieved using lightweight threads called goroutines and channels for communication.

> **Concurrency vs Parallelism:** Concurrency is about dealing with multiple tasks simultaneously. Parallelism is about executing tasks at the same time. Go focuses on concurrency, making it easy to scale tasks across processors.

---

## Goroutines

A goroutine is a lightweight thread managed by the Go runtime.

### Syntax:
```go
package main

import (
    "fmt"
    "time"
)

func say(message string) {
    for i := 0; i < 3; i++ {
        fmt.Println(message)
        time.Sleep(100 * time.Millisecond)
    }
}

func main() {
    go say("Hello")
    say("World")
}
```
### Output:
The `say("Hello")` runs concurrently with `say("World")`. The exact order of execution may vary.

---

## Channels

Channels are used to communicate between goroutines. They can send and receive values.

### Creating a Channel:
```go
ch := make(chan int) // Unbuffered channel
```

### Sending and Receiving:
```go
package main

import "fmt"

func main() {
    ch := make(chan string)

    go func() {
        ch <- "Hello from goroutine!"
    }()

    message := <-ch // Receive
    fmt.Println(message)
}
```
### Output:
```
Hello from goroutine!
```

---

## Buffered vs. Unbuffered Channels

- **Unbuffered Channel:** Requires both sender and receiver to be ready.
- **Buffered Channel:** Allows sending even if the receiver isn't ready (up to the buffer size).

### Example:
```go
ch := make(chan int, 2) // Buffered channel with size 2
ch <- 1
ch <- 2
fmt.Println(<-ch) // Output: 1
fmt.Println(<-ch) // Output: 2
```

---

## Select Statement

The `select` statement is used to wait on multiple channel operations.

### Example:
```go
package main

import "fmt"

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    go func() { ch1 <- "Channel 1" }()
    go func() { ch2 <- "Channel 2" }()

    select {
    case msg1 := <-ch1:
        fmt.Println(msg1)
    case msg2 := <-ch2:
        fmt.Println(msg2)
    }
}
```
### Output:
Either "Channel 1" or "Channel 2" depending on which channel sends first.

---

## WaitGroups

WaitGroups help wait for a collection of goroutines to finish execution.

### Example:
```go
package main

import (
    "fmt"
    "sync"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done() // Decrement counter when done
    fmt.Printf("Worker %d started\n", id)
    fmt.Printf("Worker %d finished\n", id)
}

func main() {
    var wg sync.WaitGroup

    for i := 1; i <= 3; i++ {
        wg.Add(1) // Increment counter
        go worker(i, &wg)
    }

    wg.Wait() // Block until all goroutines are done
}
```

---

## Mutex and Synchronization

A `Mutex` ensures only one goroutine accesses a critical section at a time.

### Example:
```go
package main

import (
    "fmt"
    "sync"
)

var counter int
var mutex sync.Mutex

func increment(wg *sync.WaitGroup) {
    defer wg.Done()

    mutex.Lock()
    counter++
    mutex.Unlock()
}

func main() {
    var wg sync.WaitGroup

    for i := 0; i < 10; i++ {
        wg.Add(1)
        go increment(&wg)
    }

    wg.Wait()
    fmt.Println("Final Counter:", counter)
}
```

---

## Context Package

The `context` package is used to control goroutines' lifetimes, passing deadlines, cancellations, and other signals.

### Example:
```go
package main

import (
    "context"
    "fmt"
    "time"
)

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()

    go func(ctx context.Context) {
        for {
            select {
            case <-ctx.Done():
                fmt.Println("Goroutine stopped:", ctx.Err())
                return
            default:
                fmt.Println("Working...")
                time.Sleep(500 * time.Millisecond)
            }
        }
    }(ctx)

    time.Sleep(3 * time.Second)
}
```
### Output:
```
Working...
Working...
Goroutine stopped: context deadline exceeded
```

---

## Concurrency Best Practices

1. Use goroutines for independent tasks.
2. Prefer channels over shared memory for communication.
3. Avoid deadlocks by designing clear ownership of resources.
4. Limit goroutine counts to avoid overwhelming the scheduler.
5. Use the `context` package for managing goroutine lifetimes.

---

## Resources

- [Go Official Documentation](https://golang.org/doc/)
- [A Tour of Go](https://tour.golang.org/concurrency/1)
- [Effective Go - Concurrency](https://golang.org/doc/effective_go#concurrency)

---

Happy Coding with Go Concurrency! 🚀

