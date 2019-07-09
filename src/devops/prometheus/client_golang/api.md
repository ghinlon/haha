# [api - GoDoc](https://godoc.org/github.com/prometheus/client_golang/api)

# Links

* [documentation for API client · Issue #194 · prometheus/client_golang · GitHub](https://github.com/prometheus/client_golang/issues/194#issuecomment-194319793)

# Basics


```go
	rawurl := "http://prom.org"

	c, err := api.NewClient(
		api.Config{
			Address: rawurl,
		})

	if err != nil {
		log.Fatal(err)
	}

	api := v1.NewAPI(c)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	v, w, err := api.Query(ctx, `100 - (avg by (instance) (irate(node_cpu_seconds_total{job="node",mode="idle"}[5m])) * 100)`, time.Now())

	if err != nil || w != nil {
		log.Fatal(err)
	}

	fmt.Println(v.String())

```

