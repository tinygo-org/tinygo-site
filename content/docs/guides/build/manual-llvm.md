---
title: "Manual LLVM build"
weight: 3
description: >
  How to build LLVM yourself, or "manually".
---


| You need to build LLVM manually in the following cases |
|---|
|  You would like to use it for the ESP8266 or ESP32 chips |
| You are using Windows. |
| Your Linux distribution (if you use Linux) does not ship the right LLVM version. |


## Build dependencies
***Depending on which OS you use the command to install
the dependencies needed to build LLVM will differ.*** Skip over to the command for your Operating System (OS) in this section (OS names are bolded):


**Linux Debian and Ubuntu** users may install all required tools this way:
```shell
sudo apt-get install build-essential git cmake ninja-build
```

---

**Linux Fedora** users may install all required tools with:
```shell
sudo dnf groupinstall "Development Tools"
sudo dnf install cmake ninja-build
```

---

**MacOS** users may install these tools using [Homebrew](https://brew.sh/):

```shell
brew install cmake ninja
```

---

**Windows** users may install build dependencies using [Chocolatey](https://chocolatey.org/). Install Chocolatey first, then run the following in a command prompt or PowerShell with administrative privileges:

```shell
choco install --confirm git golang mingw make cmake ninja python
```
**Windows** users should also use *Git Bash* (installed above) to run all the build commands like `make`. The TinyGo build system expects a Unix-like environment that is not normally provided by Windows but is included already in *Git Bash*.

Choco doesn't seem to add CMake automatically to the `$PATH` variable. You can do this manually if needed, in Git bash:

```shell
export PATH="$PATH:/c/Program Files/CMake/bin"
```

## Building LLVM
***The following instructions are common to all operating systems.***

Background: *[LLVM](https://llvm.org/) is a library and collection of tools for building compilers. It is the base of many newer compilers, like Clang, Swift, and the Rust compiler so they can all build on the same high-quality base. Similarly, TinyGo uses LLVM for all its low-level optimization and machine code generation so that it can produce code that's roughly as small and efficient as all these other modern compilers.*

We build LLVM from inside the TinyGo repository we cloned in the [previous step](../). The following command when run inside the TinyGo repo will take care of first downloading the LLVM source code to later build it in the next step. It places the source code in `llvm-project/` and the build output in `llvm-build/`. It only needs to be done once until the next LLVM release (every half year).

Note that the build step may take some time- feel free to grab a drink meanwhile. Warnings emitted through the compilation in this part are normal as of LLVM 17.

```shell
make llvm-source llvm-build
```

When building on Windows, add CCACHE=OFF.

```shell
make llvm-source llvm-build CCACHE=OFF
```


## Building TinyGo

Once this is finished, you can build TinyGo against this manually built LLVM:

```shell
make
```

This results in a `tinygo` binary in the `build` directory:

```shell
$ ./build/tinygo version
tinygo version 0.36.0-dev-d4189fec linux/amd64 (using go version go1.24 and LLVM version 19.1.2)
```

You have successfully built TinyGo from source. Congratulations! What's left now is to complete the [additional requirements](../additional-requirements)

#### Adding tinygo to your path

**Linux** users may choose to run TinyGo from any directory you may want to move the built binary to a location on your path
or add the `./build` directory to your path. The following shell command moves the TinyGo binary to `/user/bin` so that any user can run TinyGo. You may need root privileges to complete this step.
```shell
mv ./build/tinygo /usr/bin/
```
If you ran the above command, **uninstalling** TinyGo is as easy as running `rm /usr/bin/tinygo`.

If you'd prefer a user-based install you can move TinyGo to your `$HOME/go/bin` directory, where the Go compiler installs binaries by default. The directory should be added to your PATH if it hasn't already been added for this to work.