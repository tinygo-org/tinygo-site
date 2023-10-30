---
title: "Vim and Neovim"
weight: 9
---

In Vim and Neovim LSP servers can be configured to offer completions and documentation based on the contents of the file you are editing.
These LSP servers (i.e. `gopls` in our case) are spawned by Vim and Neovim themselves. This means we must provide them with the appropriate
environment so that they know where to find the definition of the `machine package`.

This can be done manually before starting Vim or Neovim by running `export` on the shell. However, this can become cumbersome and very
fast. That is why some available plugins automate the setup by providing a single command you can use to configure the target.

An option would be to use [tinygo.vim](https://github.com/sago35/tinygo.vim). This plugin offers the `:TinygoTarget` command within
Vim and Neovim to change the target at will. It is written in Vim Script, Vim's builtin scripting language. However, it is compatible
both with Vim and Neovim. You can find more information on how to install and use it in the project's page linked above.

A second option would be using a Neovim-native plugin: [tinygo.nvim](https://github.com/pcolladosoto/tinygo.nvim). Unlike the previous
option, this plugin is written entirely in Lua, making it only compatible with Neovim. Through the `:TinyGoSetTarget` command you
can also configure the current target for the `machine` package. Information on how to install and use it is available on its project
page, which is also linked above.

At any rate, either plugin offers a complete integration for TinyGo so that you can leverage all the advantages of `gopls`.
