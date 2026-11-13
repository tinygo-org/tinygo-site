---
title: "Zed"
weight: 9
---

In Zed, you can edit the file `~/.config/zed/settings.json` (global) or `.zed/settings.json` (project local). An example file (using the `GOROOT` and tags output from running `tinygo info <board>`) is the following (top-level) JSON:

```json
"lsp": {
    "gopls": {
      "initialization_options": {
        "env": {
          "GOROOT": "/home/user/.cache/tinygo/goroot-go1.14-f930d5b5f36579e8cbd1c139012b3d702281417fb6bdf67303c4697195b9ef1f-syscall",
          "GOFLAGS": "-tags=cortexm,baremetal,linux,arm,nrf51822,nrf51,nrf,microbit,tinygo,gc.conservative,scheduler.tasks"
        }
    }
}
```
