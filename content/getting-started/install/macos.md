---
title: "macOS"
date: 2019-04-15T11:41:45+02:00
draft: false
weight: 1
---

This page has information on how to install and use TinyGo on macOS. If you wish to build TinyGo from source, for example if you intend to contribute to the project, please take a look [here](../../guides/build).

You must have Go v1.12+ already installed on your machine in order to install TinyGo. We recommend Go v1.16+.

You can use Homebrew to install TinyGo using the following commands:

```shell
brew tap tinygo-org/tools
brew install tinygo
```

Alternatively, you can download and extract the files manually by downloading [this](https://github.com/tinygo-org/tinygo/releases/download/v0.17.0/tinygo0.17.0.darwin-amd64.tar.gz) file.

You can test that the installation is working properly by running this code which should display the version number:

```shell
$ tinygo version
tinygo version 0.17.0 darwin/amd64 (using go version go1.16 and LLVM version 11.0.0)
```

If you are only interested in compiling TinyGo code for WebAssembly then you are done with the installation.

Otherwise, please continue with the installation of the additional requirements for your desired microcontroller.

### Additional Requirements for Microcontrollers

Some boards require a special flashing tool for that particular chip, like `openocd` or `nrfjprog`. See the documentation page for your board as listed [here](../../microcontrollers/) to see which flashing tool is required for your target board.

If you are only interested in compiling TinyGo code for ARM microcontrollers then you are now done with the installation.

#### AVR (e.g. Arduino Uno)

To compile and flash TinyGo programs for AVR based processors such as the original Arduino Uno you must install some extra tools:

```shell
brew tap osx-cross/avr
brew install avr-gcc
brew install avrdude
```
