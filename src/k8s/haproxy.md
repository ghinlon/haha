# HAProxy

# Links

* [HAProxy - The Reliable, High Performance TCP/HTTP Load Balancer](https://www.haproxy.org/)
* [HAProxy version 2.0.5 - Configuration Manual](https://cbonte.github.io/haproxy-dconv/2.0/configuration.html#2)

# Configuration file format


```
HAProxy's configuration process involves 3 major sources of parameters :

  - the arguments from the command-line, which always take precedence
  - the "global" section, which sets process-wide parameters
  - the proxies sections which can take form of "defaults", "listen",
    "frontend" and "backend".
```

# Simple


```
    # Simple configuration for an HTTP proxy listening on port 80 on all
    # interfaces and forwarding requests to a single backend "servers" with a
    # single server "server1" listening on 127.0.0.1:8000
    global
        daemon
        maxconn 256

    defaults
        mode http
        timeout connect 5000ms
        timeout client 50000ms
        timeout server 50000ms

    frontend http-in
        bind *:80
        default_backend servers

    backend servers
        server server1 127.0.0.1:8000 maxconn 32


    # The same configuration defined with a single listen block. Shorter but
    # less expressive, especially in HTTP mode.
    global
        daemon
        maxconn 256

    defaults
        mode http
        timeout connect 5000ms
        timeout client 50000ms
        timeout server 50000ms

    listen http-in
        bind *:80
        server server1 127.0.0.1:8000 maxconn 32
```


Assuming haproxy is in $PATH, test these configurations in a shell with:

```
    $ sudo haproxy -f configuration.conf -c
```

# Global parameters

* [Global parameters](https://cbonte.github.io/haproxy-dconv/2.0/configuration.html#3)

```
daemon

	Makes the process fork into background. This is the recommended mode of
	operation. It is equivalent to the command line "-D" argument. It can be
	disabled by the command line "-db" argument. This option is ignored in
	systemd mode.


```







