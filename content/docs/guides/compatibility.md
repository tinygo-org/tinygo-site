---
title: "Porting code to TinyGo"
type: "docs"
description: >
  How to port existing code to TinyGo.
---

TinyGo tries to stay compatible with the `gc` compiler toolchain. It should over time become more and more compatible as new features are implemented. However, if you're porting existing code over then you may encounter a number of issues.

## Serialization and deserialization

TinyGo does have limited support for the `reflect` package, but it is not complete. This is especially noticeable with serialization/deserialization packages such as `encoding/json`.

### How to fix

Your best bet is is to use a different package than the standard library package. Packages optimized for speed will often work because they may avoid using reflection.

## `reflect.SliceHeader` and `reflect.StringHeader`

These two structs are sometimes used for unsafely casting between strings and slices. In Go, the structs look like this:

```go
type SliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}
```

In TinyGo, they look like this:

```go
type SliceHeader struct {
	Data uintptr
	Len  uintptr
	Cap  uintptr
}
```

The difference is that `Len` and `Cap` are of type `uintptr` instead of `int`. This is because TinyGo supports AVR, which is a mixed 8/16-bit architecture with 16-bit pointers. The Go language requires `int` to be either 32-bit or 64-bit. To solve this mismatch, TinyGo has chosen to use `uintptr` for the `Len` and `Cap` fields.

Also, note the comment, which explicitly states that using this struct is not portable (emphasis mine):

> SliceHeader is the runtime representation of a slice. **It cannot be used safely or portably** and its representation may change in a later release. Moreover, the Data field is not sufficient to guarantee the data it references will not be garbage collected, so programs must keep a separate, correctly typed pointer to the underlying data. 

### How to fix

Don't use these. Instead, use [`unsafe.Slice`](https://pkg.go.dev/unsafe#Slice), [`unsafe.SliceData`](https://pkg.go.dev/unsafe#SliceData), [`unsafe.String`](https://pkg.go.dev/unsafe#String), and [`unsafe.StringData`](https://pkg.go.dev/unsafe#StringData) (available since Go 1.20).

For example, here is how you can unsafely cast a string to a byte slice:

```go
// Warning: don't ever write to this byte slice!
func unsafeStringToBytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}
```

And here is how you can unsafely cast a `[]uint32` to a `[]byte` slice.

```go
// Warning: byte slice depends on endianness!
func unsafeIntToBytes(slice []uint32) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(slice))), len(slice)*4)
}
```

These functions are fully supported in TinyGo and unlike `reflect.StringHeader` and `reflect.SliceHeader` are portable across compilers and architectures.
