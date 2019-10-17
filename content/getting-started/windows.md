---
title: "Windows"
date: 2019-04-15T11:41:45+02:00
draft: false
weight: 3
---

This page has information on how to install and use TinyGo on Windows 10.

## Quick Install

If you want to use TinyGo to compile your own or sample code, you should be able to install the release version directly on your machine using Windows Subsystem for Linux (WSL) by following the "Quick Install" instructions on the Linux page located [here](../linux).

## Docker Install

The other option is to use the Docker image. This has the benefit of making no changes to your system but has a large download and installation size. For instructions on using the Docker image, please see the page [here](../using-docker).

## Windows Native Install - EXPERIMENTAL

We now have an experimental native install for Windows 10.

VERY IMPORTANT NOTE: You cannot yet create Windows binary programs using TinyGo, only MCU and WASM targets.

- You MUST use Go 1.12.x with the experimental Windows 10 native install of TinyGo.

- Install LLVM 8 from http://releases.llvm.org/download.html#8.0.1

    - Choose the pre-built binary for Windows (64-bit).

    - During the installation make sure you choose the option to add LLVM to your PATH.

- Now, download the TinyGo binary for Windows file from https://github.com/tinygo-org/tinygo/releases/download/v0.9.0/tinygo0.9.0.windows-amd64.zip

- Decompress the file like this:

    ```shell
    PowerShell Expand-Archive -Path "c:\Downloads\tinygo0.9.0.windows-amd64.zip" -DestinationPath "c:\tinygo"
    ```

- You will need to add `C:\tinygo\bin` to your PATH.

    ```shell
    set PATH=%PATH%;C:\tinygo\bin;
    ```

- Now you should be able to run the TinyGo command:

    ```
    tinygo version
    ```
