# Querying Prometheus

* [Querying basics | Prometheus](https://prometheus.io/docs/prometheus/latest/querying/basics/)
* [Operators | Prometheus](https://prometheus.io/docs/prometheus/latest/querying/operators/)
* [Query functions | Prometheus](https://prometheus.io/docs/prometheus/latest/querying/functions/)
* [Combining alert conditions – Robust Perception | Prometheus Monitoring Experts](https://www.robustperception.io/combining-alert-conditions)


In Prometheus's expression language, an expression or sub-expression can evaluate to one of four types:

* **Instant vector** - a set of time series containing a single sample for each time series, all sharing the same timestamp
* **Range vector** - a set of time series containing a range of data points over time for each time series
* **Scalar** - a simple numeric floating point value
* **String** - a simple string value; currently unused

# Time series Selectors

## Instant vector selectors

Instant vector selectors allow the selection of a set of time series and a single sample value for each at a given timestamp (instant): in the simplest form, only a metric name is specified. This results in an instant vector containing elements for all time series that have this metric name.

## Range Vector Selectors

Range vector literals work like instant vector literals, except that they select a range of samples back from the current instant. Syntactically, a range duration is appended in square brackets ([]) at the end of a vector selector to specify how far back in time values should be fetched for each resulting range vector element.


# Offset modifier

The offset modifier allows changing the time offset for individual instant and range vectors in a query.


# Vector matching

Operations between vectors attempt to find a matching element in the right-hand side vector for each entry in the left-hand side. There are two basic types of matching behavior: One-to-one and many-to-one/one-to-many.

## One-to-one vector matches

The `ignoring` keyword allows ignoring certain labels when matching, while the `on` keyword allows reducing the set of considered labels to a provided list:

```
<vector expr> <bin-op> ignoring(<label list>) <vector expr>
<vector expr> <bin-op> on(<label list>) <vector expr>
```

E.g.

Input:

```
method_code:http_errors:rate5m{method="get", code="500"}  24
method_code:http_errors:rate5m{method="get", code="404"}  30
method_code:http_errors:rate5m{method="put", code="501"}  3
method_code:http_errors:rate5m{method="post", code="500"} 6
method_code:http_errors:rate5m{method="post", code="404"} 21

method:http_requests:rate5m{method="get"}  600
method:http_requests:rate5m{method="del"}  34
method:http_requests:rate5m{method="post"} 120
```

Query:

```
method_code:http_errors:rate5m{code="500"} / ignoring(code) method:http_requests:rate5m

{method="get"}  0.04            //  24 / 600
{method="post"} 0.05            //   6 / 120
```

因为要标签完全一样也有法比，所以要忽略code这个标签


## Many-to-one and one-to-many vector matches


```
<vector expr> <bin-op> ignoring(<label list>) group_left(<label list>) <vector expr>
<vector expr> <bin-op> ignoring(<label list>) group_right(<label list>) <vector expr>
<vector expr> <bin-op> on(<label list>) group_left(<label list>) <vector expr>
<vector expr> <bin-op> on(<label list>) group_right(<label list>) <vector expr>
```

```
method_code:http_errors:rate5m / ignoring(code) group_left method:http_requests:rate5m
```

In this case the left vector contains more than one entry per method label value. Thus, we indicate this using group_left. The elements from the right side are now matched with multiple elements with the same method label on the left:

```
{method="get", code="500"}  0.04            //  24 / 600
{method="get", code="404"}  0.05            //  30 / 600
{method="post", code="500"} 0.05            //   6 / 120
{method="post", code="404"} 0.175           //  21 / 120
```

Many-to-one and one-to-many matching are advanced use cases that should be carefully considered. Often a proper use of ignoring(<labels>) provides the desired outcome.

# Aggregation operators

`without` removes the listed labels from the result vector, while all other labels are preserved the output.
`by` does the opposite and drops labels that are not listed in the by clause, even if their label values are identical between all elements of the vector.


# 我只关心早上9点到晚上10点

```
job:request_latency_seconds:mean5m{job="myjob"} > 0.5 and ON() hour() > 9 < 10
```

Here hour() returns the hour of the day in UTC as a time series with no labels. We can use comparison operators to filter this down to the time we care about. Finally as there's no matching labels on either side of the and, we specify ON() to ignore all labels when matching.




