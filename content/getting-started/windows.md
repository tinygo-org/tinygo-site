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

- You MUST use Go 1.14.x+ with the Windows 10 native install of TinyGo.

    - If you have not installed it yet, you can get it from https://golang.org/dl/

    - Choose the download link for Microsoft Windows, Windows 7 or later, Intel 64-bit processor.

- Download the TinyGo binary for Windows file from https://github.com/tinygo-org/tinygo/releases/download/v0.13.1/tinygo0.13.1.windows-amd64.zip

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

The other option is to use the Docker image. This has the benefit of making no changes to your system but has a large download and installation size. For instructions on using the Docker image, please see the page [here](../using-docker).
