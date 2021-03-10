---
title: "Build from source"
weight: 1
---

This page details how to build TinyGo from source. If you would like to install a pre-build release, please see our [quick install guide](../../getting-started/install).

Start with getting the source code. On Windows, you might want to install the [build dependencies](#build-dependencies) first.

```shell
git clone --recursive https://github.com/tinygo-org/tinygo.git
cd tinygo
```

If you want to use the latest version of LLVM instead of the latest release, please switch over to the dev branch:

```shell
git checkout dev
git submodule update --init
```

## Building TinyGo

A major dependency of TinyGo is [LLVM](https://llvm.org/). You can either use a version of LLVM already on your system or build LLVM manually. Building manually takes a long time (around an hour depending on how fast your system is) so it is recommended to use a version of LLVM already on your system if that's possible.

You need to build LLVM manually in the following cases:

  * You would like to use it for the ESP8266 or ESP32 chips.
  * You are using Windows.
  * Your Linux distribution (if you use Linux) does not ship the right LLVM version.

### With system-installed LLVM

Using a system-installed version of LLVM depends on your system, of course.

For Debian or Ubuntu you can install LLVM by adding a new apt repository (please replace `buster` with the appropriate codename). For more information about this method, see [apt.llvm.org](https://apt.llvm.org/).

```shell
echo 'deb http://apt.llvm.org/buster/ llvm-toolchain-buster-11 main' | sudo tee /etc/apt/sources.list.d/llvm.list
wget -O - https://apt.llvm.org/llvm-snapshot.gpg.key | sudo apt-key add -
sudo apt-get update
sudo apt-get install clang-11 llvm-11-dev lld-11 libclang-11-dev
```

For MacOS, you can install LLVM through [Homebrew](https://formulae.brew.sh/formula/llvm). The Clang/LLVM version from Apple is not supported by TinyGo.

```shell
brew install llvm@11
```

After LLVM has been installed, installing TinyGo should be as easy as running the following command:

```shell
go install
```

Note that you should not use `make` when you want to build using a system-installed LLVM, just use the Go toolchain. `make` is used when you want to use a self-built LLVM, as in the next section.

If you have gotten this far, please skip over to [Additional requirements](#additional-requirements) below to further set up TinyGo.

### With a self-built LLVM

You can also manually build LLVM. This is a long process which takes at least one hour on most machines. In many cases you can build TinyGo using a system-installed LLVM, see above.

You will need a few extra tools that are required during the build of LLVM, depending on your OS.

#### Build dependencies

Debian and Ubuntu users can install all required tools this way:

```shell
sudo apt-get install build-essential git cmake ninja-build
```

On MacOS you can install these tools using [Homebrew](https://brew.sh/):

```shell
brew install cmake ninja
```

On Windows you can install them using [Chocolatey](https://chocolatey.org/). Install Chocolatey first, then run the following in a command prompt or PowerShell with administrative privileges:

```shell
choco install --confirm git golang mingw make cmake ninja python
```

#### Building LLVM

The following command takes care of downloading and building LLVM. It places the source code in `llvm-project/` and the build output in `llvm-build/`. It only needs to be done once until the next LLVM release (every half year).

```shell
make llvm-build
```

#### Building TinyGo

Once this is finished, you can build TinyGo against this manually built LLVM:

```shell
make
```

This results in a `tinygo` binary in the `build` directory:

```shell
$ ./build/tinygo version
tinygo version 0.17.0 linux/amd64 (using go version go1.16 and LLVM version 11.0.0)
```

You have successfully built TinyGo from source. Congratulations!

## Additional requirements

We're not done yet. Some extra things need to be built before you can start using TinyGo.

To be able to use TinyGo on a bare-metal target, you need to generate some files first:

```shell
make gen-device
```

To be able to use TinyGo to build WebAssembly binaries, you will need to compile [wasi-libc](https://github.com/WebAssembly/wasi-libc):

```shell
make wasi-libc
```

These command may need to be re-run after some updates in TinyGo.

There are also some extra tools you will need to install, depending on your operating system. These tools are gcc-avr, avr-libc, avrdude, and openocd. Check the additional requirements for your operating system:

  * [Linux](../../getting-started/install/linux/)
  * [MacOS](../../getting-started/install/macos/)
  * [Windows](../../getting-started/install/windows/)
