---
title: "IDE Integration"
weight: 8
---


## Downloading the source

Either git clone the tinygo source from.

> https://github.com/tinygo-org/tinygo

## Go mod init in your project

Navigate to your project and type:

> go mod init

## Replace needed imports

Replace each std package you need using the replace keyword in the go.mod file.

So if you want the machine package to be resolved replace it using.

> replace machine => /path/to/tinygo/machine

Repeat this step for every package, which needs to be resolved in your project.

### Example

```go
module github.com/Nerzal/tinygo-playground

go 1.14

replace machine => /home/tobias/go/src/github.com/tinygo-org/tinygo/src/machine

require (
    machine v0.0.0-00010101000000-000000000000
    tinygo.org/x/drivers v0.13.0
)

```

## Add go module files to the packages

In order to get your package resolved, you also have to add a go mod file for the packages.

So navigate to `/path/to/tinygo/src`
And then do `go mod init`

### Example

```bash
 cd ~/go/src/github.com/tinygo-org/tinygo/src/machine
 go mod init
```

## GoPath

In order to get IDE support like autocomplete you have to add the tinygo src path to your [GOPATH](https://github.com/golang/go/wiki/GOPATH).
After adding the tinygo source path to the GOPATH go will know, where to look for the packages like `machine`

### Ubuntu Example

You can just append the path to your tinygo installation in your GOPATH.

> export GOPATH=$GOPATH:/path/to/your/tinygo

## Starting your editor with variables

The last step needed to get full code completion support is to start your editor with environment variables.

You could also just set the environment variables for your complete environment, but this would interfere you when, writing normal go code.

So we use another way.

### Gather the needed information

We need different environment variables for different microcontrollers.

You can use `tinygo info controllerName`

**Example**
To gather information needed to work with an Arduino use:

> tinygo info arduino

```bash
$ tinygo info arduino
LLVM triple:       avr-unknown-unknown
GOOS:              linux
GOARCH:            arm
build tags:        avr baremetal linux arm atmega328p atmega avr5 arduino tinygo gc.conservative scheduler.none
garbage collector: conservative
scheduler:         none
```

So now you now, that.

1. GOOS needs to be set to linux
2. GOARCH needs to be set to arm
3. GOFLAGS needs to be set to `-tags=avr,baremetal,linux,arm,atmega328p,atmega,avr5,arduino,tinygo,gc.conservative,scheduler.none`

Important is that you need to comma separate the tags.

### Start the editor

The following example should work for all editors. And more or less all operating systems. The syntax may vary depending on your os/shell.

VSCode Example:

> export GOOS=linux; export GOARCH=arm; export GOFLAGS=-tags=avr,baremetal,linux,arm,atmega328p,atmega,avr5,arduino,tinygo,gc.conservative,scheduler.none; code

### Note

This process has only been tested with the [gopls](https://github.com/golang/tools/blob/master/gopls/doc/user.md) language server. It might or might not work with other language servers.

To install gopls follow the instructions in the link above.

## TinyGo Drivers

There are already lot's of drivers for common hardware. See [this](https://github.com/tinygo-org/drivers) for more information.

### Install drivers

> go get tinygo.org/x/drivers
