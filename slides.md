---
theme: dracula
title: Go 101
highlighter: shiki
drawings:
  enabled: false
lineNumbers: true

transition: slide-left
class: text-center
hideInToc: true
mdc: true
---

# Go 101

<div class="w-md ma">
  <img src="/hello_world.png" />
</div>

<div class="abs-br m-6 flex gap-2">
  <a href="https://github.com/ImFlog/go101" target="_blank" alt="GitHub" title="Open in GitHub"
    class="text-xl slidev-icon-btn opacity-50 !border-none !hover:text-white">
    <carbon-logo-github />
  </a>
</div>

---
hideInToc: true
---

# Table of contents

<Toc minDepth="1" maxDepth="1"></Toc>

---
transition: slide-up
layout: section
---

# What is Go ?

Baby don't hurt me.
<div class="w-xs ma">
    <img src="/fancy_gopher.png"/>
</div>

---
transition: slide-up
---

## History

Go is a statically typed, compiled high-level programming language designed at Google by Robert Griesemer, Rob Pike, and
Ken Thompson[^1].

Often referred as Golang because of its former domain name "golang.org", the proper name is Go.

It's open source and supported (not owned by) by Google. It's reputed to be easy to learn and performant.

Lots of modern cloud tooling is built using Go. Some famous projects are:

- 🐳 Docker
- ☸️ Kubernetes
- 🌍 Terraform
- 🔥 Prometheus
- ...

<br/>

