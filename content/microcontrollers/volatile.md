---
title: "The volatile keyword"
weight: 3
---

Go does not have the `volatile` keyword like C/C++. This keyword is unnecessary in most desktop use cases but is required for memory mapped I/O on microcontrollers and interrupt handlers. As a workaround, any variable of a type annotated with the `//go:volatile` pragma will be marked volatile. For example:

```go
//go:volatile
type volatileBool bool

var isrFlag volatileBool
```

This is a workaround for a limitation in the Go language and should at some point be replaced with something else.
