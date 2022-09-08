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

Cast the types as appropriate. For example, code like this:

```go
func split(s []byte) (*byte, int, int) {
	header := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	return (*byte)(unsafe.Pointer(header.Data)), header.Len, header.Cap
}
```

Can easily be supported in both Go versions, with some type casts:

```go
func split(s []byte) (*byte, int, int) {
	header := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	return (*byte)(unsafe.Pointer(header.Data)), int(header.Len), int(header.Cap)
}
```

The difference is that `header.Len` and `header.Cap` are now cast to an int before returning.
