# Change the default StorageClass

# Links

* [Change the default StorageClass - Kubernetes](https://kubernetes.io/docs/tasks/administer-cluster/change-default-storage-class/)


# How

The default StorageClass has an annotation
`storageclass.kubernetes.io/is-default-class` set to `true`. Any other value or
absence of the annotation is interpreted as `false`.

```
kubectl patch storageclass <your-class-name> -p '{"metadata": {"annotations":{"storageclass.kubernetes.io/is-default-class":"true"}}}'
// Verify
kubectl get storageclass
```

Please note that at most one StorageClass can be marked as default. If two or
more of them are marked as default, a PersistentVolumeClaim without
storageClassName explicitly specified cannot be created.



