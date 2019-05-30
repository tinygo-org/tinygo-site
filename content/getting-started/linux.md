---
title: "Linux"
date: 2019-04-15T11:41:45+02:00
draft: false
weight: 1
---

This page has information on how to install and use TinyGo on Ubuntu, as well as other Linux distros.

If you want to use TinyGo to compile your own or sample code, you can install the release version directly on your machine by following the "Quick Install" instructions below.

You can instead install the full source code to the TinyGo compiler itself, generally for people who wish to contribute to the project or want to build the compiler from sources.

The third option is to use the Docker image. This has the benefit of making no changes to your system, however the image has a large download and installation size. For instructions on using the Docker image, please see the page [here](../using-docker).

## Quick Install

You must have Go v1.11+ already installed on your machine in order to install TinyGo. We recommend Go v1.12+

If you are using Ubuntu or another Debian based Linux on an Intel processor, download the DEB file from Github and install it using the following commands:

```shell
wget https://github.com/tinygo-org/tinygo/releases/download/v0.6.0/tinygo_0.6.0_amd64.deb
sudo dpkg -i tinygo_0.6.0_amd64.deb
```

If you are on a Raspberry Pi or other ARM-based Linux computer, you should use this command instead:

```shell
wget https://github.com/tinygo-org/tinygo/releases/download/v0.6.0/tinygo_0.6.0_armhf.deb
sudo dpkg -i tinygo_0.6.0_armhf.deb
```

You will need to ensure that the path to the `tinygo` executable file is in your `PATH` variable.

```shell
export PATH=$PATH:/usr/local/tinygo/bin
```

You can test that the installation is working properly by running this code which should display the version number:

```shell
$ tinygo version
tinygo version 0.6.0 linux/amd64
```

If you are only interested in compiling TinyGo code for WebAssembly then you are done with the installation.

Otherwise, please continue with the installation of the additional requirements for your desired microcontroller.

### Additional Requirements for Microcontrollers

There are some additional requirements to compile TinyGo programs that can run on microcontrollers.

#### ARM Cortex-M

To compile TinyGo programs for ARM based processors you must also install Clang 8 (`clang-8`) for building assembly files and the compiler runtime library [https://compiler-rt.llvm.org/](https://compiler-rt.llvm.org/).

One way to do this is to use the LLVM Debian/Ubuntu packages.

First, you must add the LLVM package repository to your system. Note that if you have already done this once, you do not need to do it again.

```shell
wget -O - https://apt.llvm.org/llvm-snapshot.gpg.key|sudo apt-key add -
sudo add-apt-repository "deb http://apt.llvm.org/$(lsb_release -s -c)/ llvm-toolchain-$(lsb_release -s -c)-8 main"
sudo apt-get update
```

Now you can install Clang 8 by running this command:

```shell
sudo apt-get install clang-8
```

Some boards also require a special flashing tool for that particular chip, like `openocd` or `nrfjprog`. See the documentation page for your board as listed [here](../../microcontrollers/) to see which flashing tool is required for your target board.

#### AVR (Arduino)

To compile and flash TinyGo programs for AVR based processors you must install some extra tools:

```shell
sudo apt-get install gcc-avr
sudo apt-get install avr-libc
sudo apt-get install avrdude
```

This should allow you to compile and flash TinyGo programs on an Arduino or other supported AVR-based board.

## Source Install

First, obtain the TinyGo source code:

```shell
go get -d github.com/tinygo-org/tinygo
cd $GOPATH/src/github.com/tinygo-org/tinygo
```

Once you have the code, you can install the various prerequisites using [`dep`](https://golang.github.io/dep/):

```shell
dep ensure
```

You now have two options: build LLVM manually or use a LLVM distributed with
your package manager. The advantage of a manual build is that it includes all
supported targets (including AVR), which not all prebuilt LLVM packages provide.

### With system-installed LLVM

You can use LLVM included with the package manager of your distribution. How
this is done depends on your system. If you have gotten it to work on a
different distribution, please let us know how so we can add it here.

For Debian and Ubuntu, you can install the binaries provided by LLVM on
[apt.llvm.org](http://apt.llvm.org/). For example, the following commands can be
used to install LLVM 8 on Debian Stretch:

```shell
echo 'deb http://apt.llvm.org/stretch/ llvm-toolchain-stretch-8 main' | sudo tee /etc/apt/sources.list.d/llvm.list
wget -O - https://apt.llvm.org/llvm-snapshot.gpg.key | sudo apt-key add -
sudo apt-get update
sudo apt-get install clang-8 llvm-8-dev lld-8 libclang-8-dev
```

For Arch Linux, install the [clang](https://www.archlinux.org/packages/extra/x86_64/clang/),
[llvm](https://www.archlinux.org/packages/extra/x86_64/llvm/) and
[lld](https://www.archlinux.org/packages/extra/x86_64/lld/) packages.
They must be version 8.x.

Installing TinyGo should now be as easy as:

```shell
go install
```

Note that you should not use `make` when you want to build using a
system-installed LLVM, just use the Go toolchain. `make` is used when you want
to use a self-built LLVM.

### With a self-built LLVM

You can also manually build LLVM. This is a long process which takes at least
one hour on most machines. In most cases you can build TinyGo using a
system-installed LLVM. However, some builds do not support the experimental AVR
target so you'll have to build from source if you want to use TinyGo for the
Arduino Uno. The binaries from apt.llvm.org do provide the AVR target, however.

You will need a few extra tools that are required during the build of LLVM,
depending on your distribution. Debian and Ubuntu users can install all required
tools this way:

```shell
sudo apt-get install build-essential git cmake ninja
```

The following command takes care of downloading and building LLVM. It places the
source code in `llvm-build/` and the build output in `llvm/`. It only needs to
be done once until the next LLVM release. Note that the `export` lines are
optional, but using Clang during the build speeds up the build significantly.

```shell
export CC=clang
export CXX=clang++
make llvm-build
```

Once this is finished, you can build TinyGo against this manually built LLVM:

```shell
make
```

This results in a `tinygo` binary in the `build` directory:

```shell
$ ./build/tinygo version
tinygo version 0.6.0 linux/amd64
```

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

The TinyGo Docker image contains a complete development environment for both WASM and microcontroller development already setup. For instructions on using the Docker image, please see the page [here](../using-docker).
