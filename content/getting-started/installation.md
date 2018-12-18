---
title: "Installation"
weight: 2
---

Now that you have the requirements, now download the sources. This may take a while.

    go get -u github.com/aykevl/tinygo

If you get an error like this::

    /usr/local/go/pkg/tool/linux_amd64/link: running g++ failed: exit status 1
    /usr/bin/ld: error: cannot find -lLLVM-7
    cgo-gcc-prolog:58: error: undefined reference to 'LLVMVerifyFunction'
    cgo-gcc-prolog:80: error: undefined reference to 'LLVMVerifyModule'
    [...etc...]

Or like this::

    ../go-llvm/analysis.go:17:93: fatal error: llvm-c/Analysis.h: No such file or directory
     #include "llvm-c/Analysis.h" // If you are getting an error here read bindings/go/README.txt

It means something is wrong with your LLVM installation. Make sure LLVM 7 is
installed (Debian package ``llvm-7-dev``). If it still doesn't work, you can
try running:

    cd $GOPATH/github.com/aykevl/go-llvm
    make config

And retry:

    go install github.com/aykevl/tinygo
