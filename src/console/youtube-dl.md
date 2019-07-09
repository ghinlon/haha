# [youtube-dl](https://rg3.github.io/youtube-dl/)

# Links

* [youtube-dl - ArchWiki](https://wiki.archlinux.org/index.php/Youtube-dl)


# Install

```
brew install youtube-dl
```

cfgFile: `~/.config/youtube-dl/config`

```
--proxy socks5://127.0.0.1:1080

--ignore-errors
# --no-playlist
#

# Save in ~/Videos
-o ~/dl/youtube/%(title)s.%(ext)s

# Prefer 1080p or lower resolutions
-f bestvideo[ext=mp4][height<1200]+bestaudio[ext=m4a]/bestvideo[ext=webm][height<1200]+bestaudio[ext=webm]/bestvideo[height<1200]+bestaudio/best[height<1200]/best
```

# Usage

```
youtube-dl [OPTIONS] URL
```
