# [sql_exporter - GoDoc](https://godoc.org/github.com/free/sql_exporter#Collector)

# Links

* [GitHub - free/sql_exporter: Database agnostic SQL exporter for Prometheus](https://github.com/free/sql_exporter)

# Overview

After several days, still didn't get how this software works. For me, the most
dificult is that I can't figer out how those struct work together.  

Cant figer out, a column of a row of sql results, is be as a label, or as a metric ?  

Sun Jul 14 23:41:58 CST 2019

After read [sql_exporter/sql_exporter.yml at master · free/sql_exporter · GitHub](https://github.com/free/sql_exporter/blob/master/documentation/sql_exporter.yml), I finally get something:

```
    # The result columns conceptually fall into two categories:
    #  * zero or more key columns: their values will be directly mapped to labels of the same name;
    #  * one or more value columns:
    #     * if exactly one value column, the column name name is ignored and its value becomes the metric value
    #     * with multiple value columns, a `value_label` must be defined; the column name will populate this label and
    #       the column value will popilate the metric value.
```

但是这什么意思啊？ 什么叫0或多个key列？？？他们的值会直接给有相同labels的，
谁和谁相同啊？？？  

label就是label, 分什么key_label, value_label???

Mon Jul 15 00:30:09 CST 2019


This query may result in multiple rows. There has metric with a label named
`db`, then the result of this will use the value of the `db` column to the
label `db`,kind of: 

```
mssql_log_growths{db="db1"} 12
mssql_log_growths{db="db2"} 25
mssql_log_growths{db="db3"} 32
```

```
SELECT rtrim(instance_name) AS db, cntr_value AS counter
FROM sys.dm_os_performance_counters
WHERE counter_name = 'Log Growths' AND instance_name <> '_Total'
```


# Is one query only get one metric ?

This is the core work.  

```go
func (q *Query) Collect(ctx context.Context, conn *sql.DB, ch chan<- Metric) {

	...

	for rows.Next() {
		row, err := q.scanRow(rows, dest)
		if err != nil {
			ch <- NewInvalidMetric(err)
			continue
		}
		for _, mf := range q.metricFamilies {
			mf.Collect(row, ch)
		}
	}

	...
}
```



