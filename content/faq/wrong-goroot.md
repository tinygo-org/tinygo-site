---
title: "Cannot find package \"unsafe\""
weight: 6
---

I get the following error when I run tinygo:

``` text
error: cannot find package "unsafe" in any of:
    /usr/local/Cellar/go/1.11.1/libexec/src/unsafe (from $GOROOT)
    ...
```

You can probably solve this issue by defining the GOROOT environment while running tinygo:

``` shell
GOROOT=$(go env GOROOT) tinygo ...
```

If you have to extend the PATH, you can also define both variables:

``` shell
PATH=/usr/local/Cellar/llvm/8.0.0/bin:$PATH GOROOT=$(go env GOROOT) tinygo ...
```
