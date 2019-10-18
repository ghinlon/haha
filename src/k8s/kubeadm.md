# kubeadm quickstart

# Links

* [Setup a kubernetes cluster with kubeadm | loodse | Katacoda](https://www.katacoda.com/loodse/courses/kubernetes/kubernetes-03-cluster-setup)
* [kubeadm config - Kubernetes](https://kubernetes.io/docs/reference/setup-tools/kubeadm/kubeadm-config/)

# Steps

* [Installing kubeadm - Kubernetes](https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/install-kubeadm/)
* [Get Docker Engine - Community for CentOS | Docker Documentation](https://docs.docker.com/install/linux/docker-ce/centos/)
* [Install and Set Up kubectl - Kubernetes](https://kubernetes.io/docs/tasks/tools/install-kubectl/)

## Install CRI

## systemd

## swapoff -a 

and also delete swap in `/etc/fstab`

## Installing kubeadm, kubelet and kubectl

There are some nodes on centos7:

* Setting SELinux in permissive

	```
	# Set SELinux in permissive mode (effectively disabling it)
	setenforce 0
	sed -i 's/^SELINUX=enforcing$/SELINUX=permissive/' /etc/selinux/config
	```
  
* Some users on RHEL/CentOS 7 have reported issues with traffic being routed
  incorrectly due to iptables being bypassed. You should ensure
  net.bridge.bridge-nf-call-iptables is set to 1 in your sysctl config, e.g.

	```
	cat <<EOF >  /etc/sysctl.d/k8s.conf
	net.bridge.bridge-nf-call-ip6tables = 1
	net.bridge.bridge-nf-call-iptables = 1
	EOF
	sysctl --system
	```

* Make sure that the `br_netfilter` module is loaded before this step. This can
  be done by running `lsmod | grep br_netfilter`. To load it explicitly call
  `modprobe br_netfilter`.

```
cat <<EOF > /etc/yum.repos.d/kubernetes.repo
[kubernetes]
name=Kubernetes
baseurl=https://packages.cloud.google.com/yum/repos/kubernetes-el7-x86_64
enabled=1
gpgcheck=1
repo_gpgcheck=1
gpgkey=https://packages.cloud.google.com/yum/doc/yum-key.gpg https://packages.cloud.google.com/yum/doc/rpm-package-key.gpg
EOF

yumdownloader --resolve kubelet kubeadm kubectl --disableexcludes=kubernetes
sudo um localinstall *

systemctl enable --now kubelet
```

## Config --nodeip

* [kubeadm should make the --node-ip option available · Issue #203 · kubernetes/kubeadm · GitHub](https://github.com/kubernetes/kubeadm/issues/203)

**This Step is important, before come up with this, several days life have been
taken, due to vagrant's default NAT interface**  

Mon Oct 07 21:01:41 CST 2019

In `/etc/sysconfig/kubelet`:  
on `ubuntu` use `/etc/default/kubelet`:

```
KUBELET_EXTRA_ARGS=--node-ip=10.1.1.11
```

then

```
systemctl daemon-reload
systemctl restart kubelet
```

The kubelet is now restarting every few seconds, as it waits in a crashloop for
kubeadm to tell it what to do.

## docker load images

```
// save
docker save $(docker images --format '{{.Repository}}:{{.Tag}}') -o allinone.tar

// load
docker load -i allinone.tar
```

## On master

```
// 1
$ kubeadm config images pull
// output
[config/images] Pulled k8s.gcr.io/kube-apiserver:v1.16.0
[config/images] Pulled k8s.gcr.io/kube-controller-manager:v1.16.0
[config/images] Pulled k8s.gcr.io/kube-scheduler:v1.16.0
[config/images] Pulled k8s.gcr.io/kube-proxy:v1.16.0
[config/images] Pulled k8s.gcr.io/pause:3.1
[config/images] Pulled k8s.gcr.io/etcd:3.3.15-0
[config/images] Pulled k8s.gcr.io/coredns:1.6.2

// 2
master $ kubeadm init --pod-network-cidr=192.168.0.0/16
// output
I0915 13:37:12.697153   16228 version.go:240] remote version is much newer: v1.15.3; falling back to: stable-1.14
[init] Using Kubernetes version: v1.14.6
[preflight] Running pre-flight checks
[preflight] Pulling images required for setting up a Kubernetes cluster
[preflight] This might take a minute or two, depending on the speed of your internet connection
[preflight] You can also perform this action in beforehand using 'kubeadm config images pull'
[kubelet-start] Writing kubelet environment file with flags to file "/var/lib/kubelet/kubeadm-flags.env"
[kubelet-start] Writing kubelet configuration to file "/var/lib/kubelet/config.yaml"
[kubelet-start] Activating the kubelet service
[certs] Using certificateDir folder "/etc/kubernetes/pki"
[certs] Generating "front-proxy-ca" certificate and key
[certs] Generating "front-proxy-client" certificate and key
[certs] Generating "etcd/ca" certificate and key
[certs] Generating "apiserver-etcd-client" certificate and key
[certs] Generating "etcd/peer" certificate and key
[certs] etcd/peer serving cert is signed for DNS names [master localhost] and IPs [172.17.0.35 127.0.0.1 ::1]
[certs] Generating "etcd/healthcheck-client" certificate and key
[certs] Generating "etcd/server" certificate and key
[certs] etcd/server serving cert is signed for DNS names [master localhost] and IPs [172.17.0.35 127.0.0.1 ::1]
[certs] Generating "ca" certificate and key
[certs] Generating "apiserver" certificate and key
[certs] apiserver serving cert is signed for DNS names [master kubernetes kubernetes.default kubernetes.default.svc kubernetes.default.svc.cluster.local] and IPs [10.96.0.1 172.17.0.35]
[certs] Generating "apiserver-kubelet-client" certificate and key
[certs] Generating "sa" key and public key
[kubeconfig] Using kubeconfig folder "/etc/kubernetes"
[kubeconfig] Writing "admin.conf" kubeconfig file
[kubeconfig] Writing "kubelet.conf" kubeconfig file
[kubeconfig] Writing "controller-manager.conf" kubeconfig file
[kubeconfig] Writing "scheduler.conf" kubeconfig file
[control-plane] Using manifest folder "/etc/kubernetes/manifests"
[control-plane] Creating static Pod manifest for "kube-apiserver"
[control-plane] Creating static Pod manifest for "kube-controller-manager"
[control-plane] Creating static Pod manifest for "kube-scheduler"
[etcd] Creating static Pod manifest for local etcd in "/etc/kubernetes/manifests"
[wait-control-plane] Waiting for the kubelet to boot up the control plane as static Pods from directory "/etc/kubernetes/manifests". This can take up to 4m0s
[apiclient] All control plane components are healthy after 17.503286 seconds
[upload-config] storing the configuration used in ConfigMap "kubeadm-config" in the "kube-system" Namespace
[kubelet] Creating a ConfigMap "kubelet-config-1.14" in namespace kube-system with the configuration for the kubelets inthe cluster
[upload-certs] Skipping phase. Please see --experimental-upload-certs
[mark-control-plane] Marking the node master as control-plane by adding the label "node-role.kubernetes.io/master=''"
[mark-control-plane] Marking the node master as control-plane by adding the taints [node-role.kubernetes.io/master:NoSchedule]
[bootstrap-token] Using token: z36he9.hlio8gqvxognvuuk
[bootstrap-token] Configuring bootstrap tokens, cluster-info ConfigMap, RBAC Roles
[bootstrap-token] configured RBAC rules to allow Node Bootstrap tokens to post CSRs in order for nodes to get long term certificate credentials
[bootstrap-token] configured RBAC rules to allow the csrapprover controller automatically approve CSRs from a Node Bootstrap Token
[bootstrap-token] configured RBAC rules to allow certificate rotation for all node client certificates in the cluster
[bootstrap-token] creating the "cluster-info" ConfigMap in the "kube-public" namespace
[addons] Applied essential addon: CoreDNS
[addons] Applied essential addon: kube-proxy

Your Kubernetes control-plane has initialized successfully!

To start using your cluster, you need to run the following as a regular user:

  mkdir -p $HOME/.kube
  sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
  sudo chown $(id -u):$(id -g) $HOME/.kube/config

You should now deploy a pod network to the cluster.
Run "kubectl apply -f [podnetwork].yaml" with one of the options listed at:
  https://kubernetes.io/docs/concepts/cluster-administration/addons/

Then you can join any number of worker nodes by running the following on each as root:

kubeadm join 172.17.0.35:6443 --token z36he9.hlio8gqvxognvuuk \
    --discovery-token-ca-cert-hash sha256:a4f314dd5f59656e0a42699b5ac5ac962e6b29914f4d2b0732d202dd997cf139

// 3
mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config

// 4
curl https://docs.projectcalico.org/v3.9/manifests/calico-typha.yaml -o calico.yaml
kubectl apply -f calico.yaml
```

## On node

```
// 1
node01 $ kubeadm join 172.17.0.35:6443 --token z36he9.hlio8gqvxognvuuk \
>     --discovery-token-ca-cert-hash sha256:a4f314dd5f59656e0a42699b5ac5ac962e6b29914f4d2b0732d202dd997cf139
// output
[preflight] Running pre-flight checks
[preflight] Reading configuration from the cluster...
[preflight] FYI: You can look at this config file with 'kubectl -n kube-system get cm kubeadm-config -oyaml'
[kubelet-start] Downloading configuration for the kubelet from the "kubelet-config-1.14" ConfigMap in the kube-system namespace
[kubelet-start] Writing kubelet configuration to file "/var/lib/kubelet/config.yaml"
[kubelet-start] Writing kubelet environment file with flags to file "/var/lib/kubelet/kubeadm-flags.env"
[kubelet-start] Activating the kubelet service
[kubelet-start] Waiting for the kubelet to perform the TLS Bootstrap...

This node has joined the cluster:
* Certificate signing request was sent to apiserver and a response was received.
* The Kubelet was informed of the new secure connection details.

Run 'kubectl get nodes' on the control-plane to see this node join the cluster.
```

# More about Networking

* [Installing Addons - Kubernetes](https://kubernetes.io/docs/concepts/cluster-administration/addons/)
* [Cluster Networking - Kubernetes](https://kubernetes.io/docs/concepts/cluster-administration/networking/)

