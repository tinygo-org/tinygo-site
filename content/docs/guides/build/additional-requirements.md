---
title: "Additional requirements"
weight: 4
description: >
  Build a development version of TinyGo from source if you want to help improve TinyGo or want to try the latest features.
---

We're not done yet. Some extra things need to be built before you can start using TinyGo.

If you haven't already, you need to download llvm-project (for the compiler-rt sources that are needed for most architectures).
Run the following commands inside your cloned TinyGo repository:

```shell
make llvm-source
```

To be able to use TinyGo on a bare-metal target, you need to generate some files first:

```shell
make gen-device
```

To be able to use TinyGo to build WebAssembly binaries, you will need to compile [wasi-libc](https://github.com/WebAssembly/wasi-libc) and [Binaryen](https://github.com/WebAssembly/binaryen):

```shell
make wasi-libc
make binaryen
```

These command may need to be re-run after some updates in TinyGo.

There are also some extra tools you will need to install, depending on your operating system. These tools are avrdude and openocd. Check the additional requirements for your operating system:

  * [Linux](/getting-started/install/linux/)
  * [MacOS](/getting-started/install/macos/)
  * [Windows](/getting-started/install/windows/)
