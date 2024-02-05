---
title: "Build from source"
weight: 2
description: >
  Build a development version of TinyGo from source if you want to help improve TinyGo or want to try the latest features.
---

This page details how to build TinyGo from source. If you would like to install a pre-built binary release, please see our [quick install guide](../../../getting-started/install).

You'll need [Go](https://go.dev) installed on your machine to build TinyGo. The Go team provides install documentation [here](https://go.dev/doc/install).

A major dependency of TinyGo is [LLVM](https://llvm.org/). You can either use a version of LLVM already on your system or build LLVM manually. Building manually takes a long time (around an hour depending on how fast your system is) so it is recommended to use a version of LLVM already on your system if that's possible. The links provided below show how to install LLVM one way or the other.

After finishing building TinyGo from source remember to refer to the [additional requirements](./additional-requirements.md) page to ensure you have a fully working TinyGo installation.


## Uninstalling TinyGo before build
It's highly suggested you uninstall TinyGo if it is present on your computer before proceeding.

***If you installed TinyGo via a package manager the command will depend on your operating system.***

**Linux Debian and Ubuntu** users may run the following command to uninstall TinyGo:
```shell
sudo apt remove tinygo
```

---

**Linux Fedora** users may uninstall TinyGo with:
```shell
sudo dnf remove tinygo
```

---

**MacOS** users may uninstall TinyGo with the following command:
```shell
brew uninstall tinygo
```

---

**Windows** users who installed with Scoop may run the following command to uninstall TinyGo:
```shell
scoop uninstall tinygo
```

---

**Any OS with a manually installed TinyGo**: Remove the cloned repository. This will remove LLVM along with the TinyGo root. 

You should also check if there's a remaining `tinygo` executable in your path and remove it too. 

**Linux and MacOS** may run `echo $(which tinygo)` to print the path to the existing TinyGo executable.

**Windows** users may run `where tinygo` to see the where any remaining TinyGo executable is located at.


## Repository cloning (before build)
Start with getting the source code. On Windows, you might want to install the [build dependencies](#build-dependencies) first.

```shell
git clone https://github.com/tinygo-org/tinygo.git
cd tinygo
```

Once cloned, you can can choose which branch to build. The default branch is the latest release, but you can also build the latest development version:

```shell
# If building released version this command is not necessary.
git checkout dev
```
Once the branch is selected, pull submodules:

```shell
git submodule update --init --recursive
```

Now you are ready to build TinyGo- but you must choose whether to build with a [manual LLVM install](./manual-llvm) or with a [system installed LLVM](./bring-your-own-llvm). After building you should also read [additional requirements](./additional-requirements) to make sure you've fulfilled all the requirements for the features of TinyGo you'll be using.
