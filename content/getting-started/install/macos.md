---
title: "macOS"
linkTitle: "macOS"
type: "docs"
weight: 45
description: >
  macOS install guide
---

This page has information on how to install and use TinyGo on macOS. If you wish to build TinyGo from source, for example if you intend to contribute to the project, please take a look [here](../../../docs/guides/build).

You must have Go v1.20+ already installed on your machine in order to install TinyGo.

You can use Homebrew to install TinyGo using the following commands:

```shell
brew tap tinygo-org/tools
brew install tinygo
```

### Alternative installation

#### Mac M1/M2

Download [this](https://github.com/tinygo-org/tinygo/releases/download/v0.36.0/tinygo0.36.0.darwin-arm64.tar.gz) file. Then, run the following commands:

```shell
tar xvzf tinygo0.36.0.darwin-arm64.tar.gz
export PATH=<extract location>/tinygo/bin:$PATH
```

You can test that the installation is working properly by running this code which should display the version number:

```shell
$ tinygo version
tinygo version 0.36.0 darwin/arm64 (using go version go1.24 and LLVM version 19.1.2)
```

If you are only interested in compiling TinyGo code for WebAssembly then you are done with the installation.

Otherwise, please continue with the installation of the additional requirements for your desired microcontroller.

#### Mac Intel

Download [this](https://github.com/tinygo-org/tinygo/releases/download/v0.36.0/tinygo0.36.0.darwin-amd64.tar.gz) file. Then, run the following commands:

```shell
tar xvzf tinygo0.36.0.darwin-amd64.tar.gz
export PATH=<extract location>/tinygo/bin:$PATH
```

You can test that the installation is working properly by running this code which should display the version number:

```shell
$ tinygo version
tinygo version 0.36.0 darwin/amd64 (using go version go1.24 and LLVM version 19.1.2)
```

If you are only interested in compiling TinyGo code for WebAssembly then you are done with the installation.

Otherwise, please continue with the installation of the additional requirements for your desired microcontroller.

### Additional Requirements for Microcontrollers

If you are only interested in compiling TinyGo code for ARM microcontrollers then you are now done with the installation.

Some boards require a special flashing tool for that particular chip, like `openocd` or `nrfjprog`. See the documentation page for your board as listed [here](../../../docs/reference/microcontrollers/) to see which flashing tool is required for your target board.

The boards with USB Mass Storage Device flashing method, like Raspberry Pi Pico (RP2040), reboot and detach "unsafely" during flashing procedure.  
This behaviour upsets OS and it shows a modal notification one shall unmount USB drive properly first.  
To suppress this notification, run the following command and restarting your Mac after that.  
```
sudo defaults write /Library/Preferences/SystemConfiguration/com.apple.DiskArbitration.diskarbitrationd.plist DADisableEjectNotification -bool YES && sudo pkill diskarbitrationd
```

#### AVR (e.g. Arduino Uno)

To flash TinyGo programs for AVR based processors such as the original Arduino Uno you must install `avrdude`:

```shell
brew install avrdude
```

### Development Builds

You can download the latest builds from the TinyGo `dev` branch where active development takes place.

To obtain the binary, first go to the list of recent actions for the macOS build:

https://github.com/tinygo-org/tinygo/actions/workflows/build-macos.yml?query=branch%3Adev

Click on the link for the build you want to download. The most recent one is located at the top.

Scroll down on that page to the "Artifacts" and click to download the file named "release-double-zipped".

As you might suspect from the name, the file is a compressed zip file that contains the zip file with the actual TinyGo build. Extract that to your desired location, and run it to try the latest features and fixes.
