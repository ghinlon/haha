# Plugins

# Links

# vim-easy-align

[junegunn/vim-easy-align: A Vim alignment plugin](https://github.com/junegunn/vim-easy-align)


# Automatic input method switching for vim 

* [lyokha/vim-xkbswitch: vim plugin for automatic keyboard layout switching in insert mode](https://github.com/lyokha/vim-xkbswitch)  
// Mac
* [vovkasm/input-source-switcher: Command line input source switcher for Mac.](https://github.com/vovkasm/input-source-switcher)  
// linux
* [GitHub - ierton/xkb-switch: Switch your X keyboard layouts from the command line](https://github.com/ierton/xkb-switch)  

# Color

* [colorwat.sh](http://colorswat.ch/vim/list?o=updated_at)
* [cocopon/iceberg.vim: Dark blue color scheme for Vim and Neovim](https://github.com/cocopon/iceberg.vim)
* [terminal.sexy - Terminal Color Scheme Designer](https://terminal.sexy/)

# Test Python Support

If `:echo has("python3")` returns `1`, then you have python 3 support; 

# Find

* [junegunn/fzf: A command-line fuzzy finder](https://github.com/junegunn/fzf)

同类的还有LeaderF, skim

deprecated: ctrip

# Regexp

* [BurntSushi/ripgrep: ripgrep recursively searches directories for a regex pattern](https://github.com/BurntSushi/ripgrep)

```sh
brew install ripgrep

$ sudo yum-config-manager --add-repo=https://copr.fedorainfracloud.org/coprs/carlwgeorge/ripgrep/repo/epel-7/carlwgeorge-ripgrep-epel-7.repo
$ sudo yum install ripgrep

$ sudo apt-get install ripgrep
```

deprecated: ack

# Miscellany

* [vim-airline/vim-airline: lean & mean status/tabline for vim that's light as air](https://github.com/vim-airline/vim-airline)
* [majutsushi/tagbar: Vim plugin that displays tags in a window, ordered by scope](https://github.com/majutsushi/tagbar)
* [vim-markdown](https://github.com/plasticboy/vim-markdown) 

  * zr: reduces fold level throughout the buffer
  * zR: opens all folds

  vim-markdown automatically insert the indent. By default, the number of spaces of indent is 4. If you'd like to change the number as 0, just write:

  `let g:vim_markdown_new_list_item_indent = 0`

