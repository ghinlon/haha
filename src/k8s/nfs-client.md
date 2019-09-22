# Kubernetes NFS-Client Provisioner

# Links

* [external-storage/nfs-client at master · kubernetes-incubator/external-storage · GitHub](https://github.com/kubernetes-incubator/external-storage/tree/master/nfs-client)

# What is it

nfs-client is an automatic provisioner that use your existing and already
configured NFS server to support dynamic provisioning of Kubernetes Persistent
Volumes via Persistent Volume Claims. Persistent volumes are provisioned as
`${namespace}-${pvcName}-${pvName}`.

**To note again, you must already have an NFS Server.**


# Install

* [charts/stable/nfs-client-provisioner at master · helm/charts · GitHub](https://github.com/helm/charts/tree/master/stable/nfs-client-provisioner)


```
helm install --name my-release --set nfs.server=x.x.x.x --set nfs.path=/exported/path stable/nfs-client-provisioner
```




