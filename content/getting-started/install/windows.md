---
title: "Windows"
linkTitle: "Windows"
type: "docs"
weight: 65
description: >
  Windows install guide
---

This page has information on how to install and use TinyGo on Windows 10. If you wish to build TinyGo from source, for example if you intend to contribute to the project, please take a look [here](../../../docs/guides/build).

TinyGo requires Go v1.20+ to be already installed on your machine.

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
tinygo version 0.36.0 windows/amd64 (using go version go1.24 and LLVM version 19.1.2)
```

Upgrading to the latest TinyGo version can be done via scoop with:

```shell
> scoop update tinygo
```

#### WebAssembly/WASI

If you want to compile programs on Windows that will target WebAssembly or WASI, you need to install the `binaryen` tool via Scoop:

```shell
> scoop install binaryen
```

Upgrading to the latest versions can be as easy as:

```shell
> scoop update binaryen
```

#### AVR (e.g. Arduino Uno)

If you want to flash programs on older Arduino boards such as the Arduino Uno, you need to install the `avrdude` tool via Scoop:

```shell
> scoop install avrdude
```

Upgrading to the latest versions can be as easy as:

```shell
> scoop update avrdude
```


### Manual Install

- You MUST use Go 1.20.x+ with the Windows 10 native install of TinyGo.

    - If you have not installed it yet, you can get it from https://golang.org/dl/

    - Choose the download link for Microsoft Windows, Windows 7 or later, Intel 64-bit processor.

- Download the TinyGo binary for Windows file from https://github.com/tinygo-org/tinygo/releases/download/v0.36.0/tinygo0.36.0.windows-amd64.zip

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
    tinygo version 0.36.0 windows/amd64 (using go version go1.24 and LLVM version 19.1.2)
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

### Development Builds

You can download the latest builds from the TinyGo `dev` branch where active development takes place.

To obtain the binary, first go to the list of recent actions for the Windows build:

https://github.com/tinygo-org/tinygo/actions/workflows/windows.yml?query=branch%3Adev

Click on the link for the build you want to download. The most recent one is located at the top.

Scroll down on that page to the "Artifacts" and click to download the file named "release-double-zipped".

As you might suspect from the name, the file is a compressed zip file that contains the zip file with the actual TinyGo build. Extract that to your desired location, and run it to try the latest features and fixes.
