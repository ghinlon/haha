# ip rule

# Links

* [A Quick Introduction to Linux Policy Routing - Scott's Weblog - The weblog of an IT pro focusing on cloud computing, Kubernetes, Linux, containers, and networking](https://blog.scottlowe.org/2013/05/29/a-quick-introduction-to-linux-policy-routing/)
* [Linux Advanced Routing & Traffic Control HOWTO](https://lartc.org/howto/index.html)

# Notes

* table `main` is really the default table, and table `default` is the last
  table, which after talbe `main`, generally never used.

	```
	0:      from all lookup local
	32766:  from all lookup main
	32767:  from all lookup default
	```

# Add a new rule

```
# echo 200 John >> /etc/iproute2/rt_tables
# ip rule add from 10.0.0.10 table John
# ip rule ls
0: from all lookup local
32765: from 10.0.0.10 lookup John
32766: from all lookup main
32767: from all lookup default

// generate John’s table
# ip route add default via 195.96.98.253 dev ppp2 table John
# ip route flush cache
```

Every table need a default gateway.

**routing table is just a routing table, route is decided by rule.**


