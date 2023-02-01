---
title: "Build from source"
weight: 2
description: |
  Build a development version of TinyGo from source if you want to help improve TinyGo or want to try the latest features.
---

This page details how to build TinyGo from source. If you would like to install a pre-built binary release, please see our [quick install guide](../../../getting-started/install).

Start with getting the source code. On Windows, you might want to install the [build dependencies](#build-dependencies) first.

```shell
git clone --recursive https://github.com/tinygo-org/tinygo.git
cd tinygo
```

If you want to use the latest version of TinyGo instead of the latest release, please switch over to the dev branch:

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

For **Debian** or **Ubuntu** you can install LLVM by adding a new apt repository. For more information about this method, see [apt.llvm.org](https://apt.llvm.org/). *Before copying the command below, please replace `xxxxx` with your distribution's codename*.

| Distro | Version | Codename |
|--------|------- |-----------|
| Ubuntu | 18.04  | `bionic`  |
| Ubuntu | 20.04  | `focal`   |
| Ubuntu | 20.10  | `groovy`  |
| Ubuntu | 21.04  | `hirsute` |
| Ubuntu | 22.04  | `jammy`   |
| Debian | 10     | `buster`  |
| Debian | 11     | `bullseye`|
| Debian | sid    | `unstable`|

```shell
echo 'deb http://apt.llvm.org/xxxxx/ llvm-toolchain-xxxxx-15 main' | sudo tee /etc/apt/sources.list.d/llvm.list
```

After adding the apt repository for your distribution you may install the LLVM toolchain packages:

```shell
wget -O - https://apt.llvm.org/llvm-snapshot.gpg.key | sudo apt-key add -
sudo apt-get update
sudo apt-get install clang-15 llvm-15-dev lld-15 libclang-15-dev
```

For **MacOS**, you can install LLVM through [Homebrew](https://formulae.brew.sh/formula/llvm). The Clang/LLVM version from Apple is not supported by TinyGo.

```shell
brew install llvm@15
```

For **Fedora** users you can install LLVM from the repository. Note that the version of LLVM [varies by Fedora version](https://packages.fedoraproject.org/pkgs/llvm/llvm-libs/), for example Fedora 37 has LLVM 15.

```shell
sudo dnf install llvm-devel lld-libs lld
```

After LLVM has been installed, installing TinyGo should be as easy as running the following command:

```shell
go install
```

If you are getting an `gcc` or `g++ not found` error you most likely do not have a working C++ build environment. You'll need the `build-essential` package on Debian or `sudo dnf install make automake gcc gcc-c++` for Fedora based systems.

If you are getting a build error like this, LLVM is not installed as expected:

```
# tinygo.org/x/go-llvm
../../../go/pkg/mod/tinygo.org/x/go-llvm@v0.0.0-20221028183034-8341240c0b32/analysis.go:16:10: fatal error: 'llvm-c/Analysis.h' file not found
#include "llvm-c/Analysis.h" // If you are getting an error here you need to build or install LLVM, see https://tinygo.org/docs/guides/build/
         ^~~~~~~~~~~~~~~~~~~
1 error generated.
```

This can often be fixed by specifying the LLVM version as a build tag, for example `-tags=llvm14` if you have LLVM 14 instead of LLVM 15.

Note that you should not use `make` when you want to build using a system-installed LLVM, just use the Go toolchain. `make` is used when you want to use a self-built LLVM, as in the next section.

If you have gotten this far, please skip over to [Additional requirements](#additional-requirements) below to further set up TinyGo.

### With a self-built LLVM

You can also manually build LLVM. This is a long process which takes at least one hour on most machines. In many cases you can build TinyGo using a system-installed LLVM, see above. More than 8GB of memory is recommend.

You will need a few extra tools that are required during the build of LLVM, depending on your OS.

#### Build dependencies

Debian and Ubuntu users can install all required tools this way:

```shell
sudo apt-get install build-essential git cmake ninja-build
```

Fedora users can install all required tools with:

```shell
sudo dnf groupinstall "Development Tools"
sudo dnf install cmake ninja-build
```

On MacOS you can install these tools using [Homebrew](https://brew.sh/):

```shell
brew install cmake ninja
```

On Windows you can install them using [Chocolatey](https://chocolatey.org/). Install Chocolatey first, then run the following in a command prompt or PowerShell with administrative privileges:

```shell
choco install --confirm git golang mingw make cmake ninja python
```

Use *Git Bash* (installed above) to run all the build commands like `make`. The TinyGo build system expects a Unix-like environment that is not normally provided by Windows but is included already in *Git Bash*.

Choco doesn't seem to add CMake automatically to the `$PATH` variable. You can do this manually if needed, in Git bash:

```shell
export PATH="$PATH:/c/Program Files/CMake/bin"
```

#### Building LLVM

The following command takes care of downloading and building LLVM. It places the source code in `llvm-project/` and the build output in `llvm-build/`. It only needs to be done once until the next LLVM release (every half year).

```shell
make llvm-source llvm-build
```

When building on Windows, add CCACHE=OFF.

```shell
make llvm-source llvm-build CCACHE=OFF
```

#### Building TinyGo

Once this is finished, you can build TinyGo against this manually built LLVM:

```shell
make
```

This results in a `tinygo` binary in the `build` directory:

```shell
$ ./build/tinygo version
tinygo version 0.26.0 linux/amd64 (using go version go1.19.1 and LLVM version 14.0.0)
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

  * [Linux](/getting-started/install/linux/)
  * [MacOS](/getting-started/install/macos/)
  * [Windows](/getting-started/install/windows/)
