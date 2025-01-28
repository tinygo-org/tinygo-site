---
title: "Build with system-installed LLVM"
weight: 3
description: >
  How to build TinyGo with a system-installed version of LLVM.
---

⚠️ Halt! This is the system installed LLVM guide! Please check the following table and make sure you don't need the [manual LLVM install guide](../manual-llvm) instead! 


| You need to build LLVM manually in the following cases |
|---|
|  You would like to use it for the ESP8266 or ESP32 chips |
| You are using Windows. |
| Your Linux distribution (if you use Linux) does not ship the right LLVM version. |

## System installed LLVM instructions

### Installing LLVM

See [troubleshooting](#troubleshooting) if running into issues during the instructions.

Using a system-installed version of LLVM depends on your system, of course.

#### Debian or Ubuntu

For **Debian** or **Ubuntu** you can install LLVM by adding a new apt repository. For more information about this method, see [apt.llvm.org](https://apt.llvm.org/). *Before copying the command below, please replace `xxxxx` with your distribution's codename*.

| Distro | Version | Codename  |
|--------|------- |------------|
| Ubuntu | 20.04  | `focal`    |
| Ubuntu | 22.04  | `jammy`    |
| Ubuntu | 23.04  | `lunar`    |
| Ubuntu | 23.10  | `mantic`   |
| Ubuntu | 24.04  | `noble`    |
| Ubuntu | 24.10  | `oracular` |
| Debian | 10     | `buster`   |
| Debian | 11     | `bullseye` |
| Debian | 12     | `bookworm` |
| Debian | sid    | `unstable` |

```shell
echo 'deb http://apt.llvm.org/xxxxx/ llvm-toolchain-xxxxx-18 main' | sudo tee /etc/apt/sources.list.d/llvm.list
```

After adding the apt repository for your distribution you may install the LLVM toolchain packages:

```shell
wget -O - https://apt.llvm.org/llvm-snapshot.gpg.key | sudo apt-key add -
sudo apt-get update
sudo apt-get install clang-18 llvm-18-dev lld-18 libclang-18-dev
```

#### MacOS

For **MacOS**, you can install LLVM through [Homebrew](https://formulae.brew.sh/formula/llvm). The Clang/LLVM version from Apple is not supported by TinyGo.

```shell
brew install llvm lld
```
#### Fedora

For **Fedora** users you can install LLVM from the repository. Note that the version of LLVM [varies by Fedora version](https://packages.fedoraproject.org/pkgs/llvm/llvm-libs/), for example Fedora 37 has LLVM 15.

```shell
sudo dnf install llvm-devel lld-libs lld
```


### Build TinyGo

After LLVM has been installed, installing TinyGo should be as easy as running the following command from within the cloned `tinygo` repository (see [instructions on how to clone](./..)):

```shell
go install
```

You should now have a working TinyGo installation! What's left now is to complete the [additional requirements](../additional-requirements)

Below is an example of running `tinygo version` and example output to check that
TinyGo was installed correctly (copy only what's in front of the `$` sign to your terminal!):
```shell
$ tinygo version
tinygo version 0.35.0-dev-d4189fec linux/amd64 (using go version go1.23 and LLVM version 18.1.2)
```

If not see the [troubleshooting](#troubleshooting) section.

## Troubleshooting

### `go install` command inner workings
The `go install` command will build the `tinygo` executable and store it to your currently set `$GOBIN` directory. A couple environment variables must be set for TinyGo to work after running this:

* Your go environment variable `GOBIN` points to where executables are installed with the
    `go install` command. An empty `GOBIN` variable will default to `$HOME/go/bin` path on Linux and MacOS (same as `~/go/bin` with bash).
    You can check Go's environment variables with `go env`. `go env GOBIN` will print
    only `GOBIN`'s value.
 
* Your $PATH environment variable should contain the `GOBIN` directory so that you can run `tinygo` from the command line! A reliable way to achieve this can be found at the [Go install page](https://go.dev/doc/install#). For linux it consists of adding the line below to the bottom of your `/etc/profile` file (before the `exit 0`):
    ```
    export PATH=$PATH:$HOME/go/bin
    ```

### Debian LLVM repository manual addition
If `sudo apt-get install clang-xx llvm-xx-dev lld-xx libclang-xx-dev` does not work, where `xx` is the LLVM version required by the TinyGo branch you are building i.e: 18, please try manually adding the repository. Run the following command and try the steps above again.
```shell
wget -O - https://apt.llvm.org/llvm-snapshot.gpg.key | sudo apt-key add -
```

### GCC errors
If you are getting an `gcc` or `g++ not found` error you most likely do not have a working C++ build environment. You'll need the `build-essential` package on Debian or `sudo dnf install make automake gcc gcc-c++` for Fedora based systems.

If you are getting a build error like this, LLVM is not installed as expected:

```
# tinygo.org/x/go-llvm
../../../go/pkg/mod/tinygo.org/x/go-llvm@v0.0.0-20221028183034-8341240c0b32/analysis.go:16:10: fatal error: 'llvm-c/Analysis.h' file not found
#include "llvm-c/Analysis.h" // If you are getting an error here you need to build or install LLVM, see https://tinygo.org/docs/guides/build/
         ^~~~~~~~~~~~~~~~~~~
1 error generated.
```

This can often be fixed by specifying the LLVM version as a build tag, for example `-tags=llvm14` if you have LLVM 14 instead of LLVM 16.

### Additional notes
Note that you should not use `make` when you want to build using a system-installed LLVM, just use the Go toolchain.
