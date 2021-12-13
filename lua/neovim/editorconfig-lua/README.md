# Test Neovim plugin using editorconfig3

Please don't use this plugin!

Its made entirely for me to learn making plugins for Neovim. If you are looking for a Neovim plugin to do editorconfig for you I would suggest using the one [Gpanders](https://github.com/gpanders/editorconfig.nvim) made.

## Neovim

Add path so we can load the local module

```sh
# nvim --cmd "set rtp+=./" lua/editorconfig-lua/init.lua
```

From within Neovim we can now load the module and test

```neovim
:lua require('editorconfig-lua').setup({debug = true})
:LoadEditorConfig
:messages
```
