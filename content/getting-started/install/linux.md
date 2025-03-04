---
title: "Linux"
linkTitle: "Linux"
type: "docs"
weight: 35
description: >
  Linux install guide
---

This page has information on how to install and use TinyGo on Ubuntu, as well as other Linux distributions.

If you want to use TinyGo to compile your own or sample code, you can install the release version directly on your machine by following the "Quick Install" instructions below.

If you wish to build TinyGo from source, for example if you intend to contribute to the project, please take a look [here](../../../docs/guides/build).

You can also install the development build of TinyGo if you want to test the latest version. For more information, please see [Development Builds](#development-builds).

## Quick Install

[Ubuntu/Debian](#ubuntu-debian)

[Fedora Linux](#fedora-linux)

[Arch Linux](#arch-linux)

### Ubuntu/Debian

You must have Go already installed on your machine in order to install TinyGo. We recommend Go v1.20 or above.

If you are using Ubuntu or another Debian based Linux on an Intel processor, download the DEB file from Github and install it using the following commands:

```shell
wget https://github.com/tinygo-org/tinygo/releases/download/v0.36.0/tinygo_0.36.0_amd64.deb
sudo dpkg -i tinygo_0.36.0_amd64.deb
```

If you are on a Raspberry Pi or other ARM-based Linux computer, you should use this command instead:

```shell
wget https://github.com/tinygo-org/tinygo/releases/download/v0.36.0/tinygo_0.36.0_armhf.deb
sudo dpkg -i tinygo_0.36.0_armhf.deb
```

You will need to ensure that the path to the `tinygo` executable file is in your `PATH` variable.

```shell
export PATH=$PATH:/usr/local/bin
```

You can test that the installation is working properly by running this code which should display the version number:

```shell
$ tinygo version
tinygo version 0.36.0 linux/amd64 (using go version go1.24 and LLVM version 19.1.2)
```

If you are on a 64 bit ARM OS, and running tinygo fails with "no such file or directory", you may need to install the 32 bit C++ runtime library, e.g.:

```shell
sudo apt install libstdc++6:armhf
```

If you are only interested in compiling TinyGo code for WebAssembly then you are now done with the installation.

Otherwise, please continue with the installation of the additional requirements for your desired microcontroller.

### Additional Requirements for Microcontrollers on Ubuntu/Debian

Some boards require a special flashing tool for that particular chip, like `openocd` or `nrfjprog`. See the documentation page for your board as listed [here](../../../docs/reference/microcontrollers/) to see which flashing tool is required for your target board.

If you are only interested in compiling TinyGo code for ARM microcontrollers then you are now done with the installation.

#### AVR (e.g. Arduino Uno)

To flash TinyGo programs for AVR based processors such as the original Arduino Uno you must install `avrdude`:

```shell
sudo apt-get install avrdude
```

This should allow you to flash TinyGo programs on an Arduino Uno or other supported AVR-based board.

#### You are now done with the TinyGo "Quick Install" for Ubuntu/Debian

### Fedora Linux

There is an [Fedora 30 package](https://packages.fedoraproject.org/pkgs/tinygo/tinygo/) available for the latest TinyGo release. To install it:

```shell
sudo dnf install tinygo
```

If you are only interested in compiling TinyGo code for WebAssembly then you are now done with the installation.

Otherwise, please continue with the installation of the additional requirements for your desired microcontroller

### Additional Requirements for Microcontrollers on Fedora Linux

There are some additional requirements to compile TinyGo programs that can run on microcontrollers.

#### AVR (e.g. Arduino Uno)

To compile and flash TinyGo programs for AVR based processors such as the original Arduino Uno you must install avrdude:

```shell
sudo dnf install avrdude
```

This should allow you to flash TinyGo programs on an Arduino or other supported AVR-based board.

#### You are now done with the TinyGo "Quick Install" for Fedora Linux

### Arch Linux

There is an [Arch package](https://archlinux.org/packages/extra/x86_64/tinygo/) available for the latest TinyGo release.

If you are only interested in compiling TinyGo code for WebAssembly then you are now done with the installation.

Otherwise, please continue with the installation of the additional requirements for your desired microcontroller

### Additional Requirements for Microcontrollers on Arch Linux

There are some additional requirements to compile TinyGo programs that can run on microcontrollers.

#### AVR (e.g. Arduino Uno)

To compile and flash TinyGo programs for AVR based processors such as the original Arduino Uno you must install avrdude:

[avrdude package](https://archlinux.org/packages/extra/x86_64/avrdude/)

This should allow you to flash TinyGo programs on an Arduino or other supported AVR-based board.

#### You are now done with the TinyGo "Quick Install" for Arch Linux

### Development Builds

You can download the latest builds from the TinyGo `dev` branch where active development takes place.

To obtain the binary, first go to the list of recent actions for the Linux build:

https://github.com/tinygo-org/tinygo/actions/workflows/linux.yml?query=branch%3Adev

Click on the link for the build you want to download. The most recent one is located at the top.

Scroll down on that page to the "Artifacts" and click to download the file for your desired architecture, for example "linux-amd64-double-zipped".

As you might suspect from the name, the file is a compressed zip file that contains the zip file with the actual TinyGo build. Extract that to your desired location, and run it to try the latest features and fixes.
