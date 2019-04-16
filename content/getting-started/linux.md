---
title: "Linux"
date: 2019-04-15T11:41:45+02:00
draft: false
weight: 1
---

This page has information on how to install and use TinyGo on Ubuntu, as well as other Linux distros.

If you want to use TinyGo to compile your own or sample code, you can install the release version directly on your machine by following the "Quick Install" instructions below.

You can instead install the full source code to the TinyGo compiler itself, generally for people who wish to contribute to the project or want to build the compiler from sources.

The third option is to use the Docker image. This has the benefit of making no changes to your system, however the image has a large download and installation size.

## Quick Install

If you are using Ubuntu or another Debian based Linux, download the DEB file from Github and install it using the following commands:

```shell
wget https://github.com/tinygo-org/tinygo/releases/download/v0.5.0/tinygo0.5.0.linux-amd64.deb
sudo dpkg -i tinygo0.5.0.linux-amd64.deb
```

You will need to ensure that the path to the `tinygo` executable file is in your `PATH` variable.

```shell
export PATH=$PATH:/usr/local/tinygo/bin
```

You can test that the installation is working properly by running this code which should display the version number:

```shell
$ tinygo version
tinygo version 0.5.0 linux/amd64
```

If you are only interested in compiling TinyGo code for WebAssembly then you are done with the installation.

Otherwise, please continue with the installation of the additional requirements for your desired microcontroller.

### Additional Requirements for Microcontrollers

There are some additional requirements to compile TinyGo programs that can run on microcontrollers.

#### ARM Cortex-M

To compile TinyGo programs for ARM based processors you must install CLang 8 (`clang-8`) for building assembly files and the compiler runtime library [https://compiler-rt.llvm.org/](https://compiler-rt.llvm.org/).

Some boards also require a special flashing tool for that particular chip, like `openocd` or `nrfjprog`. See the documentation page for your board as listed [here](../../microcontrollers/) to see which flashing tool is required for your target board.

#### AVR (Arduino)

To compile and flash TinyGo programs for AVR based processors you must install some extra tools:

```shell
sudo apt-get install avr-gcc
sudo apt-get install avr-libc
sudo apt-get install avrdude
```

This should allow you to compile and flash TinyGo programs on an Arduino or other supported AVR-based board.

## Source Install

You can install TinyGo from source, which includes installing the LLVM compiler toolchain.

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
make llvm-source
make llvm-build
```

Once the LLVM toolchain is installed, you can build the TinyGo binary that is linked to local system libraries like this:

```shell
make tinygo
```

The `tinygo` executable file will be placed into the `build` directory.

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

The TinyGo Docker image contains a complete development environment for both WASM and microcontroller development already setup. For instructions on using the Docker image, please see the page [here](./using-docker).
