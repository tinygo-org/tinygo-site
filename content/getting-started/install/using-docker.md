---
title: "Docker"
linkTitle: "Docker"
type: "docs"
weight: 55
---

You can use our Docker image to be able to run the TinyGo compiler on your computer without having to install all the dependencies. Read on to learn how.

## Installing

    docker pull tinygo/tinygo:0.36.0

## Using

The paths used here are automatically resolved by `tinygo` relative to the installation directory.
For your own code, you will probably want to use absolute paths.

A docker container exists for easy access to the TinyGo CLI. For example, to compile `wasm.wasm` for the WebAssembly export example:

    docker run --rm -v $(pwd):/src tinygo/tinygo:0.36.0 tinygo build -o wasm.wasm -target=wasm examples/wasm/export

See the [WebAssembly page](../../../docs/guides/webassembly) for more information on executing the compiled
WebAssembly.

To compile `blinky1.hex` targeting an ARM microcontroller, such as the PCA10040:

    docker run --rm -v $(pwd):/src tinygo/tinygo:0.36.0 tinygo build -o /src/blinky1.hex -size=short -target=pca10040 examples/blinky1

To compile `blinky1.hex` targeting an AVR microcontroller such as the Arduino:

    docker run --rm -v $(pwd):/src tinygo/tinygo:0.36.0 tinygo build -o /src/blinky1.hex -size=short -target=arduino examples/blinky1

For projects that have remote dependencies outside of the standard library and
go code within your own project, you will need to map your entire `$GOPATH`
into the docker image for those dependencies to be found:

    docker run -v $GOPATH:/go -e "GOPATH=/go" tinygo/tinygo:0.36.0 tinygo build -o /go/src/github.com/myuser/myrepo/wasm.wasm -target wasm --no-debug /go/src/github.com/myuser/myrepo/wasm-main.go

**note: At this time, tinygo does not resolve dependencies from the /vendor/ folder within your project.**

For microcontroller development you must flash your hardware devices from your host environment, since you cannot run `tinygo flash` from inside the docker container.

So your workflow could be:

- Compile TinyGo code using the Docker container into a HEX file.
- Flash the HEX file from your host environment to the target microcontroller.

### Development Builds

You can also use the Docker image with the latest builds from the TinyGo `dev` branch where active development takes place:

    docker pull tinygo/tinygo-dev:latest
