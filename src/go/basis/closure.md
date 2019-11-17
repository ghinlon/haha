# Closure

# Links

* 

# Overview

Regardless of what you do with a closure it will always be able to access the
variables that are local to the scope it was created in — which in this case
means that fn will always have access to the `next` variable.

```go
func myMiddleware(next http.Handler) http.Handler {
    fn := func(w http.ResponseWriter, r *http.Request) {
        // TODO: Execute our middleware logic here...
        next.ServeHTTP(w, r)
    }

    return http.HandlerFunc(fn)
}
```

