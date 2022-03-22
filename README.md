# koron-go/ptr

[![PkgGoDev](https://pkg.go.dev/badge/github.com/koron-go/ptr)](https://pkg.go.dev/github.com/koron-go/ptr)
[![Actions/Go](https://github.com/koron-go/ptr/workflows/Go/badge.svg)](https://github.com/koron-go/ptr/actions?query=workflow%3AGo)
[![Go Report Card](https://goreportcard.com/badge/github.com/koron-go/ptr)](https://goreportcard.com/report/github.com/koron-go/ptr)

Package `ptr` provides convert functions between value and pointer with using
generics.

## Usage

Import `github.com/koron-go/ptr`.

Get a pointer of a value.

```go
ptr.P(123)   // (*int)
ptr.P("foo") // (*string)
```

Get a value which referenced by a pointer.

```go
n := 123
ptr.V(&n) // 123 (int)

s := "abc"
ptr.V(&s) // "abc" (string)
```

