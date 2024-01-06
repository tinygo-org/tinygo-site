---
title: "Build with system-installed LLVM"
weight: 3
description: >
  How to build TinyGo with a system-installed version of LLVM.
---

⚠️ Halt! This is the system installed LLVM guide! Please check the following table and make sure you don't need the [manual LLVM install guide](./manual-llvm) instead! 


| You need to build LLVM manually in the following cases |
|---|
|  You would like to use it for the ESP8266 or ESP32 chips |
| You are using Windows. |
| Your Linux distribution (if you use Linux) does not ship the right LLVM version. |


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
echo 'deb http://apt.llvm.org/xxxxx/ llvm-toolchain-xxxxx-16 main' | sudo tee /etc/apt/sources.list.d/llvm.list
```

After adding the apt repository for your distribution you may install the LLVM toolchain packages:

```shell
wget -O - https://apt.llvm.org/llvm-snapshot.gpg.key | sudo apt-key add -
sudo apt-get update
sudo apt-get install clang-16 llvm-16-dev lld-16 libclang-16-dev
```

For **MacOS**, you can install LLVM through [Homebrew](https://formulae.brew.sh/formula/llvm). The Clang/LLVM version from Apple is not supported by TinyGo.

```shell
brew install llvm
```

For **Fedora** users you can install LLVM from the repository. Note that the version of LLVM [varies by Fedora version](https://packages.fedoraproject.org/pkgs/llvm/llvm-libs/), for example Fedora 37 has LLVM 15.

```shell
sudo dnf install llvm-devel lld-libs lld
```

After LLVM has been installed, installing TinyGo should be as easy as running the following command:

```shell
go install
```

You should now have a working TinyGo installation!

If you have gotten this far, please refer to [Additional requirements](./additional-requirements) to further set up TinyGo.

## Final notes

### `go install` command inner workings
The `go install` command will build the `tinygo` executable and store it to your currently set `$GOBIN` directory. A couple environment variables must be set for TinyGo to work after running this:

* Your `GOBIN` should be set to where you want go binaries to be stored to. Below is an example (for linux systems)
    ```shell
    go env -w GOBIN=$HOME/local/bin
    ```
* The directory should exist! Below is a way to create the directory on linux.
    ```shell
    mkdir -p $HOME/local/bin
    ```
* Your $PATH environment variable should contain the directory so that you can run `tinygo` from the command line! A reliable way to achieve this can be found at the [Go install page](https://go.dev/doc/install#). For linux it consists of adding the line below to the bottom of your `/etc/profile` file (before the `exit 0`):
    ```
    export PATH=$PATH:$HOME/local/bin
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
