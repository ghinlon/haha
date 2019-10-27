# node_exporter

# Links

* [GitHub - prometheus/node_exporter: Exporter for machine metrics](https://github.com/prometheus/node_exporter)

# Textfile Collector

* [Exposition formats | Prometheus](https://prometheus.io/docs/instrumenting/exposition_formats/)

To use it, set the `--collector.textfile.directory` flag on the Node exporter.
The collector will parse all files in that directory matching the glob `*.prom`
using the text format.





