# [nginx](https://nginx.org/en/docs/)

# Links

# Installation

Just install from prebuild package: 

[nginx: Linux packages](https://nginx.org/en/linux_packages.html#RHEL-CentOS)

# Beginner’s Guide

* [Beginner’s Guide](https://nginx.org/en/docs/beginners_guide.html)

```
nginx -s <stop|quit|reload|reopen>

    stop — fast shutdown
    quit — graceful shutdown
    reload — reloading the configuration file
    reopen — reopening the log files
```

# OpenResty Lua Support

* [OpenResty® - Official Site](https://openresty.org/en/)
* [GitHub - openresty/lua-nginx-module: Embed the Power of Lua into NGINX HTTP servers](https://github.com/openresty/lua-nginx-module#installation)

# Getting Started

* [OpenResty - Getting Started](https://openresty.org/en/getting-started.html)

Config PATH

```
# put them in ~/.bashrc or ~/.bash_profile
PATH=/usr/local/openresty/nginx/sbin:$PATH
export PATH
```

# Test

```
mkdir ~/work
cd ~/work
mkdir logs/ conf/
```

Create a simple plain text file named `conf/nginx.conf` with the following
contents in it:

```
worker_processes  1;
error_log logs/error.log;
events {
    worker_connections 1024;
}
http {
    server {
        listen 8080;
        location / {
            default_type text/html;
            content_by_lua_block {
                ngx.say("<p>hello, world</p>")
            }
        }
    }
}
```

Run:

```
nginx -p `pwd`/ -c conf/nginx.conf
```

We can use curl to access our new web service that says HelloWorld:

`curl http://localhost:8080/`

If everything is okay, we should get the output

`<p>hello, world</p>`

