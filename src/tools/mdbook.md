# mdbook

# Links

* [rust-lang-nursery/mdBook: Create book from markdown files. Like Gitbook but implemented in Rust](https://github.com/rust-lang-nursery/mdBook)
* [mdBook - mdBook Documentation](https://rust-lang-nursery.github.io/mdBook/)
* [Support in-repo links, with verification and output-dependent conversion · Issue #431 · rust-lang-nursery/mdBook · GitHub](https://github.com/rust-lang-nursery/mdBook/issues/431)

# 为什么弃用gitbook

脚本语言实现的东西 每个很小很小很小的功能或模块都是一个脚本文件.

所以稍微一点点复杂的功能 每次执行都意味着发生了巨量的硬盘读写

所以就是慢 

脚本语言就做做脚本语言适合做的简单的事情就好了

# Install

This requires at least [Rust](https://www.rust-lang.org/en-US/) 1.20 and Cargo
to be installed. Once you have installed Rust, type the following in the
terminal:

```
cargo install mdbook
```

# Basic Usage

```
mdbook build
mdbook watch
mdbook serve
mdbook clean
```

# Deploying Your Book to GitHub Pages

[Continuous Integration - mdBook Documentation](https://rust-lang-nursery.github.io/mdBook/continuous-integration.html)


I have a question is whether dir `book` is needed to push to repo ?

不需要的，配置文件里的`script`里生成的就会被拿去使用

# 在readme.md里面相对引用子目录里的md

是可以的，不过不能引用`readmd.md`，因为`readmd.md`会被处理成`index.html`，这是
一个特殊情况

[Configuration - mdBook Documentation](https://rust-lang-nursery.github.io/mdBook/format/config.html)

The following preprocessors are available and included by default:

* links: Expand the {{# playpen}} and {{# include}} handlebars helpers in
a chapter.
* index: Convert all chapter files named README.md into index.md. That is
to say, all README.md would be rendered to an index file index.html in the
rendered book.


