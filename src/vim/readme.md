# VIM

# Block Insert and Change text

Insert/Change/Replace: 

	CTRL-V -> {motion} -> I/c/r -> Ctrl-[

# Multiple Change

* Search it with '/' or '*' or any search method
* Change word with gn motion. e.g. cgnfoo<esc>
* Repeat via . command

# vim takes a very long time to start up

First, try running Vim with the following command: `vim -X`

If `-X` helps, you can get the same effect by adding the following line to your
vimrc file:

`set clipboard=exclude:.*`

# Turning off auto indent when pasting text into vim

`set paste`

* [configuration - Turning off auto indent when pasting text into vim - Stack
Overflow](https://stackoverflow.com/questions/2514445/turning-off-auto-indent-when-pasting-text-into-vim)

```vim
Mac users can avoid auto formatting by reading directly from the pasteboard
with:

:r !pbpaste

// linux
:r !xsel -p
```
