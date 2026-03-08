---
title: "Helix"
weight: 9
---

The Helix editor provides [configuration](https://docs.helix-editor.com/languages.html#languagestoml-files)
options for customizing LSP servers at a global and per project level. You
can override the Go LSP configuration with environment variables that enable
`tinygo` to work with the Helix editor.

The environment variables needed to be set are `GOROOT`, `GOARCH`, `GOOS`,
and `GOFLAGS`, which will enable `gopls` to find and understand the `machine`
package. The correct values for these variables can be determined using the
`tinygo info <platform>` subcommand. For example:

```shell
$ tinygo info arduino
LLVM triple:       avr
GOOS:              linux
GOARCH:            arm
build tags:        avr baremetal linux arm atmega328p atmega avr5 arduino tinygo purego osusergo math_big_pure_go gc.conservative scheduler.none serial.uart tinygo.unicore
garbage collector: conservative
scheduler:         none
cached GOROOT:     /home/user/.cache/tinygo/goroot-da7db6571e7c2710512ce772df3d2bea4770b700224fdc1ed6bb3c6d38aa4082
```

Create a `.helix/languages.toml` file in the root of your project directory, or
in `~/.config/helix/languages.toml` if you want this override to apply globally.

In this file, create a `tinygo` language-server specification, and then override
the default `go` language to make use of the `tinygo` specific LSP.

```toml
[language-server.tinygo]
command = "gopls"
environment = { "GOROOT" = "/home/user/.cache/tinygo/goroot-da7db6571e7c2710512ce772df3d2bea4770b700224fdc1ed6bb3c6d38aa4082", "GOOS" = "linux", "GOARCH" = "arm", "GOFLAGS" = "-tags=avr,baremetal,linux,arm,atmega328p,atmega,avr5,arduino,tinygo,purego,osusergo,math_big_pure_go,gc.conservative,scheduler.none,serial.uart,tinygo.unicore" }

[[language]]
name = "go"
language-servers = [{name="tinygo"}]
```

Note that the long `environment` mapping must be on one line - this is a quirk
of the TOML configuration language.

This configuration inherits and overrides the default `go` language in Helix
configuration to make use of our custom `tinygo` LSP definition - one that uses
`gopls` preconfigured with the necessary environment variables for understanding
your `tinygo` project.
