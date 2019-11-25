# Zsh

# Links

* [Zsh - ArchWiki](https://wiki.archlinux.org/index.php/Zsh)
* [GitHub - zplug/zplug: A next-generation plugin manager for zsh](https://github.com/zplug/zplug)
* [Beautifying your terminal with Zsh, Prezto, & Powerlevel9k](https://medium.com/@oldwestaction/beautifying-your-terminal-with-zsh-prezto-powerlevel9k-9e8de2023046)
* [Why zsh is Cooler than Your Shell](https://www.slideshare.net/brendon_jag/why-zsh-is-cooler-than-your-shell?next_slideshow=1)
* [Comparison of ZSH frameworks and plugin managers](https://gist.github.com/laggardkernel/4a4c4986ccdcaf47b91e8227f9868ded)
* [goreliu/zshguide: Zsh 开发指南](https://github.com/goreliu/zshguide)

# Install

```
brew install zsh zsh-completions
```

# Installing Fonts

* [GitHub - powerline/fonts: Patched fonts for Powerline users.](https://github.com/powerline/fonts)
* [ryanoasis/nerd-fonts: Iconic font aggregator, collection, and patcher. 40+ patched fonts, over 3,600 glyph/icons, includes popular collections such as Font Awesome & fonts such as Hack](https://github.com/ryanoasis/nerd-fonts)

```
git clone https://github.com/powerline/fonts.git --depth=1
cd fonts && ./install.sh

// Install nerdfont
brew tap homebrew/cask-fonts
brew cask install font-hack-nerd-font
```

# Comparison of Zsh plugin managers

* [Comparison of ZSH frameworks and plugin managers · GitHub](https://gist.github.com/laggardkernel/4a4c4986ccdcaf47b91e8227f9868ded)

# zplugin

```
mkdir -p ~/.zplugin
git clone https://github.com/zdharma/zplugin.git ~/.zplugin/bin
zcompile ~/.zplugin/bin/zplugin.zsh
```

# load completition file

* [Command-line completion | Docker Documentation](https://docs.docker.com/machine/completion/)

