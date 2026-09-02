---
title: "UEFI applications"
weight: 2
description: |
  How to use TinyGo to create UEFI applications that the firmware starts before an operating system.
---

TinyGo can compile programs into UEFI applications. A UEFI application is a PE32+ executable. The firmware loads and starts it before an operating system starts. You can use this to write boot loaders, firmware test tools, and diagnostic tools in Go.

UEFI support is new in TinyGo 0.42.0. The target name is `uefi-amd64`. Only 64-bit x86 is supported.

## Write the program

The `device/uefi` package gives you access to the UEFI system table, the boot services, and the console. The `println` function writes to the firmware console, so a first program can be very small:

```go
package main

import (
	"device/uefi"
)

func main() {
	println("Hello from TinyGo UEFI")
	println("Running forever...")

	for {
		uefi.CpuPause()
	}
}
```

When `main` returns, the application gives control back to the firmware. The firmware then continues the boot sequence. The loop in the example above keeps the application on the screen.

## Build the program

Build the program with the `uefi-amd64` target. Give the output file the name that the firmware looks for:

```shell
tinygo build -target=uefi-amd64 -o BOOTX64.EFI main.go
```

The result is an EFI application:

```shell
$ file BOOTX64.EFI
BOOTX64.EFI: PE32+ executable (EFI application) x86-64, for MS Windows, 15 sections
```

## Run the program in QEMU

QEMU with the OVMF firmware is the quickest way to test a UEFI application. Install `qemu-system-x86_64` and the `ovmf` package for your operating system.

The firmware looks for the removable media boot path on the EFI system partition. Put the executable in the `EFI/BOOT` directory:

```shell
mkdir -p EFI/BOOT
tinygo build -target=uefi-amd64 -o EFI/BOOT/BOOTX64.EFI main.go
```

Then start QEMU. The `fat:rw:` prefix tells QEMU to show the current directory to the guest as a FAT file system, so you do not have to make a disk image:

```shell
qemu-system-x86_64 \
    -machine q35 \
    -m 1024 \
    -smp 2 \
    -drive if=pflash,format=raw,readonly=on,file=/usr/share/OVMF/OVMF_CODE_4M.fd \
    -drive format=raw,file=fat:rw:$PWD \
    -nographic
```

The path to the OVMF firmware file is different on each operating system. On Debian and Ubuntu it is `/usr/share/OVMF/OVMF_CODE_4M.fd`. On Fedora it is `/usr/share/edk2/ovmf/OVMF_CODE.fd`. Use `-nographic` to send the firmware console to your terminal. To stop QEMU, press `Ctrl-A` and then `X`.

## Run the program on hardware

To run the application on a real machine, copy it to a USB drive:

- Format the USB drive with a FAT32 file system.
- Copy the executable to `EFI/BOOT/BOOTX64.EFI` on the drive.
- Start the machine from the USB drive.

You must turn off Secure Boot, or sign the executable with a key that the firmware accepts. Firmware does not start an unsigned application while Secure Boot is on.

## Use the UEFI services

The `device/uefi` package gives you the system table and the boot services:

- `uefi.ST()` returns the `EFI_SYSTEM_TABLE`.
- `uefi.BS()` returns the `EFI_BOOT_SERVICES`.
- `uefi.GetImageHandle()` returns the image handle of your application.

The `machine` package makes the most useful parts available with shorter names. It has the `EFI_STATUS` values, the error values, `machine.GetTime`, `machine.ConsoleOut`, and `machine.StandardError`.

`machine.ConsoleOut` returns an `io.Writer` for the firmware console, so you can use the `fmt` package with it. This writer does not change `\n` into `\r\n`. Write `\r\n` if you want the cursor to go back to the start of the line:

```go
package main

import (
	"fmt"
	"machine"
)

func main() {
	out := machine.ConsoleOut()
	fmt.Fprint(out, "Hello from TinyGo UEFI\r\n")

	t, status := machine.GetTime()
	if status != machine.EFI_SUCCESS {
		fmt.Fprintf(out, "could not read the time: %v\r\n", machine.StatusError(status))
		return
	}

	fmt.Fprintf(out, "firmware time is %04d-%02d-%02d %02d:%02d:%02d\r\n",
		t.Year, t.Month, t.Day, t.Hour, t.Minute, t.Second)
}
```

The boot services return an `EFI_STATUS` value instead of a Go error. Compare the value with `uefi.EFI_SUCCESS`, or use `uefi.StatusError` to make an `error` from it.

Goroutines work on this target, because the `tasks` scheduler is the default. The `time` package also works. The runtime reads the time stamp counter for monotonic time, and the UEFI runtime services for the wall clock.

To wait for a UEFI event, use `uefi.WaitForEvent`. This function gives time to the other goroutines while it waits:

```go
var event uefi.EFI_EVENT
status := uefi.BS().CreateEvent(uefi.EVT_TIMER, uefi.TPL_CALLBACK, nil, nil, &event)
if status != uefi.EFI_SUCCESS {
	return uefi.StatusError(status)
}
defer uefi.BS().CloseEvent(event)

// The trigger time is in units of 100ns, so this is one second.
uefi.BS().SetTimer(event, uefi.TimerRelative, 10000000)
uefi.WaitForEvent(event)
```

The `device/uefi` package has methods for only some of the boot services. To use a protocol that the package does not wrap, get the protocol with `HandleProtocol` or `LocateProtocol`. Then call the function pointers in it with `uefi.UefiCall0` to `uefi.UefiCall10`. These functions use the UEFI calling convention.

To convert text between Go strings and the UTF-16 strings that UEFI uses, use the `uefi.StringToCHAR16Z` and `uefi.CHAR16PtrToString` functions.

## Limitations

- Only the `uefi-amd64` target is available. There is no target for 32-bit x86 or for ARM.
- The target uses the leaking garbage collector. The runtime asks the firmware for a heap at start, but it never frees an allocation. Do not make many short-lived allocations in an application that runs for a long time.
- There is no operating system, so the parts of the standard library that need one do not work. Files, networking, and processes are not available.
- The `device/uefi` package wraps a small number of protocols and services. Use `uefi.UefiCall0` to `uefi.UefiCall10` for the rest.
- Console input is not supported yet. The `ConIn` field of the system table has no wrapper, and `os.Stdin` reads nothing.
