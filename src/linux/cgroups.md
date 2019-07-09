# cgroups

# Links

* [Resource Management Guide - Red Hat Customer Portal](https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/7/html/resource_management_guide/)


# DEFAULT CGROUP HIERARCHIES and RESOURCE CONTROLLERS IN LINUX KERNEL

systemd automatically mounts hierarchies for important kernel resource controllers in the `/sys/fs/cgroup/` directory.

Find the list of currently mounted resource controllers in `/proc/cgroups`

```sh
$ cat /proc/cgroups 
#subsys_name    hierarchy       num_cgroups     enabled
cpuset  3       1       1
cpu     5       70      1
cpuacct 5       70      1
blkio   12      70      1
memory  2       215     1
devices 9       70      1
freezer 6       2       1
net_cls 4       1       1
perf_event      11      1       1
net_prio        4       1       1
hugetlb 8       1       1
pids    10      80      1
rdma    7       1       1
```

# Systemd Unit Types

* Service — A process or a group of processes
* Scope — A group of externally created processes.
* Slice — A group of hierarchically organized units.

	Slices describe resource limits. I have one slice that limits the resources used by all my services as a whole.

# CREATING CONTROL GROUPS

The `systemd-run` command is used to create and start a transient service or scope unit and run a custom command in the unit.

```sh
systemd-run --unit=name --scope --slice=slice_name command
// Use the optional --scope parameter to create a transient scope unit instead of service unit that is created by default.
// With the --slice option, you can make your newly created service or scope unit a member of a specified slice. Replace slice_name with the name of an existing slice (as shown in the output of systemctl -t slice), or create a new slice by passing a unique name. By default, services and scopes are created as members of the system.slice.
```

E.g.

`systemd-run --unit=toptest --slice=test top -b`


or use config file.

# MODIFYING CONTROL GROUPS

```
systemctl set-property name parameter=value
systemctl set-property --runtime name property=value	// makes your settings transient

systemctl set-property httpd.service CPUShares=600 MemoryLimit=500M
```

# Modifying Unit Files

default cfg file: `/etc/systemd/system.conf`
unit cfg file: `/etc/systemd/system/httpd.service.d/cpu.conf`

```sh
[Service]
CPUShares=1500		// default 1024
CPUQuota=20%
MemoryLimit=1G	
BlockIOWeight=100	// Choose a single value between 10 and 1000, the default setting is 1000.
BlockIODeviceWeight=device_name value
BlockIOReadBandwidth=device_name value
BlockIOWriteBandwidth=device_name value
BlockIODeviceWeight=/home/jdoe 750
BlockIOReadBandwith=/var/log 5M
DeviceAllow=device_name options
// device_name stands for a path to a device node or a device group name as specified in /proc/devices. Replace options with a combination of r, w, and m to allow the unit to read, write, or create device nodes.
DevicePolicy=value		// strict, closed, auto
Slice=slice_name		// default is system.slice. 
```

Users can thus monitor the usage of the processor with the `systemd-cgtop` command.

reload:

```sh
systemctl daemon-reload
systemctl restart httpd.service
```

# OBTAINING INFORMATION ABOUT CONTROL GROUPS

```
systemd-cgls	// view the hierarchy of control groups
systemd-cgls name
systemd-cgls memory
systemd-cgtop 	// monitor their resource consumption in real time.
machinectl	// dedicated to monitoring Linux containers

cat proc/PID/cgroup
```