<!-- Footer -->
[^1]: [Wikipedia](https://en.wikipedia.org/wiki/Go_(programming_language))
---
layout: two-cols
transition: slide-up
---

## Features

The language features are:

- C style syntax
- memory safety
- garbage collection
- structural typing
- CSP-style concurrency

::right::

## Tooling

With a strong tooling ecosystem

- embedded compiler
- binaries with multi-arch support
- linting / formatting
- dependency management
- testing framework

---
layout: section
transition: slide-up
---

# Language basics

<div class="w-xs ma">
    <img src="/abc_gopher.png"/>
</div>

---
transition: slide-up
---

## Getting started

<br/>

1. After installing via the [official website](https://go.dev/doc/install) or with [asdf](https://asdf-vm.com/).
2. You can create a new project with `go mod init hello`.
3. And write your first program !

<br/>

```go {monaco-run}
package main

import "fmt"

func main() {
	fmt.Println("Hello world !")
}
```

---
transition: slide-up
zoom: 0.8
---

## Types

| Type                                                     | Zero value | Description                                        |
|----------------------------------------------------------|------------|----------------------------------------------------|
| `bool`                                                   | `false`    | `true` or `false`                                  |
| `string`                                                 | `""`       | UTF-8 string                                       |
| `int`, `int8`, `int16`, `int32`, `int64`                 | `0`        | Signed integers                                    |
| `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr` | `0`        | Unsigned integers                                  |
| `byte`                                                   | `0`        | Alias for `uint8`                                  |
| `rune`                                                   | `0`        | Alias for `int32`, represents a Unicode code point |
| `float32`, `float64`                                     | `0`        | Floating point numbers                             |
| `complex64`, `complex128`                                | `0+0i`     | Complex numbers                                    |

You can create a new type based on an existing type.

```go
type MyInt int // Creates a new type checked at compile time
type MyString = string // Creates an alias, can be replaced by a string in the code
```

---
transition: slide-up
layout: two-cols
layoutClass: gap-16
---

## Variables

- `var` keyword is used to declare a variable.
- `:=` is a shorthand for declaring and initializing a variable. It's called the walrus operator.
- Most of the time you can omit the type as Go is a statically typed language, it can infer the type of a variable.
- Multiple variables can be declared at once.

```go
var a int
var b = 2
c := 3
d, e := 4, "toto"
```

::right::

## Constants

In Go everything is mutable (😢). The only exception is the constants.

You can define immutable values using the `const` keyword.
Constants can be character, string, boolean, or numeric values.

As it is evaluated at compile time you cannot use the walrus operator.

```go
const a = 1
```

---
transition: slide-up
---

## Functions

You can use the keyword `func` to define a function. The return type is declared after the parameters.

```go {monaco-run}
import "fmt"

func add(a int, b int) (int, int, int) {
  return a, b, a + b
}

func main() {
  _, _, sum := add(1, 2)
  fmt.Println(sum)
}
```

<br />

## Visibility

In Go, the visibility of a function or a variable is determined by the first letter of the name. If it's uppercase, it's
exported, otherwise it's scoped to its package only.

---
transition: slide-up
layout: two-cols
layoutClass: gap-16
---

## If

```go
if a > 0 {
  fmt.Println("a is positive")
} else if a == 0 {
  fmt.Println("a is zero")
} else {
  fmt.Println("a is negative")
}
```

## Switch

```go
switch a {
case 0:
  fmt.Println("a is zero")
case 1:
    fmt.Println("a is one")
default:
    fmt.Println("a is something else")
    }
```

::right::

## For

```go
for i := 0; i < 10; i++ {
  fmt.Println(i)
}
```

## Range

```go
for i, v := range []int{1, 2, 3} {
  fmt.Println(i, v)
}
```

## While

```go
for a < 10 {
  a++
}
```

---
transition: slide-up
zoom: 0.9
---

## Structures

You can define a structure using the `struct` keyword.
It's a collection of fields that creates a new type in your application.

You can attach methods to a structure or create methods that returns your type.

```go {monaco-run}
import "fmt"

type Person struct {
  Name string
  Age int
}

func (p Person) SayHello() Person {
  fmt.Println("Hello, my name is", p.Name)
  return p
}

func main() {
  p := Person{"John", 42}
  _ = p.SayHello()
}
```

<br />

---
transition: slide-up
---

## Interfaces

Interface if a very important concept in Go.
To put simply, an interface is a set of method signatures that a type must have to implement the interface.

In go you don't declare that you implement an interface, it is implicit.

```go
type Speaker interface {
  Speak() string
}
```

You can then use the interface as a type. Any structure that has a method `Speak` that returns a string is a `Speaker`
and thus can be used as a parameter of type `Speaker`.

```go
func saySomething(s Speaker) {
  fmt.Println(s.Speak())
}
```

It's called *structural typing*. Like *duck typing* but at compile time.

<!--
If it walks like a duck and quacks like a duck, then it is a duck.
-->

---
transition: slide-left
---

## Pointers

In Go, by default everything is passed by value. If you want to pass a reference you can use the `&` operator to get the
address of a variable.

```go
func main() {
  a := 1
  b := &a
  fmt.Println(*b)
}
```

Function can take receiver with a pointer value, this way you can modify the values of the current structure.

```go
type Person struct {
  Name string
  Age int
}

func (p *Person) Anonymize() {
  p.Name = "anynomous"
}
```

⚠️ References default value is nil, so we can get nil pointer exceptions.

<!--
This can be scary at first, and some tricky bugs can happen, by default, we should use pointers only when needed.
-->
---
transition: slide-up
layout: section
---

# Error management

A boring good design.

<div class="w-xs ma">
    <img src="/bomb_gopher.png"/>
</div>

---
transition: slide-up
zoom: 0.9
---

## Errors

In go errors are values, and they are used to signal that something went wrong.

Everything is explicit, but you have to handle error all the time.

```go {monaco-run}
import(
    "fmt"
    "errors"
)

func negate(a int) (int, error) {
    if a == 0 {
        return 0, errors.New("cannot be zero")
    }
    return -a, nil
}

func main() {
    _, err := negate(0)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("No error")
    }
}
```

---
transition: slide-up
zoom: 0.8
---

## Creating errors

<div class="grid grid-cols-2 gap-4">
  <div>

Error is in fact an interface, that looks like this:

```go
type error interface {
    Error() string
}
```

You can thus create your own error type:

```go
// Like this
var AlreadyNegative = errors.New("already negative")

// Or like this
type Negate0 struct {
    Message string
}

func (e Negate0) Error() string {
    return e.Message
}

func negate(a int) (int, error) {
    if a == 0 {
        return 0, Negate0{"cannot be zero"}
    }
    if a < 0 {
        return 0, AlreadyNegative
    }
    return -a, nil
}
```

</div>
<div>
Errors can be wrapped to add context to them.

```go
import "fmt"

func doSomething() error {
    err := someFunction(-1)
    if err != nil {
        return fmt.Errorf("failed to do something: %w", err)
    }
    return nil
}
```

You can also accumulate errors using the `errors` package.

```go
import "errors"

func doSomething() error {
    var err error
    for _, strErr := range []string{"toto", "tata", "tutu"} {
        errors.join(err, strErr)
    }
    return err
}
```

</div>
</div>

<br />
---
zoom: 0.8
---

## Handling errors

You can assert on the error type using the `errors.Is` function.
This also works on the wrapped errors.

```go
import "errors"

func main() {
    err := doSomething()
    switch {
    case errors.Is(err, MyError{}):
        fmt.Println("MyError")
    case errors.Is(err, ErrSomethingElse):
        fmt.Println("Something else")
    default: {}
    }
}
```

Or you could use the `errors.As` function to extract the error and use it.

```go
import "errors"

func main() {
    err := doSomething()
    var myErr *MyError
    if errors.As(err, &myErr) {
        fmt.Println(myErr.someSpecificFunction())
    }
}
```

---

# Testing

Testing is built-in in Go.
You can write tests in the same package as the code you want to test you just have to create a file with the `_test.go`
suffix.

```go
import "testing"

func TestAdd(t *testing.T) {
  got := add(1, 2)
  if got != 3 {
    t.Errorf("add(1, 2) = %d; want 3", got)
  }
}
```

You can test on private function as soon as the tests are in the same package.

---
layout: section
transition: slide-up
---

# Concurrency

One of the big strengths of Go. Built-in, easy to use.

<div class="w-xs ma">
    <img src="/magic_gopher.png"/>
</div>

---
transition: slide-up
---

## Goroutines

A goroutine is a lightweight thread managed by the Go runtime.

Starting a new goroutine is as simple as using the keyword `go` before a function call.

```go {monaco-run}
import (
  "fmt"
  "time"
)

func main() {
  go func() {
    fmt.Println("Hello from another goroutine")
  }()
  fmt.Println("Hello from the main goroutine")

  time.Sleep(1 * time.Second)
}
```

---
transition: slide-up
---

## Channels

Channels are a typed conduit through which you can send and receive values with the channel operator, <-.
The arrow direction is the direction the data flows.

```go {monaco-run}
import "fmt"

func main() {
  ch := make(chan int)
  go func() {
    ch <- 42
  }()
  fmt.Println(<-ch)
}
```

By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without
explicit locks or condition variables.

---
transition: slide-left
---

## Wait groups

A wait group waits for a collection of goroutines to finish.

```go {monaco-run}
import (
  "fmt"
  "sync"
)

func main() {
  var wg sync.WaitGroup
  wg.Add(1)
  go func() {
    fmt.Println("Hello from another goroutine")
    wg.Done()
  }()
  fmt.Println("Hello from the main goroutine")
  wg.Wait()
  fmt.Println("Bye")
}
```

---
layout: section
transition: slide-up
---

# Tooling

Batteries included

<div class="w-xs ma">
    <img src="/tool_gopher.png"/>
</div>

---
transition: slide-up
zoom: 0.9
---

## Using dependencies

Two files are created when you create a new project with `go mod init`:

* `go.mod`-> list dependencies and their versions.
* `go.sum`-> store the hash values of contents of the module's dependencies.

<br/>

Here is a list of commands to interact with your dependencies:

| Command           | Description                | Example                                             |
|-------------------|----------------------------|-----------------------------------------------------|
| `go get`          | Install dependencies       | `go get github.com/gofiber/fiber/v3`                |
| `go install`      | Install an executable      | `go install github.com/swaggo/swag/cmd/swag@latest` |
| `go mod tidy`     | Remove unused dependencies | `go mod tidy`                                       |
| `go mod download` | Download dependencies      | `go mod download`                                   |

---
zoom: 0.8
---

## Building

The Go ecosystem provides a lot of tools to help write cleaner and production ready code.

| Command       | Description         | Example             |
|---------------|---------------------|---------------------|
| `go fmt`      | Format code         | `go fmt ./...`      |
| `go vet`      | Check for errors    | `go vet ./...`      |
| `go lint`     | Lint code           | `golangci-lint run` |
| `go test`     | Run tests           | `go test ./...`     |
| `go build`    | Build an executable | `go build -o myapp` |
| `go run`      | Run a program       | `go run main.go`    |
| `go generate` | Generate code       | `go generate ./...` |
| `go clean`    | Clean the project   | `go clean`          |

There are also some handy flags to test concurrent code and avoid race conditions.

---
layout: image-right
image: /grow_gopher.png
---
# Going further

Lots of things have not been explained in this presentation like:
* pointers
* generics
* context management
* ...

You can find more information on the official Go supports:
* [Go official website](https://go.dev/)
* [Go by example](https://gobyexample.com/)
* [A tour of Go](https://tour.golang.org/welcome/1)
* [100 Go mistakes](https://100go.co)
