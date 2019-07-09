# Hugo

# Links

* [Hugo](https://gohugo.io/)

# Install

```
// macOS
brew install hugo
```
# Quick Start

[Quick Start | Hugo](https://gohugo.io/getting-started/quick-start/)

```
hugo new site <dir>
```
# Theme

[Complete List | Hugo Themes](https://themes.gohugo.io/)

```
cd <hugodir>
git init
git submodule add https://github.com/budparr/gohugo-theme-ananke.git themes/ananke
echo 'theme = "ananke"' >> config.toml
```

# Serve

```
hugo server -D
```

