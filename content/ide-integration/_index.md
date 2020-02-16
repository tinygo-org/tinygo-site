---
title: "IDE Integration"
weight: 8
---

## GoPath

In order to get IDE support like autocomplete you have to add the tinygo src path to your [GOPATH](https://github.com/golang/go/wiki/GOPATH).
After adding the tinygo source path to the GOPATH go will know, where to look for the packages like `machine`

### Ubuntu Example
You can just append the path to your tinygo installation in your GOPATH.

>  export GOPATH=$GOPATH:/path/to/your/tinygo

## TinyGo Drivers

There are already lot's of drivers for common hardware. See [this](https://github.com/tinygo-org/drivers) for more information.

### Install drivers

> go get tinygo.org/x/drivers
