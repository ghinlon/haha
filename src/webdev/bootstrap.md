# Bootstrap

# Links

* [Introduction · Bootstrap](https://getbootstrap.com/docs/4.3/getting-started/introduction/)
* [GitHub - twbs/bootstrap: The most popular HTML, CSS, and JavaScript framework for developing responsive, mobile first projects on the web.](https://github.com/twbs/bootstrap)

# Install

* [Releases · twbs/bootstrap · GitHub](https://github.com/twbs/bootstrap/releases)

```
unzip bootstrap-4.3.1-dist.zip
```

# QuickStart

* [What is Bootstrap and How Do I Use It? – Tania Rascia](https://www.taniarascia.com/what-is-bootstrap-and-how-do-i-use-it/)

Serve it on `/static/`:

```go
fileserver := http.FileServer(http.Dir("./ui/static/bootstrap"))
mux.Handle("/static/", http.StripPrefix("/static", fileserver))
```


CSS

Copy-paste the stylesheet <link> into your <head> before all other stylesheets
to load our CSS.

```
<!-- Bootstrap CSS -->
<link href="/static/css/bootstrap.min.css" rel="stylesheet" />
```

JS

Many of our components require the use of JavaScript to function. Specifically,
they require jQuery, Popper.js, and our own JavaScript plugins. Place the
following `<script>`s near the end of your pages, right before the closing
`</body>` tag, to enable them. jQuery must come first, then Popper.js, and then
our JavaScript plugins.


```
<!-- Optional JavaScript -->
<!-- jQuery first, then Popper.js, then Bootstrap JS -->
<script src="https://code.jquery.com/jquery-3.3.1.slim.min.js" integrity="sha384-q8i/X+965DzO0rT7abK41JStQIAqVgRVzpbzo5smXKp4YfRvH+8abtTE1Pi6jizo" crossorigin="anonymous"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.7/umd/popper.min.js" integrity="sha384-UO2eT0CpHqdSJQ6hJty5KVphtPhzWj9WO1clHTMGa3JDZwrnQq4sF86dIHNDz0W1" crossorigin="anonymous"></script>
<script src="/static/js/bootstrap.min.js"></script>
```

# Navigation Bar

* [Navbar · Bootstrap](https://getbootstrap.com/docs/4.3/components/navbar/#)

# Jumbotron

* [Jumbotron · Bootstrap](https://getbootstrap.com/docs/4.3/components/jumbotron/)

# Theme

* [GitHub - ColorlibHQ/gentelella: Free Bootstrap 4 Admin Dashboard Template](https://github.com/ColorlibHQ/gentelella)

* [bootstrap · GitHub Topics · GitHub](https://github.com/topics/bootstrap?o=desc&s=stars)


