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
```

# Usage

```
youtube-dl [OPTIONS] URL

-f best 					// best resolution
-f bestvideo+bestaudio 		// best resolution
```

# Downoad Subtitle

* [subtitle-options](https://github.com/ytdl-org/youtube-dl#subtitle-options)

```
youtube-dl --write-auto-sub --skip-download
```

