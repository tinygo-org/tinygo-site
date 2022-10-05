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

## Quick Install

[Ubuntu/Debian](#ubuntu-debian)

[Fedora Linux](#fedora-linux)

[Arch Linux](#arch-linux)

### Ubuntu/Debian

You must have Go already installed on your machine in order to install TinyGo. We recommend Go v1.18 or above.

If you are using Ubuntu or another Debian based Linux on an Intel processor, download the DEB file from Github and install it using the following commands:

```shell
wget https://github.com/tinygo-org/tinygo/releases/download/v0.26.0/tinygo_0.26.0_amd64.deb
sudo dpkg -i tinygo_0.26.0_amd64.deb
```

If you are on a Raspberry Pi or other ARM-based Linux computer, you should use this command instead:

```shell
wget https://github.com/tinygo-org/tinygo/releases/download/v0.26.0/tinygo_0.26.0_armhf.deb
sudo dpkg -i tinygo_0.26.0_armhf.deb
```

You will need to ensure that the path to the `tinygo` executable file is in your `PATH` variable.

```shell
export PATH=$PATH:/usr/local/bin
```

You can test that the installation is working properly by running this code which should display the version number:

```shell
$ tinygo version
tinygo version 0.26.0 linux/amd64 (using go version go1.19.1 and LLVM version 14.0.0)
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

To compile and flash TinyGo programs for AVR based processors such as the original Arduino Uno you must install some extra tools:

```shell
sudo apt-get install gcc-avr
sudo apt-get install avr-libc
sudo apt-get install avrdude
```

This should allow you to compile and flash TinyGo programs on an Arduino or other supported AVR-based board.

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

To compile and flash TinyGo programs for AVR based processors such as the original Arduino Uno you must install some extra tools:

```shell
sudo dnf install avr-gcc
sudo dnf install avr-libc
sudo dnf install avrdude
```

This should allow you to compile and flash TinyGo programs on an Arduino or other supported AVR-based board.

#### You are now done with the TinyGo "Quick Install" for Fedora Linux

### Arch Linux

There is an [Arch package](https://archlinux.org/packages/community/x86_64/tinygo/) available for the latest TinyGo release.

If you are only interested in compiling TinyGo code for WebAssembly then you are now done with the installation.

Otherwise, please continue with the installation of the additional requirements for your desired microcontroller

### Additional Requirements for Microcontrollers on Arch Linux

There are some additional requirements to compile TinyGo programs that can run on microcontrollers.

#### AVR (e.g. Arduino Uno)

To compile and flash TinyGo programs for AVR based processors such as the original Arduino Uno you must install some extra tools:

[avr-gcc package](https://www.archlinux.org/packages/community/x86_64/avr-gcc/)

[avr-libc package](https://www.archlinux.org/packages/community/any/avr-libc/)

[avrdude package](https://www.archlinux.org/packages/community/x86_64/avrdude/)

This should allow you to compile and flash TinyGo programs on an Arduino or other supported AVR-based board.

#### You are now done with the TinyGo "Quick Install" for Arch Linux

