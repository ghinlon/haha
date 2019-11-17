# Asynchronous Lint Engine

# Links

* [GitHub - dense-analysis/ale](https://github.com/dense-analysis/ale)

# golang

must config this, default doesn't use `gopls`:

```
let g:ale_linters_explicit = 1
let g:ale_linters = {
\   'go': ['gopls'],
\}
```


