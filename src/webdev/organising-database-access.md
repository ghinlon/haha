# Organising Database Access

# Links

* [Practical Persistence in Go: Organising Database Access - Alex Edwards](https://www.alexedwards.net/blog/organising-database-access)

# Overview

In general, it is good practice to inject dependencies into your handlers. 

# Closures for Dependency Injection

Since method form of handlers can't spread across multiple packages. so that's
a bad practice. In that case, an alternative approach is to create a config
package exporting an Application struct and have your handler functions close
over this to form a closure. Very roughly:

```go
func main() {
    app := &config.Application{
        ErrorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
    }

    mux.Handle("/", handlers.Home(app))
}

func Home(app *config.Application) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        ...
        ts, err := template.ParseFiles(files...)
        if err != nil {
            app.ErrorLog.Println(err.Error())
            http.Error(w, "Internal Server Error", 500)
            return
        }
        ...
    }
}
```

also see: [_tree · GitHub](https://gist.github.com/alexedwards/5cd712192b4831058b21)



