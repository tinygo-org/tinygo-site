---
title: "Optimizing binaries"
weight: 11
description: |
  How to improve speed or reduce code size for TinyGo programs.
---

There are various ways TinyGo programs can be optimized for speed, size, or debuggability.

## Optimizing for code size

By default, binaries are already optimized for code size (`-opt=z`, equivalent to `-Oz`, is the default). However, there are various things you can do to improve code size.

### Do not import large packages

You may be tempted to import packages like `fmt` for simple formatted printing. But `fmt` is very large, so unless you absolutely need it, it is recommended to try to avoid it.

In many cases, you can use `println` instead:

```go
println("2 + 2 =", 2 + 2) // prints: "2 + 2 = 4"
```

### Remove debug symbols (WebAssembly)

You can omit debug symbols to reduce the size of WebAssembly binaries using the `-no-debug` flag:

    $ tinygo build -o test.wasm examples/serial && ls -lh test.wasm
    -rwxrwxr-x 1 ayke ayke 93K  4 sep 17:04 test.wasm
    
    $ tinygo build -o test.wasm -no-debug examples/serial && ls -lh test.wasm
    -rwxrwxr-x 1 ayke ayke 30K  4 sep 17:04 test.wasm

As you can see, debug information is ⅔ of this binary so removing it helps a lot!

This may also help on desktop systems (Windows, macOS, Linux) if you want small binaries, but you can also simply use the `strip` utility. It does not have an effect on microcontrollers (Arduino etc) because the debug information is not stored in the firmware image.

### Other options

You can also try the following:

  - Disable goroutines using `-scheduler=none`. It will remove support for the `go` keyword and some related features.
  - Don't print panic messages using `-panic=trap`. This will make debugging harder but it can cut down on code size.
  - Disable the GC using `-gc=leaking` (for very short-lived programs only).
  - Try `-opt=s` and `-opt=2`. While these options normally increase binary size (for higher performance), in rare cases they will actually reduce binary size.

With all these flags applied, the binary is reduced from 93K to just 1.6K!

    $ tinygo build -o test.wasm -no-debug -panic=trap -scheduler=none -gc=leaking examples/serial && ls -lh test.wasm
    -rwxrwxr-x 1 ayke ayke 1,6K  4 sep 17:30 test.wasm

## Optimize for speed

Flags you can try, roughly in order of importance:

  - `-opt=2` to optimize for speed over code size. The default is to optimize for size (`-opt=z`).
  - `-gc=leaking` disables the garbage collector, which can sometimes have a large effect (especially on WebAssembly). Of course, memory will never be freed so this is only appropriate for very short-lived programs.
  - `-scheduler=none` disables the scheduler. This won't usually have much of an effect, but might help on WebAssembly because WebAssembly doesn't have native support for stack switching ([yet](https://github.com/WebAssembly/stack-switching)).
  - `-panic=trap` might help a little bit by removing some code.


## Optimize for debugging

While TinyGo binaries by default include debug information, they are also heavily optimized for size. Therefore, debugging might be difficult by default. You can try the following to make it easier:

  - `-opt=1` will disable some optimization passes which can make debugging somewhat easier.
  - `-opt=0` will remove most optimization passes, which should make debugging even easier than `-opt=1`. However, because so many passes are disabled, some programs don't work well anymore.
  - Of course, don't use `-no-debug`.
