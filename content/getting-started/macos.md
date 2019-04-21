---
title: "macOS"
date: 2019-04-15T11:41:45+02:00
draft: false
weight: 1
---

This page has information on how to install and use TinyGo on macOS.

If you want to use TinyGo to compile your own or sample code, you can install the release version directly on your machine by following the "Quick Install" instructions below.

You can also install the full source code to the TinyGo compiler itself, generally for people who wish to contribute to the project or want to build the compiler from sources directly.

The third option is to use the Docker image. This has the benefit of making no changes to your system but has a large download and installation size. For instructions on using the Docker image, please see the page [here](../using-docker).

## Quick Install

You must have Go v1.11+ already installed on your machine in order to install TinyGo.

You can use Homebrew to install TinyGo using the following commands:

```shell
brew tap tinygo-org/tools
brew install tinygo
```

You can test that the installation is working properly by running this code which should display the version number:

```shell
$ tinygo version
tinygo version 0.5.0 darwin/amd64
```

If you are only interested in compiling TinyGo code for WebAssembly then you are done with the installation.

Otherwise, please continue with the installation of the additional requirements for your desired microcontroller.

### Additional Requirements for Microcontrollers

There are some additional requirements to compile TinyGo programs that can run on microcontrollers.

#### ARM Cortex-M

In order to develop for ARM-based microcontrollers you will need to install LLVM 8:

```shell
brew install llvm
```

#### AVR (Arduino)

If you want to compile code for AVR-based microcontrollers such as Arduino, you will need to install gcc-avr:

```shell
brew tap osx-cross/avr
brew install avr-gcc
brew install avrdude
```

## Source Install

You can instead install TinyGo from source, which includes installing the full LLVM compiler toolchain.

First, obtain the TinyGo source code:

```shell
go get -d github.com/tinygo-org/tinygo
cd $GOPATH/src/github.com/tinygo-org/tinygo
```

Once you have the code, you can install the various prerequisites using [`dep`](https://golang.github.io/dep/):

```shell
dep ensure
```

Now you can run the following make tasks to download and build the LLVM toolchain. Please note that this is likely to take at least 1 hour even on a fast machine.

```shell
make llvm-build
```

Once the LLVM toolchain is installed, you can build the TinyGo binary that is linked to local system libraries like this:

```
make
```

The result of the build will be placed into the `build` directory.

### Additional Requirements for Microcontrollers

Before anything can be built for a bare-metal target, you need to generate some
files first:

```shell
make gen-device
```

This will generate register descriptions, interrupt vectors, and linker scripts
for various devices. Also, you may need to re-run this command after updates,
as some updates cause changes to the generated files.

The same additional requirements to compile TinyGo programs that can run on microcontrollers must be fulfilled when installing TinyGo from source. Please follow [these instructions](#additional-requirements-for-microcontrollers) above.

## Docker Install

For instructions on using the Docker image, please see the page [here](../using-docker).
