
Go P()
======

Print debug messages and variable contents with ease

## Overview

Printing debug messages and dumping variable contents in Go via `fmt.Printf()` is clunky to me.  I hate to type things like this over and over again:

```go
import "fmt"

fmt.Printf("name:%#v age:#%v\n", name, age)
```

Not only it's painful to type, it's also potentially confusing since there might be legitimate `fmt.*` calls around. I'm positive that there must be a **dedicated debug message printer with these ideas behind it**:

1. Debug message printer **must not require initialization**.
2. Debug message printer **must be nothing else, but a debug message printer**. It must be distinctly visible in code.
3. Debug message printer **must be invoked the same way regardless of place of invocation**.

This tiny package gives you function `P()` which makes producing debug output tons easier.

## Examples

```go
P("Joe")
P(25)
// Output:
// "Joe"
// 25

type person struct {
    firstName string
    gender string
}

P(person{"Joe", "m"})
// Output:
// gopee_test.person{firstName:"Joe", gender:"m"}

// Basic labels.
P("name:", "Joe", "age:", 25)
// Output:
// name:"Joe"
// age:25

// Compact labels, smarter output.
P(1.1, " name:", "Joe", " age:", 25)
// Output:
// 1.1 name:"Joe" age:25
```

## Setup

On your development host(s):

```sh
$ go get github.com/dadooda/gopee
```

In your code:

```go
import . "github.com/dadooda/gopee"

func main() {
  P("Hello, world!")
}
```

## Setup (inline)

Inline setup makes `P()` implicitly available throughout your package code and tests without need for additional `import` statements here and there. For inline setup:

1. Download `p.go` (along with `p_test.go` if you're pedantic) into your project.
    ```sh
    $ curl https://raw.githubusercontent.com/dadooda/gopee/master/{p,p_test}.go -OO
    ```
2. Edit `p*.go` files to replace `package mypee` (and possibly `package mypee_test`) with relevant name(s) of your package(s).
3. All done, now you can call `P()` just anywhere in your code and tests.

## Cheers!

Feedback of any kind is highly appreciated.

&mdash; Alex Fortuna, &copy; 2019
