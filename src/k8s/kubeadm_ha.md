# Creating Highly Available clusters with kubeadm

# Links

* [Creating Highly Available clusters with kubeadm - Kubernetes](https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/high-availability/)
* [Creating Highly Available Clusters with kubeadm - V1-12](https://v1-12.docs.kubernetes.io/docs/setup/independent/high-availability/)
* [Set up a High Availability etcd cluster with kubeadm - Kubernetes](https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/setup-ha-etcd-with-kubeadm/)
* [etcd/clustering.md at master · etcd-io/etcd · GitHub](https://github.com/etcd-io/etcd/blob/master/Documentation/op-guide/clustering.md)

# Options for Highly Available topology

* [Options for Highly Available topology - Kubernetes](https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/ha-topology/)

* With stacked control plane nodes. This approach requires less infrastructure.
  The etcd members and control plane nodes are co-located.
* With an external etcd cluster. This approach requires more infrastructure.
  The control plane nodes and etcd members are separated.

# First Steps

* [HAProxy - The Reliable, High Performance TCP/HTTP Load Balancer](https://www.haproxy.org/)

## Create load balancer for kube-apiserver

1. Create a kube-apiserver load balancer with a name that resolves to DNS.  
   **MUST**: make sure DNS works correctly, at least put it in `/etc/hosts`,
   and write a record has current hostname in the `/etc/hosts`
1. Add control plane nodes to the load balancer and test the connection:  
   `nc -v LOAD_BALANCER_IP PORT`  

haproxy.cfg:

```
listen k8s
    bind *:443
	mode tcp
    server master1 10.1.1.11:6443 check fall 3 rise 2
    server master2 10.1.1.12:6443 check fall 3 rise 2
```

Config /etc/hosts

```
10.1.1.11 k8s.go
```

## kubeadm config image pull

# Steps for the first control plane node

* [Issue when using kubeadm with multiple network interfaces · Issue #33618 · kubernetes/kubernetes · GitHub](https://github.com/kubernetes/kubernetes/issues/33618)

# kubeadm api

* [v1beta2 - GoDoc](https://godoc.org/k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1beta2)

To print the defaults for "init" and "join" actions use the following commands:

```
kubeadm config print init-defaults
kubeadm config print join-defaults
```


Example Config File:

```
apiVersion: kubeadm.k8s.io/v1beta2
kind: InitConfiguration
localAPIEndpoint:
  advertiseAddress: 10.1.1.11
  bindPort: 6443
nodeRegistration:
  name: master1
---
apiVersion: kubeadm.k8s.io/v1beta2
kind: ClusterConfiguration
clusterName: k8s.go
controlPlaneEndpoint: k8s.go
imageRepository: k8s.gcr.io
networking:
  dnsDomain: k8s.go
  serviceSubnet: 10.96.0.0/12
  podSubnet: 192.168.0.0/16
```

1. kubeadm init

```
sudo kubeadm init --node-name master01 --apiserver-advertise-address 10.1.1.11 \
     --control-plane-endpoint k8s.go \
	 --pod-network-cidr 192.168.0.0/16 \
	 --service-dns-domain k8s.go \
	 --upload-certs \
	 -v 5
```

some options:
	
```
  --config string                        Path to a kubeadm configuration file.

  --upload-certs                         Upload control-plane certificates to the kubeadm-certs Secret.
  --node-name string                     Specify the node name.
  --apiserver-advertise-address string   The IP address the API Server will advertise it's listening on.If not set the default network interface will be used
  --control-plane-endpoint string        Specify a stable IP address or DNS name for the control plane.
  --pod-network-cidr string              Specify range of IP addresses for the pod network. If set, the control plane will automatically allocate CIDRs for every node.
  --service-dns-domain string            Use alternative domain for services, e.g. "myorg.internal". (default "cluster.local")
  --image-repository string              Choose a container registry to pull control plane images from (default "k8s.gcr.io")

Please note that the certificate-key gives access to cluster sensitive data, keep it secret!
As a safeguard, uploaded-certs will be deleted in two hours; If necessary, you can use
"kubeadm init phase upload-certs --upload-certs" to reload certs afterward.
```
   
2. Write the output join commands that are returned to a text file for later use.  

3. Install CNI

4. kubeadm join

* [kubeadm token - Kubernetes](https://kubernetes.io/docs/reference/setup-tools/kubeadm/kubeadm-token/)
* [issue with multiple interface when kubeadm join](https://github.com/kubernetes/kubernetes/issues/33618#issuecomment-516114988):

Config /etc/hosts

```
# must not config master02 in hosts when use vagrant
10.1.1.11 k8s.go
```

```
// control-plane
sudo kubeadm join k8s.go:6443 --token xxx.xxx \
  --discovery-token-ca-cert-hash sha256:xxx \
  --control-plane --certificate-key xxx \
  --apiserver-advertise-address 10.1.1.12 \
  --node-name master02 \
  -v 5

// worker
kubeadm join k8s.go:6443 --token xxx.xxx \
    --discovery-token-ca-cert-hash sha256:xxx \
	--node-name node01 \
	-v 5
```

# Web UI (Dashboard)

* [Web UI (Dashboard) - Kubernetes](https://kubernetes.io/docs/tasks/access-application-cluster/web-ui-dashboard/)
* [dashboard/creating-sample-user.md at master · kubernetes/dashboard · GitHub](https://github.com/kubernetes/dashboard/blob/master/docs/user/access-control/creating-sample-user.md)

## Create user

```
apiVersion: v1
kind: ServiceAccount
metadata:
  name: admin-user
  namespace: kubernetes-dashboard
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: admin-user
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-admin
subjects:
- kind: ServiceAccount
  name: admin-user
  namespace: kubernetes-dashboard
```


## Bearer Token

```
kubectl -n kubernetes-dashboard describe secret $(kubectl -n kubernetes-dashboard get secret | grep admin-user | awk '{print $1}')
```

# Tear down

* [Creating a single master cluster with kubeadm - Kubernetes](https://v1-13.docs.kubernetes.io/docs/setup/independent/create-cluster-kubeadm/#tear-down)

```
kubectl drain <node name> --delete-local-data --force --ignore-daemonsets
kubectl delete node <node name>
```

Then, on the node being removed, reset all kubeadm installed state:

```
kubeadm reset
```

Then:

```
iptables -F && iptables -t nat -F && iptables -t mangle -F && iptables -X
// or
ipvsadm -C
```

