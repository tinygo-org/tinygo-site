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

### Visual Studio Code

In VS Code, you can edit the file `.vscode/settings.json` in the root of your project. If the `.vscode` directory does not yet exist, create it. It's a normal JSON file where you need to set the `go.toolsEnvVars` property. An example file (again, using the above configuration) is the following:

```json
{
    "go.toolsEnvVars": {
        "GOROOT": "/home/user/.cache/tinygo/goroot-go1.14-f930d5b5f36579e8cbd1c139012b3d702281417fb6bdf67303c4697195b9ef1f-syscall",
        "GOFLAGS": "-tags=cortexm,baremetal,linux,arm,nrf51822,nrf51,nrf,microbit,tinygo,gc.conservative,scheduler.tasks"
    }
}
```

After creating or modifying this file, you will likely need to restart VS Code to apply these settings.

Alternatively, there is a [VSCode addon](https://marketplace.visualstudio.com/items?itemName=tinygo.vscode-tinygo) that will create these settings for you.

### Other IDEs

Other IDEs will likely need a different setup. You can try starting them with these environment variables set in your shell or configuring these environment variables somewhere in your Go language server settings.

As an example, this is an alternative way to start VS Code with the settings above (using `bash`):

```
GOROOT=/home/user/.cache/tinygo/goroot-go1.14-f930d5b5f36579e8cbd1c139012b3d702281417fb6bdf67303c4697195b9ef1f-syscall GOFLAGS=-tags=cortexm,baremetal,linux,arm,nrf51822,nrf51,nrf,microbit,tinygo,gc.conservative,scheduler.tasks code
```

### Using tinygo-edit

There is a CLI tool called [tinygo-edit](https://github.com/sago35/tinygo-edit) you can use it to Gather the needed build flags and starting the editor using the correct environment variables. Using the CLI you don't need to do the steps manualy.

### Using an alias

Another alternative to automate the process of fetching the needed environment variables and starting your editor, you could also create an alias.

#### Ubuntu Example using vscode

```bash
export VISUAL=code
export EDITOR="$VISUAL"
alias startTinyGoArduino="GOOS=linux GOARCH=arm GOFLAGS=-tags=$(tinygo info arduino|grep 'build tags'|awk -F: '{print $2}' | sed -e 's/^[[:space:]]*//'|sed -e 's/[[:space:]]/,/g') $EDITOR"
```

## TinyGo Drivers

There are already lot's of drivers for common hardware. See [this](https://github.com/tinygo-org/drivers) for more information.

### Using drivers with Go modules

When you're using Go modules, drivers are automatically added when you import a driver (such as `tinygo.org/x/drivers/ws2812`).

### Install drivers using GOPATH

When you're not using Go modules, you need to download the drivers separately. This works like any Go package:

> go get -d tinygo.org/x/drivers
