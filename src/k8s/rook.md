# Rook

# Links

* [Rook Docs: quickstart](https://rook.io/docs/rook/v1.1/quickstart-toc.html)

# Install

* [Rook Docs: k8s-pre-reqs](https://rook.io/docs/rook/v1.1/k8s-pre-reqs.html)


Rook Ceph requires a Linux kernel built with the RBD module. 

```
modprobe rbd
```

LVM package

```
# Centos
sudo yum install -y lvm2

# Ubuntu
sudo apt-get install -y lvm2
```

## Deploy the Rook Operator

It is recommended that the rook operator be installed into the rook-ceph
namespace (you will install your clusters into separate namespaces).

* [helm-operator](https://rook.io/docs/rook/v1.1/helm-operator.html)

 ```
helm repo add rook-release https://charts.rook.io/release
helm search rook-ceph
helm install --namespace rook-ceph rook-release/rook-ceph

// or
helm fetch rook-release/rook-ceph
// custom value
helm install --namespace rook-ceph -f value.custom.yaml rook-release/rook-ceph 

// output
NOTES:
The Rook Operator has been installed. Check its status by running:
  kubectl --namespace rook-ceph get pods -l "app=rook-ceph-operator"

Visit https://rook.io/docs/rook/master for instructions on how to create and configure Rook clusters

Note: You cannot just create a CephCluster resource, you need to also create a namespace and
install suitable RBAC roles and role bindings for the cluster. The Rook Operator will not do
this for you. Sample CephCluster manifest templates that include RBAC resources are available:

- https://rook.github.io/docs/rook/master/ceph-quickstart.html
- https://github.com/rook/rook/blob/master/cluster/examples/kubernetes/ceph/cluster.yaml

Important Notes:
- The links above are for the unreleased master version, if you deploy a different release you must find matching manifests.
- You must customise the 'CephCluster' resource at the bottom of the sample manifests to met your situation.
- Each CephCluster must be deployed to its own namespace, the samples use `rook-ceph` for the cluster.
- The sample manifests assume you also installed the rook-ceph operator in the `rook-ceph` namespace.
- The helm chart includes all the RBAC required to create a CephCluster CRD in the same namespace.
- Any disk devices you add to the cluster in the 'CephCluster' must be empty (no filesystem and no partitions).
- In the 'CephCluster' you must refer to disk devices by their '/dev/something' name, e.g. 'sdb' or 'xvde'.
 ```
 
## Create a Rook Ceph Cluster

* [ceph-cluster-crd](https://rook.io/docs/rook/v1.1/ceph-cluster-crd.html)



# Production Environments

* [ceph-examples](https://rook.io/docs/rook/v1.1/ceph-examples.html)
* [Ceph Common Issues](https://rook.io/docs/rook/v1.1/ceph-common-issues.html)


# Rook Toolbox

* [Rook Toolbox](https://rook.io/docs/rook/v1.1/ceph-toolbox.html)

```
kubectl -n rook-ceph exec -it $(kubectl -n rook-ceph get pod -l "app=rook-ceph-tools" -o jsonpath='{.items[0].metadata.name}') bash

ceph status
ceph osd status
ceph df
rados df
```







