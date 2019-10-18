# Calico: Secure networking for the cloud native era

# Links

* [Installing Calico for policy and networking (recommended)](https://docs.projectcalico.org/v3.9/getting-started/kubernetes/installation/calico)

# Install

**Note**: 

`pod-network-cidr` can't assign save network as node.

```
// 1
curl https://docs.projectcalico.org/v3.9/manifests/calico.yaml -O

// 2
POD_CIDR="<your-pod-cidr>" \
sed -i -e "s?192.168.0.0/16?$POD_CIDR?g" calico.yaml

// 3
kubectl apply -f calico.yaml
```

