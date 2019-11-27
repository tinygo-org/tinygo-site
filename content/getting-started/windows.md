---
title: "Windows"
date: 2019-04-15T11:41:45+02:00
draft: false
weight: 3
---

This page has information on how to install and use TinyGo on Windows 10.

## Windows Native Install

We now have a native install for Windows 10.

VERY IMPORTANT NOTE: You cannot yet create Windows binary programs using TinyGo, only MCU and WASM targets.

- You MUST use Go 1.13.x+ with the Windows 10 native install of TinyGo.

    - If you have not installed it yet, you can get it from https://golang.org/dl/

    - Choose the download link for Microsoft Windows, Windows 7 or later, Intel 64-bit processor.

- Install LLVM 9 from http://releases.llvm.org/download.html

    - Choose the pre-built binary for Windows (64-bit).

    - If you already have a previous version of LLVM installed, you should choose the option to uninstall it that you will be presented with during the LLVM 9 install.

    - Make sure to choose "C:\Program Files\LLVM\" as the destination folder during the install.

    - During the installation make sure you choose the option to add LLVM to your PATH.

- Now, download the TinyGo binary for Windows file from https://github.com/tinygo-org/tinygo/releases/download/v0.10.0/tinygo0.10.0.windows-amd64.zip

- Decompress the file like this:

    - First double click on the downloaded ZIP file to open it.

    - Now drag the "tinygo" folder in the ZIP file window onto your "C:" drive.

    - When the folder is finished extracting, you can close the ZIP file window.

- You will need to add `C:\tinygo\bin` to your PATH.

    ```shell
    set PATH=%PATH%;"C:\tinygo\bin";
    ```

- Now you should be able to run the TinyGo command:

    ```
    tinygo version
    ```

### Flashing boards

The `tinygo flash` command does not work correctly yet. However you can use `tinygo build` to build the binary, and then manually flash the binary to the target board. See the following examples.

#### Adafruit Circuit Playground Express

- First build the binary in `UF2` format:

    ```shell
    tinygo build -o flash.uf2 -target circuitplay-express examples/blinky1
    ```

- Plug in the Circuit Playground Express board to the USB port.

- Double tap the "RESET" button.

- Once the Circuit Playground Express board appears as a flash drive, copy the `flash.uf2` file to the `CPLAYBOOT` folder.

#### Arduino Nano33 IoT

- You must install the "BOSSA" flashing utility first. You can download it from https://github.com/shumatech/BOSSA/releases/download/1.9.1/bossa-x64-1.9.1.msi

- During the installation, you should choose to put it into `c:\Program Files`

- After the installation, you must add it to your PATH:

    ```shell
    set PATH=%PATH%;"c:\Program Files\BOSSA";
    ```

- Test that you have installed "BOSSA" correctly by running this command:

    ```shell
    bossac --help
    ```

- Now you can build the binary in `BIN` format:

    ```shell
    tinygo build -o flash.bin -target arduino-nano33 path\to\code
    ```

- Plug in the Arduino Nano33 IoT board to the USB port.

- Double tap the "RESET" button.

- Windows 10 should display a notification with the name of the COM port that the board has been assigned to. For example `COM4`.

- Flash the binary to your Arduino Nano33 board by running this command:

    ```shell
    bossac -d -i -e -w -v -R --port=COM4 --offset=0x2000 flash.bin
    ```

## WSL Install

If you want to use TinyGo to compile your own or sample code, you should be able to install the release version directly on your machine using Windows Subsystem for Linux (WSL) by following the "Quick Install" instructions on the Linux page located [here](../linux).

## Docker Install

The other option is to use the Docker image. This has the benefit of making no changes to your system but has a large download and installation size. For instructions on using the Docker image, please see the page [here](../using-docker).
