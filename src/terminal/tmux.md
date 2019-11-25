# tmux

# Links

* [A Gentle Introduction to tmux – Hacker Noon](https://hackernoon.com/a-gentle-introduction-to-tmux-8d784c404340)
* [Home · tmux/tmux Wiki](https://github.com/tmux/tmux/wiki)
* [Learn tmux | Terminal Multiplexer Tutorials by thoughtbot](https://thoughtbot.com/upcase/tmux)
* [What's your prefix key? : tmux](https://www.reddit.com/r/tmux/comments/4a1694/whats_your_prefix_key/)

# Install

```
// MacOS
brew install tmux
```

# Basic

```
Prefix: ctrl-b

tmux new
prefix d		detach tmux
prefix ?		help


tmux ls		list sessions
tmux attach-session -t 3
tmux a #	attach the last session

// panes

tmux new -s {name}	start a new session with a specific name
tmux a -t {name}	attach a session by using the name

prefix "		split a pane horizontally
prefix %		split a pane vertically
prefix o		Cycle through panes
prefix ;		Cycle between previous and current
prefix x		Kill current pane
prefix z		make a pane full screen, hit again to shrink it back to it's previous size

prefix :resize-pane -U/D/L/R {count} 	resize-pane up/down/left/right {count} lines

// windows

prefix c		create a new window
prefix {num}		Switch to window {num}
prefix ,		rename current window

tmux kill-server	kill tmux server, along with all sessions.
```

# Configuration

file: `.tmux.conf`

load config: `prefix :source-file .config/tmux/tmux.conf`

```
# remap prefix from 'C-b' to 'C-a'
unbind C-b
set-option -g prefix C-g
bind-key C-g send-prefix

# Start window numbering at 1
set -g base-index 1
```

# TPM

[GitHub - tmux-plugins/tpm: Tmux Plugin Manager](https://github.com/tmux-plugins/tpm)

Install

```
git clone https://github.com/tmux-plugins/tpm ~/.tmux/plugins/tpm

prefix + I 		// Install plugin
prefix + U 		// Update plugin

# type this in terminal if tmux is already running
tmux source ~/.tmux.conf
```

# reorder tmux window

* [How do I reorder tmux windows? - Super User](https://superuser.com/questions/343572/how-do-i-reorder-tmux-windows)

to let window number 3 and window number 1 swap their positions:

```go
swap-window -s 3 -t 1
```
`
