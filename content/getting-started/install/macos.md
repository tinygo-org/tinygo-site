---
title: "macOS"
linkTitle: "macOS"
type: "docs"
weight: 45
description: >
  macOS install guide
---

This page has information on how to install and use TinyGo on macOS. If you wish to build TinyGo from source, for example if you intend to contribute to the project, please take a look [here](../../../docs/guides/build).

You must have Go v1.18+ already installed on your machine in order to install TinyGo.

You can use Homebrew to install TinyGo using the following commands:

```shell
brew tap tinygo-org/tools
brew install tinygo
```

### Alternative installation

Download [this](https://github.com/tinygo-org/tinygo/releases/download/v0.26.0/tinygo0.26.0.darwin-amd64.tar.gz) file. Then, run the following commands:

```shell
tar xvzf tinygo-0.26.0.darwin-amd64.tar.gz
export PATH=<extract location>/tinygo/bin:$PATH
```

You can test that the installation is working properly by running this code which should display the version number:

```shell
$ tinygo version
tinygo version 0.26.0 darwin/amd64 (using go version go1.19.1 and LLVM version 14.0.0)
```

If you are only interested in compiling TinyGo code for WebAssembly then you are done with the installation.

Otherwise, please continue with the installation of the additional requirements for your desired microcontroller.

### Additional Requirements for Microcontrollers

If you are only interested in compiling TinyGo code for ARM microcontrollers then you are now done with the installation.

Some boards require a special flashing tool for that particular chip, like `openocd` or `nrfjprog`. See the documentation page for your board as listed [here](../../../docs/reference/microcontrollers/) to see which flashing tool is required for your target board.

#### AVR (e.g. Arduino Uno)

To compile and flash TinyGo programs for AVR based processors such as the original Arduino Uno you must install some extra tools:

```shell
brew tap osx-cross/avr
brew install avr-gcc
brew install avrdude
```

#### Xtensa ESP32

To compile TinyGo programs for the Xtensa ESP32 based processors from Espressif, you must install some extra tools:

```shell
brew tap tasanakorn/homebrew-esp32
brew install xtensa-esp32-elf
```
