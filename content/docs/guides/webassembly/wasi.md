---
title: "Using WASI"
weight: 3
description: |
  How to use TinyGo with the WebAssembly System Interface (WASI).
---

TinyGo is very useful for compiling programs for use on servers and other edge devices (WASI).

TinyGo programs can run in [Extism](https://github.com/extism/extism), [Fastly Compute@Edge](https://developer.fastly.com/learning/compute/go/), [Fermyon Spin](https://developer.fermyon.com/spin/go-components), [tau](https://github.com/taubyte/tau), [wazero](https://wazero.io/languages/tinygo/) and many other WebAssembly runtimes.

Both WASI Preview 1 (wasip1) and WASI Preview 2 (wasip2) are currently supported.

Here is a small TinyGo program for use within a WASI host application:

```go
package main

//go:wasmexport add
func add(x, y uint32) uint32 {
	return x + y
}

// main is required for the `wasip1` target, even if it isn't used.
func main() {}
```

To compile the above TinyGo program for use on any WASI runtime:

```shell
GOOS=wasip1 GOARCH=wasm tinygo build -o main.wasm main.go
```

## Commands and reactors

A WASI module is either a *command* or a *reactor*. A command is an ordinary program: the host calls its `_start` function once, and that runs `main`. A reactor works more like a library: the host calls its `_initialize` function once to set it up, and can then call its exported functions whenever it needs to.

The `-buildmode` flag selects which one TinyGo builds. For WASI Preview 1 (`-target=wasip1`, or `GOOS=wasip1 GOARCH=wasm`):

- `-buildmode=default`
  Build a command. The program exits when `main` returns. Use this for programs that do their work in `main`.

- `-buildmode=c-shared`
  Build a reactor whose exported functions are `//go:wasmexport` functions, as in the example above. `main` is not called. Each call to an exported function runs in a goroutine, so it can use goroutines, channels, mutexes and `time.Sleep`. Use this for plugins and other modules whose functions a host calls. The standard Go toolchain builds WASI reactors the same way.

- `-buildmode=wasi-legacy`
  Build a command whose `_start` returns when `main` returns instead of exiting, so that the host can call `//export` functions afterwards. This is how TinyGo 0.34 and earlier built WASI commands. Use it only for hosts that depend on that behavior; see [The wasi-legacy build mode](#the-wasi-legacy-build-mode).

For WASI Preview 2 (`-target=wasip2`) and for `-target=wasm-unknown`, the default build mode is `c-shared`.

## The wasi-legacy build mode

Up to TinyGo 0.34, a WASI command's `_start` returned when `main` returned, and the module stayed usable. Several hosts came to rely on that: they call `_start` to set the program up and then call its exported functions, as if it were a reactor. TinyGo 0.35 changed `_start` to exit through `proc_exit` when `main` returns, as a WASI command should, and that broke those hosts. TinyGo 0.36 added `-buildmode=wasi-legacy` to bring the old behavior back for them.

New code should use `-buildmode=c-shared` instead. Use `wasi-legacy` when you need a host that expects the older behavior to keep working, for example a plugin system built on [wazero](https://wazero.io/) that runs `_start` before calling a plugin's exports.

Here is a program built for such a host:

```go
package main

var counter int32

func main() {
	// Set up state that the exported functions use.
	counter = 10
}

//export increment
func increment() int32 {
	counter++
	return counter
}
```

```shell
tinygo build -target=wasip1 -buildmode=wasi-legacy -o plugin.wasm main.go
```

And a host using wazero, which calls `_start` when it instantiates the module and then calls `increment`:

```go
mod, err := runtime.InstantiateWithConfig(ctx, wasm, wazero.NewModuleConfig())
if err != nil {
	log.Fatal(err)
}
results, err := mod.ExportedFunction("increment").Call(ctx) // returns 11
```

Built with the default build mode, the same call fails, because the module exited at the end of `_start`. wazero reports `module closed with exit_code(0)`.

A `wasi-legacy` module behaves as follows:

- **`_start` runs the package initializers and `main`, then returns** instead of calling `proc_exit`. Global variables and the heap keep their contents, so exported functions can use whatever `main` set up.
- **Export functions with `//export`, not `//go:wasmexport`.** A `//go:wasmexport` function called after `main` has returned panics with `//go:wasmexport function called after main.main returned`, as it would in the default build mode.
- **Exported functions must not block.** After `_start` returns, an `//export` function runs outside any goroutine. With the default scheduler, if it sleeps, waits on a channel, a locked `sync.Mutex` or a `sync.WaitGroup`, or blocks in any other way, the program panics. A goroutine it starts never runs, because nothing runs the scheduler after `_start` has returned. With `-scheduler=none`, `time.Sleep` sleeps in place, but any wait that needs another goroutine to proceed panics. Use `-buildmode=c-shared` for exported functions that need to block.
- **Goroutines work while `main` runs.** Since TinyGo 0.45, `wasi-legacy` works with the default scheduler. Earlier releases require `-scheduler=none`, so programs that use goroutines cannot use `wasi-legacy` with them. When `main` returns, `_start` returns even if other goroutines are still waiting, as in the default build mode.
- **`os.Exit` still exits.** Calling `os.Exit` ends the program through `proc_exit` with the given code, and the module cannot be used after that.
- **C standard I/O is flushed when `main` returns** (TinyGo 0.45 and later), as `_start` did before TinyGo 0.35.
