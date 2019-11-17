# [Package template](https://golang.org/pkg/text/template/)

# Links

# Text and spaces 

* `{{- `: all trailing white space is trimmed from the immediately preceding text.
* ` -}}`: all leading white space is trimmed from the immediately following text. 
* `{{-3}}`: parses as an action containing the number -3

For instance:

```
"{{23 -}} < {{- 45}}"
```

will generate:

```
"23<45"
```

# func (*Template) Funcs  

Funcs adds the elements of the argument map to the template's function map. It
must be called before the template is parsed. It panics if a value in the map
is not a function with appropriate return type. However, it is legal to
overwrite elements of the map. The return value is the template, so calls can
be chained. 

```go
func (t *Template) Funcs(funcMap FuncMap) *Template
```


