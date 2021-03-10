---
title: "Windows"
date: 2019-04-15T11:41:45+02:00
draft: false
weight: 3
---

This page has information on how to install and use TinyGo on Windows 10. If you wish to build TinyGo from source, for example if you intend to contribute to the project, please take a look [here](../../guides/build).

TinyGo can run on Windows 10, however please note that TinyGo does not support building exectuables for Windows. It only supports microcontrollers and WebAssembly on Windows.

TinyGo requires Go v1.16+ to be already installed on your machine.

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
tinygo version 0.17.0 windows/amd64 (using go version go1.16 and LLVM version 11.0.0)
```

Upgrading to the latest TinyGo version can be done via scoop with:

```shell
> scoop update tinygo
```

#### Installing Arduino Dependencies

You can install dependencies required for Arduino Development (`avr-gcc` and `avrdude`) also via Scoop:

```shell
> scoop install avr-gcc
> scoop install avrdude
```

Upgrading to the latest versions can be as easy as:

```shell
> scoop update avr-gcc
> scoop update avrdude
```


### Manual Install

- You MUST use Go 1.16.x+ with the Windows 10 native install of TinyGo.

    - If you have not installed it yet, you can get it from https://golang.org/dl/

    - Choose the download link for Microsoft Windows, Windows 7 or later, Intel 64-bit processor.

- Download the TinyGo binary for Windows file from https://github.com/tinygo-org/tinygo/releases/download/v0.17.0/tinygo0.17.0.windows-amd64.zip

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
    tinygo version 0.17.0 windows/amd64 (using go version go1.16 and LLVM version 11.0.0)
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
