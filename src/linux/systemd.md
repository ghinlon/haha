# systemd

# Links

* [Systemd Essentials: Working with Services, Units, and the Journal | DigitalOcean](https://www.digitalocean.com/community/tutorials/systemd-essentials-working-with-services-units-and-the-journal)
* [Understanding Systemd Units and Unit Files | DigitalOcean](https://www.digitalocean.com/community/tutorials/understanding-systemd-units-and-unit-files)
* [Chapter 10. Managing Services with systemd - Red Hat Customer Portal](https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/7/html/system_administrators_guide/chap-managing_services_with_systemd)
* [Using systemd as a better cron – Horrible Hacks – Medium](https://medium.com/horrible-hacks/using-systemd-as-a-better-cron-a4023eea996d)

# Conflicts between dnsmasq and systemd-resolved

I agree that disabling systemd-resolved is a bad idea and should be avoided.

[Ubuntu 17.04 - Conflicts between systemd-resolved, dnsmasq and resolvconf](https://cnly.github.io/2017/09/02/ubuntu1704-dnsmasq-resolvconf.html)

Append to Config File `/etc/default/dnsmasq`:

```
DNSMASQ_EXCEPT=lo
```

# change default editor to vim for "sudo systemctl edit <unit>"

```
export SYSTEMD_EDITOR="vim" 
```

And then sudo visudo and add this line:

```
Defaults  env_keep += "SYSTEMD_EDITOR"
```

# Basic

```
systemctl start|stop|restart|reload|enable|disable|status nginx.service
systemctl is-active|is-enabled|is-failed application.service	    // to check
systemctl mask|unmask nginx.service	    // mark a unit as completely unstartable, automatically or manually, by linking it to /dev/null. 
systemctl   // default to systemctl list-units
systemctl list-units --all
systemctl list-unit-files
systemctl cat nginx.service
systemctl list-dependencies nginx.service
systemctl list-dependencies --all nginx.service	    // To expand all dependent units recursively
systemctl show nginx.service	    // to see the low-level details of the unit's settings
systemctl edit nginx.service
systemctl edit --full nginx.service
systemctl daemon-reload	    // After modifying a unit file, you should reload the systemd process itself to pick up your changes

systemctl rescue		// single-user environment 
systemctl --no-wall rescue
systemctl emergency		// the most minimal environment
systemctl --no-wall emergency

systemctl halt|poweroff|reboot|suspend|hibernate|hybrid-sleep

systemctl list-timers

journalctl
journalctl -b	// see the journal entries from the current boot
journalctl -k // To see only kernel messages, such as those that are typically represented by dmesg
journalctl -k -b
journalctl -u nginx.service
journalctl -b -u nginx.service
journalctl --list-boots
journalctl --since "2015-01-10" --until "2015-01-11 03:00"
journalctl --since yesterday
journalctl --since 09:00 --until "1 hour ago"
journalctl _PID=8088
journalctl _UID=33 --since today
journalctl -F _GID
journalctl /usr/bin/bash
journalctl -p err -b
    0: emerg
    1: alert
    2: crit
    3: err
    4: warning
    5: notice
    6: info
    7: debug

journalctl -n 20	// as tail -n 20
journalctl -f		// as tail -f
journalctl --disk-usage

```

# Targets (Runlevels)

```
systemctl list-unit-files --type=target
systemctl get-default	    // To view the default target 
systemctl set-default multi-user.target
systemctl list-dependencies multi-user.target	    // To see what units are tied to a target
```

# Time

```
timedatectl list-timezones
timedatectl set-timezone zone
timedatectl status
```

# Unit Files

```
/lib/systemd/system/
/usr/lib/systemd/system/
/etc/systemd/system/	    // this directory location take precedence over any of the other locations
/run/systemd/system/	    // run-time
```

The correct way to do this is to create a directory named after the unit file with `.d` appended on the end. So for a unit called `example.service`, a subdirectory called `example.service.d` could be created. Within this directory a file ending with `.conf` can be used to override or extend the attributes of the system's unit file.


E.g.

`app.service`

```systemd
[Unit]
Description= describe the name and basic functionality of the unit
Documentation=
Requires= any units upon which this unit essentially depends. These units are started in parallel with the current unit by default.
Wants= similar to Requires=, but less strict.
BindsTo= similar to Requires=,  but also causes the current unit to stop when the associated unit terminates.
Before= I'm before these units. they are after me.
After= I'm after these units. they are before me.
Conflicts= 
Condition...=
Assert...=


[Install] 
WantedBy= For instance, if the current unit has WantedBy=multi-user.target, a directory called multi-user.target.wants will be created within /etc/systemd/system (if not already available) and a symbolic link to the current unit will be placed within. Disabling this unit removes the link and removes the dependency relationship.

RequiredBy= similar to the WantedBy= directive, but instead specifies a required dependency that will cause the activation to fail if not met. When enabled, a unit with this directive will create a directory ending with .requires.

Alias= another name
Also= Supporting units that should always be available when this unit is active can be listed here. 
DefaultInstance= 

[Service]
Type= This is important because it tells systemd how to correctly manage the servie and find out its state. 

RemainAfterExit= commonly used with the oneshot type. It indicates that the service should be considered active even after the process exits.

PIDFile= if Type = forking

BusName= if Type = dbus

NotifyAccess= if Type = notify, this can be "none", "main", "all", default "none"

ExecStart= if Type != "oneshot", this may only be specified once 
ExecStartPre=
ExecStartPost=
ExecReload=
ExecStop= If this is not given, the process will be killed immediately when the service is stopped.
ExecStopPost=
RestartSec= the amount of time to wait before attempting to restart the service.
Restart=  This can be set to values like "always", "on-success", "on-failure", "on-abnormal", "on-abort", or "on-watchdog". 
TimeoutSec= This configures the amount of time that systemd will wait when stopping or stopping the service before marking it as failed or forcefully killing it. 

[Socket]
ListenStream= TCP
ListenDatagram= UDP
ListenSequentialPacket= Unix
ListenFIFO= you can also specify a FIFO buffer instead of a socket.
```

# Template Unit Files

`example@.service` be used to create `example@instance1.service`

```
%n		// full resulting unit name
%N		// reverse escaping

%p		// unit name prefix
%P

%i		// instance name
%I		

%f		// unescaped instance name or the prefix name, prepended with a /
%c		// indicate the control group

%u		// user name
%U		// UID
%H		// host name
%%		// literal %%
```


E.g.

`/usr/lib/systemd/system/postifix.service`

```systemd
[Unit]
Description=Postfix Mail Transport Agent
After=syslog.target network.target
Conflicts=sendmail.service exim.service

[Service]
Type=forking
PIDFile=/var/spool/postfix/pid/master.pid
EnvironmentFile=-/etc/sysconfig/network
ExecStartPre=-/usr/libexec/postfix/aliasesdb
ExecStartPre=-/usr/libexec/postfix/chroot-update
ExecStart=/usr/sbin/postfix start
ExecReload=/usr/sbin/postfix reload
ExecStop=/usr/sbin/postfix stop

[Install]
WantedBy=multi-user.target
```

# Creating Custom Unit Files

1. Prepare the executable file with the custom service. 
2. Create a unit file in the /etc/systemd/system/ directory and make sure it has correct file permissions.

	```sh
	touch /etc/systemd/system/name.service 
	chmod 664 /etc/systemd/system/name.service
	```
3. cfg

	```sh
	[Unit]
	Description=service_description
	After=network.target

	[Service]
	ExecStart=path_to_executable
	Type=forking
	PIDFile=path_to_pidfile

	[Install]
	WantedBy=default.target
	```

E.g.

```sh
cat /etc/issue
Ubuntu 18.04.1 LTS \n \l

// 1

file /lib/systemd/system/shadowsocks-libev*
/lib/systemd/system/shadowsocks-libev-local@.service:  ASCII text
/lib/systemd/system/shadowsocks-libev-redir@.service:  ASCII text
/lib/systemd/system/shadowsocks-libev-server@.service: ASCII text
/lib/systemd/system/shadowsocks-libev.service:         ASCII text
/lib/systemd/system/shadowsocks-libev-tunnel@.service: ASCII text

$ cat /lib/systemd/system/shadowsocks-libev-server@.service 
#  This file is part of shadowsocks-libev.
#
#  Shadowsocks-libev is free software; you can redistribute it and/or modify
#  it under the terms of the GNU General Public License as published by
#  the Free Software Foundation; either version 3 of the License, or
#  (at your option) any later version.
#
#  This is a template unit file. Users may copy and rename the file into
#  config directories to make new service instances. See systemd.unit(5)
#  for details.

[Unit]
Description=Shadowsocks-Libev Custom Server Service for %I
Documentation=man:ss-server(1)
After=network.target

[Service]
Type=simple
CapabilityBoundingSet=CAP_NET_BIND_SERVICE
ExecStart=/usr/bin/ss-server -c /etc/shadowsocks-libev/%i.json

[Install]
WantedBy=multi-user.target

$ cat /lib/systemd/system/shadowsocks-libev-local@.service 
#  This file is part of shadowsocks-libev.
#
#  Shadowsocks-libev is free software; you can redistribute it and/or modify
#  it under the terms of the GNU General Public License as published by
#  the Free Software Foundation; either version 3 of the License, or
#  (at your option) any later version.
#
#  This is a template unit file. Users may copy and rename the file into
#  config directories to make new service instances. See systemd.unit(5)
#  for details.

[Unit]
Description=Shadowsocks-Libev Custom Client Service for %I
Documentation=man:ss-local(1)
After=network.target

[Service]
Type=simple
CapabilityBoundingSet=CAP_NET_BIND_SERVICE
ExecStart=/usr/bin/ss-local -c /etc/shadowsocks-libev/%i.json

[Install]
WantedBy=multi-user.target

// 2
file /etc/shadowsocks-libev/config.json 
/etc/shadowsocks-libev/config.json: ASCII text

cd /lib/systemd/system/
mv shadowsocks-libev-local\@.service shadowsocks-libev-local\@config.service 
systemctl enable shadowsocks-libev-local\@config.service 
Created symlink /etc/systemd/system/multi-user.target.wants/shadowsocks-libev-local@config.service → /lib/systemd/system/shadowsocks-libev-local@config.service.
systemctl enable shadowsocks-libev-local\@config.service 
journalctl -u shadowsocks-libev-local\@config -f
```

