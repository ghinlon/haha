# Go

# Great Articles

* [Less is exponentially more](https://commandcenter.blogspot.com/2012/06/less-is-exponentially-more.html)  
	`If C++ and Java are about type hierarchies and the taxonomy of types, Go is about composition.`
 
# Tutorial

* [The actor model in 10 minutes](https://www.brianstorti.com/the-actor-model/)
* [Ways To Do Things - Peter Bourgon - Release Party #GoSF - YouTube](https://www.youtube.com/watch?v=LHe1Cb_Ud_M&t=15m45s)
* [Peter Bourgon Â· Go: Best Practices for Production Environments](https://peter.bourgon.org/go-in-production/#formatting-and-style)
* [CodeReviewComments Â· golang/go Wiki](https://github.com/golang/go/wiki/CodeReviewComments)
* [Build You Own Web Framework In Go | Nicolas Merouze](https://www.nicolasmerouze.com/build-web-framework-golang)(**Recommendation**)
* [Rethinking Visual Programming with Go · divan's blog](https://divan.dev/posts/visual_programming_go/)
* [Go Proverbs](https://go-proverbs.github.io/)
* [A Tour of Go](https://tour.golang.org/welcome/1)
* [Go by Example](https://gobyexample.com/)
* [Learn Go in Y Minutes](https://learnxinyminutes.com/docs/go/)
* [enocom/gopher-reading-list: A curated selection of blog posts on Go](https://github.com/enocom/gopher-reading-list)
* [golang/go Wiki](https://github.com/golang/go/wiki)
* [Learn · golang/go Wiki · GitHub](https://github.com/golang/go/wiki/Learn)
* [GoLove的博客](http://www.cnblogs.com/golove/)
* [Network Programming with Go (golang)](https://ipfs.io/ipfs/QmfYeDhGH9bZzihBUDEQbCbTc5k5FZKURMUoUvfmc27BwL/index.html)
* The [Go Blog](https://blog.golang.org/) has a large archive of informative Go articles. 
* Once you have Go installed, the [Go Documentation](https://golang.org/doc/) is a great place to continue. It contains references, tutorials, videos, and more. 
* If you need help with the standard library, see the [package reference](https://golang.org/pkg/). For help with the language itself, you might be surprised to find the [Language Spec](https://golang.org/ref/spec) is quite readable. 
* The [First Class Functions in Go](https://golang.org/doc/codewalk/functions/) codewalk gives an interesting perspective on Go's function types.
* [Writing Web Applications](https://golang.org/doc/articles/wiki/)
* [“Introduction to bufio package in Golang” @mlowicki](https://medium.com/golangspec/introduction-to-bufio-package-in-golang-ad7d1877f762)
* [“In-depth introduction to bufio.Scanner in Golang” @mlowicki](https://medium.com/golangspec/in-depth-introduction-to-bufio-scanner-in-golang-55483bb689b4)
* [Arrays, slices (and strings): The mechanics of 'append'](https://blog.golang.org/slices)
* [Go Slices: usage and internals](https://blog.golang.org/go-slices-usage-and-internals)
* [Go Data Structures](https://research.swtch.com/godata)
* ["Slice Tricks" Wiki page ](https://golang.org/wiki/SliceTricks)
* [Strings, bytes, runes and characters in Go](https://blog.golang.org/strings)
* [Text normalization in Go](https://blog.golang.org/normalization)
* [The Absolute Minimum Every Software Developer Absolutely, Positively Must Know About Unicode and Character Sets (No Excuses!).](http://www.joelonsoftware.com/articles/Unicode.html) 
* [Establishing and Running Large Scale Internet Based (TCP/IP) Services](http://mars.netanya.ac.il/~unesco/cdrom/booklet/HTML/index.html)

# Install

[Download Site](https://golang.org/dl/)

```
sudo tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz

export GOPATH=~/go
export GOBIN=${GOPATH}/bin
export PATH=${GOBIN}:/usr/local/go/bin:${PATH}
```

# Editor

* [neovim/neovim](https://github.com/neovim/neovim)  

## Install

macOS

```
tar xzvf nvim-macos.tar.gz
export PATH=~/nvim-osx64/bin/nvim:$PATH
nvim -version
```

## Plugins

* [fatih/vim-go](https://github.com/fatih/vim-go)  
* [fatih/vim-go-tutorial: Tutorial for vim-go](https://github.com/fatih/vim-go-tutorial)
* [ncm2/ncm2: Completion Framework for Neovim](https://github.com/ncm2/ncm2)  
  Slim, Fast and Hackable Completion Framework for Neovim.**Really** fast.  
* [mdempsky/gocode: An autocompletion daemon for the Go programming language](https://github.com/mdempsky/gocode)
* [sebdah/vim-delve: Neovim / Vim integration for Delve](https://github.com/sebdah/vim-delve)

deprecated

* [zchee/deoplete-go: Asynchronous Go completion for Neovim. deoplete source for Go.](https://github.com/zchee/deoplete-go)  

more about completion:

* [Best vim completion manager? : vim](https://www.reddit.com/r/vim/comments/9a8c3m/best_vim_completion_manager/)

## Install

install vim-plug

```
// Neovim
curl -fLo ~/.local/share/nvim/site/autoload/plug.vim --create-dirs \
    https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim

// then
:PlugInstall
```

# Misc

If `:echo has("python3")` returns 1, then you have python 3 support; 

# xcrun error

```
xcrun: error: invalid active developer path (/Library/Developer/CommandLineTools), missing xcrun at: /Library/Developer/CommandLineTools/usr/bin/xcrun 

(1 of 1): xcrun: error: invalid active developer path (/Library/Developer/CommandLineTools), missing xcrun at: /Library/Developer/CommandLineTools/usr/bin/xcrun
```

Shooting: 

Install CommandLineTools

