# Routers


# Links

Recommendation by Order:

go-chi/chi

* [chi - GoDoc](https://godoc.org/github.com/go-chi/chi)
* [go-chi/chi: lightweight, idiomatic and composable router for building Go HTTP services](https://github.com/go-chi/chi)

gorilla/mux

* [mux - GoDoc](https://godoc.org/github.com/gorilla/mux)
* [gorilla/mux: A powerful HTTP router and URL matcher for building Go web servers with 🦍](https://github.com/gorilla/mux)

go-zoo/bone

* [bone - GoDoc](https://godoc.org/github.com/go-zoo/bone)
* [go-zoo/bone: Lightning Fast HTTP Multiplexer](https://github.com/go-zoo/bone)

bmizerany/pat

* [pat - GoDoc](https://godoc.org/github.com/bmizerany/pat)
* [bmizerany/pat](https://github.com/bmizerany/pat)

Alternative Routers:

* [julienschmidt/httprouter](https://github.com/julienschmidt/httprouter)

  it can’t cope with [conflicting
  patterns](https://github.com/julienschmidt/httprouter/issues/175) caused by
  wildcard parameters, which can be a problem for applications that are using
  a RESTful routing scheme (in our case `/snippet/create` and `/snippet/:id`
  are conflicting patterns).

# Overview


