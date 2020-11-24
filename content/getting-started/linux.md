---
title: "Linux"
date: 2019-04-15T11:41:45+02:00
draft: false
weight: 1
---

This page has information on how to install and use TinyGo on Ubuntu, as well as other Linux distributions.

If you want to use TinyGo to compile your own or sample code, you can install the release version directly on your machine by following the "Quick Install" instructions below.

You can instead install the full source code to the TinyGo compiler itself, generally for people who wish to contribute to the project or want to build the compiler from sources.

The third option is to use the Docker image. This has the benefit of making no changes to your system, however the image has a large download and installation size. For instructions on using the Docker image, please see the page [here](../using-docker).

## Quick Install

[Ubuntu/Debian](#ubuntu-debian)

[Fedora Linux](#fedora-linux)

[Arch Linux](#arch-linux)

### Ubuntu/Debian

You must have Go already installed on your machine in order to install TinyGo. We recommend Go v1.14 or above.

If you are using Ubuntu or another Debian based Linux on an Intel processor, download the DEB file from Github and install it using the following commands:

```shell
wget https://github.com/tinygo-org/tinygo/releases/download/v0.16.0/tinygo_0.16.0_amd64.deb
sudo dpkg -i tinygo_0.16.0_amd64.deb
```

If you are on a Raspberry Pi or other ARM-based Linux computer, you should use this command instead:

```shell
wget https://github.com/tinygo-org/tinygo/releases/download/v0.16.0/tinygo_0.16.0_arm.deb
sudo dpkg -i tinygo_0.16.0_arm.deb
```

You will need to ensure that the path to the `tinygo` executable file is in your `PATH` variable.

```shell
export PATH=$PATH:/usr/local/tinygo/bin
```

You can test that the installation is working properly by running this code which should display the version number:

```shell
$ tinygo version
tinygo version 0.16.0 linux/amd64 (using go version go1.15 and LLVM version 10.0.1)
```

If you are only interested in compiling TinyGo code for WebAssembly then you are now done with the installation.

Otherwise, please continue with the installation of the additional requirements for your desired microcontroller.

### Additional Requirements for Microcontrollers on Ubuntu/Debian

Some boards require a special flashing tool for that particular chip, like `openocd` or `nrfjprog`. See the documentation page for your board as listed [here](../../microcontrollers/) to see which flashing tool is required for your target board.

If you are only interested in compiling TinyGo code for ARM microcontrollers then you are now done with the installation.

#### AVR (e.g. Arduino Uno)

To compile and flash TinyGo programs for AVR based processors such as the original Arduino Uno you must install some extra tools:

```shell
sudo apt-get install gcc-avr
sudo apt-get install avr-libc
sudo apt-get install avrdude
```

This should allow you to compile and flash TinyGo programs on an Arduino or other supported AVR-based board.

#### You are now done with the TinyGo "Quick Install" for Ubuntu/Debian

### Fedora Linux

There is an [Fedora 30 package](https://apps.fedoraproject.org/packages/tinygo) available for the latest TinyGo release. To install it:

```shell
sudo dnf install tinygo
```

If you are only interested in compiling TinyGo code for WebAssembly then you are now done with the installation.

Otherwise, please continue with the installation of the additional requirements for your desired microcontroller

### Additional Requirements for Microcontrollers on Fedora Linux

There are some additional requirements to compile TinyGo programs that can run on microcontrollers.

#### AVR (e.g. Arduino Uno)

To compile and flash TinyGo programs for AVR based processors such as the original Arduino Uno you must install some extra tools:

```shell
sudo dnf install avr-gcc
sudo dnf install avr-libc
sudo dnf install avrdude
```

This should allow you to compile and flash TinyGo programs on an Arduino or other supported AVR-based board.

#### You are now done with the TinyGo "Quick Install" for Fedora Linux

### Arch Linux

There is an [AUR package](https://aur.archlinux.org/packages/tinygo-bin/) available for the latest TinyGo release.

If you are only interested in compiling TinyGo code for WebAssembly then you are now done with the installation.

Otherwise, please continue with the installation of the additional requirements for your desired microcontroller

### Additional Requirements for Microcontrollers on Arch Linux

There are some additional requirements to compile TinyGo programs that can run on microcontrollers.

#### AVR (e.g. Arduino Uno)

To compile and flash TinyGo programs for AVR based processors such as the original Arduino Uno you must install some extra tools:

[avr-gcc package](https://www.archlinux.org/packages/community/x86_64/avr-gcc/)

[avr-libc package](https://www.archlinux.org/packages/community/any/avr-libc/)

[avrdude package](https://www.archlinux.org/packages/community/x86_64/avrdude/)

This should allow you to compile and flash TinyGo programs on an Arduino or other supported AVR-based board.

#### You are now done with the TinyGo "Quick Install" for Arch Linux

## Source Install

***If you have already followed the "Quick Install" instructions above for your distro, you do not need to perform a source install. You are now done with the needed installation. The "Source Install" is for when you want to contribute to TinyGo.***

Start with getting the source code:

```shell
git clone --recursive https://github.com/tinygo-org/tinygo.git
cd tinygo
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
used to install LLVM 10 on Debian Buster:

```shell
echo 'deb http://apt.llvm.org/buster/ llvm-toolchain-buster-10 main' | sudo tee /etc/apt/sources.list.d/llvm.list
wget -O - https://apt.llvm.org/llvm-snapshot.gpg.key | sudo apt-key add -
sudo apt-get update
sudo apt-get install clang-10 llvm-10-dev lld-10 libclang-10-dev
```

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

On Debian use `ninja-build` package instead of `ninja`.

The following command takes care of downloading and building LLVM. It places the
source code in `llvm-project/` and the build output in `llvm-build/`. It only
needs to be done once until the next LLVM release. Note that the `export` lines
are optional, but using Clang during the build speeds up the build
significantly.

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
tinygo version 0.16.0 linux/amd64 (using go version go1.15 and LLVM version 10.0.1)
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
