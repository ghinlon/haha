# cpu time


# Links

* [linux - How are CPU time and CPU usage the same? - Server Fault](https://serverfault.com/questions/648704/how-are-cpu-time-and-cpu-usage-the-same/648708#648708)
* [Understanding Machine CPU usage – Robust Perception | Prometheus Monitoring Experts](https://www.robustperception.io/understanding-machine-cpu-usage)
* [ps utility in linux (procps), how to check which CPU is used - Stack Overflow](https://stackoverflow.com/questions/5732192/ps-utility-in-linux-procps-how-to-check-which-cpu-is-used)
* [Scheduling In Go : Part I - OS Scheduler](https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part1.html)
* [Resource Management Guide - Red Hat Customer Portal](https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/6/html/resource_management_guide/)
* [stress project page](https://people.seas.harvard.edu/~apw/stress/)
* [让 CPU 告诉你硬盘和网络到底有多慢 | Cizixs Write Here](https://cizixs.com/2017/01/03/how-slow-is-disk-and-network/)
* [bertimus9 / systemstat — Bitbucket](https://bitbucket.org/bertimus9/systemstat)


CPU time is allocated in discrete time slices (ticks). For a certain number of time slices, the cpu is busy, other times it is not (which is represented by the idle process). In the picture below the CPU is busy for 6 of the 10 CPU slices. 6/10 = .60 = 60% of busy time (and there would therefore be 40% idle time).

![cputime](img/cputime.png)


A percentage is defined as "a number or rate that is expressed as a certain number of parts of something divided into 100 parts". So in this case, those parts are discrete slices of time and the something is busy time slices vs idle time slices -- the rate of busy to idle time slices.

Since CPUs operate in GHz (billionths of cycles a second). The operating system slices that time in smaller units called ticks. They are not really 1/10 of a second. The tick rate in windows is 10 million ticks in a second and in Linux it is sysconf(_SC_CLK_TCK) (usually 100 ticks per second).

In something like top, the busy cpu cycles are then further broken down into percentages of things like user time and system time. In top on Linux and perfmon in Windows, you will often get a display that goes over 100%, that is because the total is 100% * the_number_of_cpu_cores.

In an operating system, it is the scheduler's job to allocate these precious slices to processes, so the scheduler is what reports this.

**so rate cputime is the usage per second. **

# Node_exporter

This metric comes from `/proc/stat` and tell us how many seconds each CPU spent doing each type of work:

* user: The time spent in userland
* system: The time spent in the kernel
* iowait: Time spent waiting for I/O
* idle: Time the CPU had nothing to do
* irq&softirq: Time servicing interrupts
* guest: If you are running VMs, the CPU they use
* steal: If you are a VM, time other VMs "stole" from your CPUs

These modes are mutually exclusive. A high iowait means that you are disk or network bound, high user or system means that you are CPU bound.

`sum by (mode, instance) (irate(node_cpu_seconds_total{job="node"}[5m]))`

As these values always sum to one second per second for each cpu, the per-second rates are also the ratios of usage. We can use this to calculate the percentage of CPU used, by subtracting the idle usage from 100%:

`100 - (avg by (instance) (irate(node_cpu_seconds_total{job="node",mode="idle"}[5m])) * 100)`

# Howto eat cpu time

最近有个项目，需要吃cpu, 但是不能100％使用，要有法调整使用的%比。

无法直接loop, loop会100％使用cpu，cpu就是这样设计的。

发现这个代码可以占用一个固定量的cpu

```go
func main() {
        // tick := time.NewTicker(time.Nanosecond) // 300%
        // tick := time.NewTicker(time.Microsecond) // 190%
        tick := time.NewTicker(time.Millisecond) // 9%
        // tick := time.NewTicker(time.Second) // 0%
	defer tick.Stop()
	/*
	无法直接这样放goroutine里使的，这种goroutine无办法释放
	for range tick.C {
        }
	*/
	// 正确的方法：
	for range tick.C {
		select {
			case <- stop:
				return
				default:
		}
	}

}
```

# Restricting process CPU usage using cgroups

see [cgroups](cgroups.md)

# stress

这个是吃cpu最方便的工具，配合cgroups，非常好使

```
cat stress.service 
[Unit]
Description=stress
Wants=stress.timer

[Service]
Type=simple
ExecStart=/usr/bin/stress --cpu 2 --io 1 --vm 1 --vm-bytes 128M --timeout 10s --verbose
RemainAfterExit=yes
Slice=stress.slice

[Install]
WantedBy=multi-user.target
```

```
cat stress.timer
[Unit]
Description=Run stress every 2-3 minutes
Requires=stress.service

[Timer]
Unit=stress.service
OnUnitInactiveSec=2min		// m is different from min
RandomizedDelaySec=1min
AccuracySec=1s

[Install]
WantedBy=timers.target

```

```
cat stress.service.d/cpu.conf
[Service]
CPUQuota=20%
```

or 

```
cat stress.slice
[Unit]
Description=Limited resources Slice
DefaultDependencies=no
Before=slices.target

[Slice]
CPUQuota=20%
MemoryLimit=2.7G
```

# Check CPU Usage

## /proc/stat

content in `/proc/stat/`, if used to calculate percentage, must read twice, E.g. :

`idle1 - idle0 / total1 - total0`

## ps 

```sh
ps -C "process" -L -o pid,lwp,pcpu,cpuid,time

    -C: select the process named "process"
    -L: list the process threads
    -o: specify output info
        pid: process id
        lwp: light weight process (thread)
        pcpu: CPU usage (percent)
        cpuid: CPU id
        time: thread time (from start)
```

## top

```sh
top -d delay -n iterations
top -n 1		// once
```

## iostat

This utility display system’s average CPU utilization **since the last reboot**. 

To list the individual report use -c, -d and -n switch for CPU utilization, device utilization and network file system utilization.

```sh
iostat	    // displays cpu, device, network file system
iostat -c	// cpu
iostat -d	// device	
iostat -n	// network file system
```
## vmstat
## mpstat
## sar

```sh
sar 2 5				// interval count
sar -u [ALL] 2 5 		// -u cpu
sar -P ALL 2 3 			// all cpu
sar -P 0 2 3 			// the first cpu

sar -r 				// memory
sar -d 				// device
sar -n ALL|DEV|NDEV|...		// network
sar -q				// load
```

# Process vs Thread

Your program is just a series of machine instructions that need to be executed one after the other sequentially. To make that happen, the operating system uses the concept of a Thread. It’s the job of the Thread to account for and sequentially execute the set of instructions it’s assigned. Execution continues until there are no more instructions for the Thread to execute. This is why I call a Thread, “a path of execution”.

**Every program you run creates a Process and each Process is given an initial Thread.** Threads have the ability to create more Threads. All these different Threads run independently of each other and scheduling decisions are made at the Thread level, not at the Process level. Threads can run concurrently (each taking a turn on an individual core), or in parallel (each running at the same time on different cores). Threads also maintain their own state to allow for the safe, local, and independent execution of their instructions.

以前我认为进程是一级的，一个进程有法有很多线程。

其实不是。

程序拉进内存就是进程，所有程序启动都会分配一个初始线程。线程是进程执行的路.thread is used for os scheduler to schedule cpu time.
