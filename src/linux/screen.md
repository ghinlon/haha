# GNU Screen

# Links

* [How To Use Linux Screen | Linuxize](https://linuxize.com/post/how-to-use-linux-screen/)

# Start

```
screen -S session_name
```
# Basic Config

Config File `.screenrc`:

```
escape ^jj

startup_message off
vbell off

hardstatus on
altscreen on
defscrollback 10000
```

# Basic Command

```
prefix c 	Create a new window (with shell)
prefix " 	List all window
prefix 0 	Switch to window 0 (by number )
prefix A 	Rename the current window
prefix S 	Split current region horizontally into two regions
prefix | 	Split current region vertically into two regions
prefix tab 	Switch the input focus to the next region
prefix Ctrl+a 	Toggle between the current and previous region
prefix Q 	Close all regions but the current one
prefix X 	Close the current region

prefix d 	Detach from Linux Screen Session
```
```
screen -r	Reattach to a Linux Screen
screen -ls	List
```
