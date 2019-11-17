# Alertmanager

# Links

* [prometheus/alertmanager: Prometheus Alertmanager](https://github.com/prometheus/alertmanager)
* [Alerting overview | Prometheus](https://prometheus.io/docs/alerting/overview/)
* [Integrations | Prometheus](https://prometheus.io/docs/operating/integrations/#alertmanager-webhook-receiver)
* [Alerting | Prometheus](https://prometheus.io/docs/practices/alerting/)
* [Configuration | Prometheus](https://prometheus.io/docs/alerting/configuration/)

# Packages

* [template - GoDoc](https://godoc.org/github.com/prometheus/alertmanager/template)

# SMS

* [GitHub - messagebird/sachet: SMS alerts for Prometheus' Alertmanager](https://github.com/messagebird/sachet)

# <webhook_config>

```
# Whether or not to notify about resolved alerts.
[ send_resolved: <boolean> | default = true ]

# The endpoint to send HTTP POST requests to.
url: <string>

# The HTTP client's configuration.
[ http_config: <http_config> | default = global.http_config ]
```

The Alertmanager will send HTTP POST requests in the following JSON format to
the configured endpoint:

```
{
  "version": "4",
  "groupKey": <string>,    // key identifying the group of alerts (e.g. to deduplicate)
  "status": "<resolved|firing>",
  "receiver": <string>,
  "groupLabels": <object>,
  "commonLabels": <object>,
  "commonAnnotations": <object>,
  "externalURL": <string>,  // backlink to the Alertmanager.
  "alerts": [
    {
      "status": "<resolved|firing>",
      "labels": <object>,
      "annotations": <object>,
      "startsAt": "<rfc3339>",
      "endsAt": "<rfc3339>",
      "generatorURL": <string> // identifies the entity that caused the alert
    },
    ...
  ]
}
```

