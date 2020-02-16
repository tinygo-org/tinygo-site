---
title: "IDE Integration"
weight: 8
---

## GoPath

In order to get IDE support like autocomplete you have to add the tinygo src path to your [GOPATH](https://github.com/golang/go/wiki/GOPATH).
After adding the tinygo source path to the GOPATH go will know, where to look for the packages like `machine`

### Move tinygo src files
Create the following folder structure inside the tinygo src folder

> /github.com/tinygo-org/tinygo

Now move all folders & files to

> path/to/your/tinygo/src/github.com/tinygo-org/tinygo


### Ubuntu Example
You can just append the path to your tinygo installation in your GOPATH.

> export GOPATH=$GOPATH:/usr/local/tinygo/src/github.com/tinygo-org/tinygo

