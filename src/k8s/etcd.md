# etcd

# Links

* [etcd/Documentation at master · etcd-io/etcd · GitHub](https://github.com/etcd-io/etcd/tree/master/Documentation)
* [etcd/faq.md at master · etcd-io/etcd · GitHub](https://github.com/etcd-io/etcd/blob/master/Documentation/faq.md)

# Getting Started

[etcd/dl_build.md at master · etcd-io/etcd · GitHub](https://github.com/etcd-io/etcd/blob/master/Documentation/dl_build.md)

# Install

[Releases · etcd-io/etcd · GitHub](https://github.com/etcd-io/etcd/releases/)

```
go get -v go.etcd.io/etcd
go get -v go.etcd.io/etcd/etcdctl

// start
etcd

// test
etcdctl put foo bar
// Should Output: OK
```

# Demo

Very Straightforward

[etcd/demo.md at master · etcd-io/etcd · GitHub](https://github.com/etcd-io/etcd/blob/master/Documentation/demo.md)

# Local multi-member cluster

[etcd/local_cluster.md at master · etcd-io/etcd · GitHub](https://github.com/etcd-io/etcd/blob/master/Documentation/dev-guide/local_cluster.md)

# Clustering

* [etcd/clustering.md at master · etcd-io/etcd · GitHub](https://github.com/etcd-io/etcd/blob/master/Documentation/op-guide/clustering.md)
* [etcd/runtime-configuration.md at master · etcd-io/etcd · GitHub](https://github.com/etcd-io/etcd/blob/master/Documentation/op-guide/runtime-configuration.md)
* [etcd/runtime-reconf-design.md at master · etcd-io/etcd · GitHub](https://github.com/etcd-io/etcd/blob/master/Documentation/op-guide/runtime-reconf-design.md)


# Backing up an etcd cluster

* [Operating etcd clusters for Kubernetes - Kubernetes](https://kubernetes.io/docs/tasks/administer-cluster/configure-upgrade-etcd/#backing-up-an-etcd-cluster)

