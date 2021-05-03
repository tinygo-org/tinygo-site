---
title: "IDE Integration"
weight: 8
---

IDEs need to have certain environment variables set before they work with TinyGo: `GOROOT` and `GOFLAGS`. You can determine the correct values from the `tinygo info` command (starting with TinyGo 0.15).

To get the correct values for these flags, run the following command. Here it is for `microbit`, replace it with something else if you need it.
    
```bash
$ tinygo info microbit
LLVM triple:       armv6m-none-eabi
GOOS:              linux
GOARCH:            arm
build tags:        cortexm baremetal linux arm nrf51822 nrf51 nrf microbit tinygo gc.conservative scheduler.tasks
garbage collector: conservative
scheduler:         tasks
cached GOROOT:     /home/user/.cache/tinygo/goroot-go1.14-f930d5b5f36579e8cbd1c139012b3d702281417fb6bdf67303c4697195b9ef1f-syscall
```

  * The `GOROOT` value needed is the cached GOROOT given in the output (`/home/user/.cache/tinygo/goroot-go1.14-f930d5b5f36579e8cbd1c139012b3d702281417fb6bdf67303c4697195b9ef1f-syscall` in the example).

  * The `GOFLAGS` value can be determined from the build tags. Take the entire list of build tags, replace spaces with commas, and add `-tags=` in front. For this example, the output would be `-tags=cortexm,baremetal,linux,arm,nrf51822,nrf51,nrf,microbit,tinygo,gc.conservative,scheduler.tasks`.

Now you need to configure your IDE with these values.

### Using tinygo-edit

There is a CLI tool called [tinygo-edit](https://github.com/sago35/tinygo-edit) you can use it to Gather the needed build flags and starting the editor using the correct environment variables. Using the CLI you don't need to do the steps manualy.

### Using an alias

Another alternative to automate the process of fetching the needed environment variables and starting your editor, you could also create an alias.

## TinyGo Drivers

There are already lot's of drivers for common hardware. See [this](https://github.com/tinygo-org/drivers) for more information.

### Using drivers with Go modules

When you're using Go modules, drivers are automatically added when you import a driver (such as `tinygo.org/x/drivers/ws2812`).

### Install drivers using GOPATH

When you're not using Go modules, you need to download the drivers separately. This works like any Go package:

> go get -d tinygo.org/x/drivers
