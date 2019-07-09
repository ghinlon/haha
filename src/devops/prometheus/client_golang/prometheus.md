# [client_golang prometheus](https://godoc.org/github.com/prometheus/client_golang/prometheus)

# Links

* [Instrumenting a Go application | Prometheus](https://prometheus.io/docs/guides/go-application/)
* [Metric types | Prometheus](https://prometheus.io/docs/concepts/metric_types/)

# What is Collector ?

```go
type Collector interface {
    Describe(chan<- *Desc)	
	Collect(chan<- Metric)
}
```

Collector is the interface implemented by anything that can be used by
Prometheus to collect metrics. A Collector has to be registered for collection. 

The stock metrics provided by this package (Gauge, Counter, Summary, Histogram,
Untyped) are also Collectors (which only ever collect one metric, namely
itself). 


# What is Metric ?

I think, That's a Collector is a kind of Reader, and a Metric is a kind of
Writer, an Encoder.

What the different is, Reader is used to Read into a `[]byte`, while Collector
is used to Read into a `chan<- promethheus.Metric`, and a Collector is Read
From a Metric, Cause MetricVec is not implemented `Collector Interface`, so how it
can be collected.

```go
type Metric interface {
	Desc() *Desc
	Write(*dto.Metric) error
}

type invalidMetric struct {
	desc *Desc
	err  error
}
func (m *invalidMetric) Desc() *Desc { return m.desc }
func (m *invalidMetric) Write(*dto.Metric) error { return m.err }


```

