# [mpv.io](https://mpv.io/)

# Links

* [Tutorial: Playing around with MPlayer - Linux Academy Blog](https://linuxacademy.com/blog/linux/tutorial-playing-around-with-mplayer/)


# Install

```
brew cask install mpv
```

# Subtitle

```
       --sub-auto=<no|exact|fuzzy|all>, --no-sub-auto
              Load additional subtitle files matching the video filename. The parameter specifies how external subtitle files are matched. exact is enabled by default.

              no     Don't automatically load external subtitle files.

              exact  Load the media filename with subtitle file extension (default).

              fuzzy  Load all subs containing media filename.

              all    Load all subs in the current and --sub-file-paths directories.

       --sub-file-paths=<path-list>
              Specify extra directories to search for subtitles matching the video.  Multiple directories can be separated by ":" (";" on Windows).  Paths can be rela-
              tive or absolute. Relative paths are interpreted relative to video file directory.  If the file is a URL, only absolute paths and sub configuration  sub-
              directory will be scanned.

                 Example

                        Assuming  that  /path/to/video/video.avi  is  played  and --sub-file-paths=sub:subtitles is specified, mpv searches for subtitle files in these
                        directories:

                 o /path/to/video/

                 o /path/to/video/sub/

                 o /path/to/video/subtitles/

                 o the sub configuration subdirectory (usually ~/.config/mpv/sub/)

              This is a list option. See List Options for details.
```

cfgFile: `~/.config/mpv/mpv.conf`

```
sub-auto=all
sub-file-paths=sub:subtitles
```
