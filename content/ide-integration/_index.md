---
title: "IDE Integration"
weight: 8
---

## GoPath

In order to get IDE support like autocomplete you have to add the tinygo src path to your [GOPATH](https://github.com/golang/go/wiki/GOPATH).
After adding the tinygo source path to the GOPATH go will know, where to look for the packages like `machine`

### Ubuntu Example

> export GOPATH=$GOPATH:/usr/local/tinygo

