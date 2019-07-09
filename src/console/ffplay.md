# ffplay

# Links

* [ffplay Documentation](https://ffmpeg.org/ffplay.html#While-playing)

```
q, ESC
    Quit.
f
    Toggle full screen.
p, SPC
    Pause.
m
    Toggle mute.
9, 0
    Decrease and increase volume respectively.
/, *
    Decrease and increase volume respectively.
a
    Cycle audio channel in the current program.
v
    Cycle video channel.
t
    Cycle subtitle channel in the current program.
c
    Cycle program.
w
    Cycle video filters or show modes.
s
    Step to the next frame. 
```

infile subtitle:

```
infile=<xxx.mp4>; ffplay -vf subtitles=$infile:si=0 $infile
```
