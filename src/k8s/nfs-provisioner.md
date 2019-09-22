# 

# Links

* [external-storage/nfs at master · kubernetes-incubator/external-storage · GitHub](https://github.com/kubernetes-incubator/external-storage/tree/master/nfs)

# This is a in-cluster nfs server

Choose some volume for your nfs-provisioner instance to store its state & data
in and mount the volume at `/export` in `deploy/kubernetes/deployment.yaml`. It
doesn't have to be a `hostPath` volume, it can e.g. be a PVC. Note that the
volume must have a supported file system on it: any local filesystem on Linux
is supported & **NFS is not supported**.



