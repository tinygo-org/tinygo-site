---
title: "VS Code"
weight: 9
---

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


#### Ubuntu Example using vscode

```bash
export VISUAL=code
export EDITOR="$VISUAL"
alias startTinyGoArduino="GOOS=linux GOARCH=arm GOFLAGS=-tags=$(tinygo info arduino|grep 'build tags'|awk -F: '{print $2}' | sed -e 's/^[[:space:]]*//'|sed -e 's/[[:space:]]/,/g') $EDITOR"
```
