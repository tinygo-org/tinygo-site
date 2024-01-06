---
title: "Build from source"
weight: 2
description: >
  Build a development version of TinyGo from source if you want to help improve TinyGo or want to try the latest features.
---

This page details how to build TinyGo from source. If you would like to install a pre-built binary release, please see our [quick install guide](../../../getting-started/install).

A major dependency of TinyGo is [LLVM](https://llvm.org/). You can either use a version of LLVM already on your system or build LLVM manually. Building manually takes a long time (around an hour depending on how fast your system is) so it is recommended to use a version of LLVM already on your system if that's possible. The links provided below show how to install LLVM one way or the other.

### Repository cloning (before build)
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
git submodule update --init
```

Now you are ready to build TinyGo- but you must choose whether to build with a [manual LLVM install](./manual-llvm) or with a [system installed LLVM](./bring-your-own-llvm). After building you should also read [additional requirements](./additional-requirements) to make sure you've fulfilled all the requirements for the features of TinyGo you'll be using.