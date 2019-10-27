# [client_golang - GoDoc](https://godoc.org/github.com/prometheus/client_golang)

# Links

* [GitHub - prometheus/client_golang: Prometheus instrumentation library for Go applications](https://github.com/prometheus/client_golang)
* [Metric and label naming | Prometheus](https://prometheus.io/docs/practices/naming/) 
* [Instrumenting a Go application | Prometheus](https://prometheus.io/docs/guides/go-application/)
* [Metric types | Prometheus](https://prometheus.io/docs/concepts/metric_types/)

# Install

```
go get github.com/prometheus/client_golang/...
```
# Metrics and Collectors

In addition to the fundamental metric types Gauge, Counter, Summary, Histogram,
and Untyped, a very important part of the Prometheus data model is the
partitioning of samples along dimensions called labels, which results in metric
vectors. The fundamental types are GaugeVec, CounterVec, SummaryVec,
HistogramVec, and UntypedVec. 

While only the fundamental metric types implement the Metric interface, both
the metrics and their vector versions implement the Collector interface. 

**A Collector manages the collection of a number of Metrics**, but for
convenience, a Metric can also “collect itself”. 

Note that Gauge, Counter, Summary, Histogram, and Untyped are interfaces
themselves while GaugeVec, CounterVec, SummaryVec, HistogramVec, and UntypedVec
are not. (**What does interface themselves mean ?**)

To create instances of Metrics and their vector versions, you need a suitable …
Opts struct, i.e. GaugeOpts, CounterOpts, SummaryOpts, HistogramOpts, or
UntypedOpts. 

If you just need to call a function to get a single float value to collect as
a metric, GaugeFunc, CounterFunc, or UntypedFunc might be interesting
shortcuts. 


```go
type Collector interface {
    Describe(chan<- *Desc)	
	Collect(chan<- Metric)
}


type Metric interface {
	Desc() *Desc
	Write(*dto.Metric) error
}
```

So what does a `Metric` interface do ?

it returns a `*Desc` and Write a `*dto.Metric`.

# Collector

A Collector has to be registered for collection. 

# Registry

# func DescribeByCollect

```go
func DescribeByCollect(c Collector, descs chan<- *Desc) {
	metrics := make(chan Metric)
	go func() {
		c.Collect(metrics)
		close(metrics)
	}()
	for m := range metrics {
		descs <- m.Desc()
	}
}
```





# Metric names

e.g.

```
prometheus_notifications_total
process_cpu_seconds_total
http_request_duration_seconds
```

Rules:

`app_what_unit_[total]`


# What to instrument

* [Instrumenting a Go service for Prometheus | There is no magic here](https://alex.dzyoba.com/blog/go-prometheus-service/)
* [The RED Method: key metrics for microservices architecture](https://www.weave.works/blog/the-red-method-key-metrics-for-microservices-architecture/)

four golden signals:

* Traffic or Request Rate
* Errors
* Latency or Duration of the requests
* Saturation

