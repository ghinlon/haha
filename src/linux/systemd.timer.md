# systemd.timer

# Links

* [Linux Task Scheduler | Scheduling Tasks with systemd Timers](https://coreos.com/os/docs/latest/scheduling-tasks-with-systemd-timers.html)
* [RHEL7: How to use Systemd timers. - CertDepot](https://www.certdepot.net/rhel7-use-systemd-timers/)

# List

```
systemctl list-timers
systemctl list-timers --all
```

# How to Use

date.service:

```
[Unit]
Description=Prints date into /tmp/date file

[Service]
Type=oneshot
ExecStart=/usr/bin/sh -c '/usr/bin/date >> /tmp/date'

[Install]
WantedBy=date.timer
```

date.timmer:

```
[Unit]
Description=Run date.service every 10 minutes

[Timer]
OnCalendar=*:0/10
Unit=promreporter.service

[Install]
WantedBy=timers.target
```

This config will run `date.service` every 10 minutes. 

