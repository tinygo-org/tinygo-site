---
title: "Windows"
date: 2019-04-15T11:41:45+02:00
draft: false
weight: 3
---

This page has information on how to install and use TinyGo on Windows 10.

## Windows Native Install

We now have a native install for Windows 10.

> **VERY IMPORTANT NOTE:** 
> You cannot yet create Windows binary programs using TinyGo, only MCU and WASM targets.

TinyGo requires Go v1.14+ to be already installed on your machine.

### Quick Install via Scoop

You can use [Scoop](https://scoop.sh/) to install TinyGo and dependencies.

If you haven't installed Go already, you can do so with the following command:

```shell
> scoop install go
```

Followed by TinyGo itself:

```shell
> scoop install tinygo
```

Your `$PATH` environment variable will be updated via the scoop package. By default a shim is created in `~/scoop/shims/tinygo`.

You can test that the installation was successful by running the `version` command which should display the version number:

```shell
> tinygo version
tinygo version 0.16.0 windows/amd64 (using go version go1.15 and LLVM version 10.0.1)
```

Upgrading to the latest TinyGo version can be done via scoop with:

```shell
> scoop update tinygo
```

### Manual Install

- You MUST use Go 1.14.x+ with the Windows 10 native install of TinyGo.

    - If you have not installed it yet, you can get it from https://golang.org/dl/

    - Choose the download link for Microsoft Windows, Windows 7 or later, Intel 64-bit processor.

- Download the TinyGo binary for Windows file from https://github.com/tinygo-org/tinygo/releases/download/v0.16.0/tinygo0.16.0.windows-amd64.zip

- Decompress the file like this:

    - First double click on the downloaded ZIP file to open it.

    - Now drag the "tinygo" folder in the ZIP file window onto your "C:" drive.

    - When the folder is finished extracting, you can close the ZIP file window.

- You will need to add `C:\tinygo\bin` to your PATH.

    ```shell
    > set PATH=%PATH%;"C:\tinygo\bin";
    ```

- Now you should be able to run the TinyGo command:

    ```
    > tinygo version
    tinygo version 0.16.0 windows/amd64 (using go version go1.15 and LLVM version 10.0.1)
    ```

### Flashing boards

#### Adafruit Circuit Playground Express

Many boards from Adafruit, such as the Circuit Playground Express, use the `UF2` format. This binary format does not require installing any additional flashing tools.

- Plug in the Circuit Playground Express board to the USB port.

- Double tap the "RESET" button.

- Once the Circuit Playground Express board appears as a flash drive, run the following command:

    ```shell
    tinygo flash -target circuitplay-express examples/blinky1
    ```

The board should restart and begin running your program.

#### Arduino Nano33 IoT

- You must install the "BOSSA" flashing utility first. You can download it from https://github.com/shumatech/BOSSA/releases/download/1.9.1/bossa-x64-1.9.1.msi

- During the installation, make sure that you choose to put it into `c:\Program Files`. The default in the installer will not work correct, you must choose `c:\Program Files`.

- After the installation, you must add BOSSA to your PATH:

    ```shell
    set PATH=%PATH%;"c:\Program Files\BOSSA";
    ```

- Test that you have installed "BOSSA" correctly by running this command:

    ```shell
    bossac --help
    ```

Now you can flash your Arduino Nano33 board like this:

- Plug in the Arduino Nano33 IoT board to the USB port.

- Double tap the "RESET" button.

- Windows 10 should display a notification with the name of the COM port that the board has been assigned to. For example `COM4`.

- Run the following command:

    ```shell
    tinygo flash -target arduino-nano33 path\to\code
    ```

The board should restart and begin running your program.

## WSL Install

If you want to use TinyGo to compile your own or sample code, you should be able to install the release version directly on your machine using Windows Subsystem for Linux (WSL) by following the "Quick Install" instructions on the Linux page located [here](../linux).

## Docker Install

Another option is to use the Docker image. This has the benefit of making no changes to your system but has a large download and installation size. For instructions on using the Docker image, please see the page [here](../using-docker).

## Source Install

***If you have already followed the "Windows Native Install" instructions above, you do not need to perform a source install. You are now done with the needed installation. The "Source Install" is for when you want to contribute to TinyGo.***

Be warned that building TinyGo on Windows is not tested as well as building TinyGo on other operating systems (such as Linux). If you want to contribute to TinyGo but don't need to run natively on Windows, it may be easier and faster to do development within [WSL](https://docs.microsoft.com/en-us/windows/wsl/install-win10). See the [Linux page](../linux) for how to build TinyGo on Linux.

### Dependencies

You will need to have the following programs installed on your Windows system and configured to be accessible in your PATH variable:

  * Git
  * Go 1.14
  * MinGW-w64
  * GNU Make
  * CMake
  * Ninja
  * Python

The easiest way to install all these dependencies is through [Chocolatey](https://chocolatey.org/). Install Chocolatey first, and then run the following command in a command prompt or PowerShell with administrative privileges:

    choco install --confirm git golang mingw make cmake ninja python

Now open a Git Bash window for the remaining steps. The Git Bash window provides a Bash shell with some standard Unix utilities for convenience.

The first thing to do is download the source code:

```shell
git clone --recursive https://github.com/tinygo-org/tinygo.git
cd tinygo
```

Unfortunately there is no way to use a binary release of LLVM to build against (like on Linux and MacOS) so we'll have to build LLVM from scratch. This is a long process which takes at least one hour on most machines.

The following command takes care of downloading and building LLVM. It places the
source code in `llvm-project/` and the build output in `llvm-build/`. It only needs to
be done once until the next LLVM release.

```shell
make llvm-build
```

Once this is finished, you can build TinyGo against this manually built LLVM:

```shell
make
```

This results in a `tinygo.exe` binary in the `build` directory:

```text
$ ./build/tinygo version
tinygo version 0.16.0 windows/amd64 (using go version go1.15 and LLVM version 10.0.1)
```

### Additional Requirements for Microcontrollers

Before anything can be built for a bare-metal target, you need to generate some
files first:

```shell
make gen-device
```

This will generate register descriptions, interrupt vectors, and linker scripts
for various devices. Also, you may need to re-run this command after updates,
as some updates cause changes to the generated files.

The same additional requirements to compile TinyGo programs that can run on microcontrollers must be fulfilled when installing TinyGo from source. Please follow [these instructions](#additional-requirements-for-microcontrollers) above.
