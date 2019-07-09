# redis: a data structures server

# Links

* [Redis Quick Start – Redis](https://redis.io/topics/quickstart)
* [An introduction to Redis data types and abstractions – Redis](https://redis.io/topics/data-types-intro)

# Port

`port 6379`

# PING

`redis-cli ping`

# Example

```
$ redis-cli                                                                
redis 127.0.0.1:6379> ping
PONG
redis 127.0.0.1:6379> set mykey somevalue
OK
redis 127.0.0.1:6379> get mykey
"somevalue"
```




