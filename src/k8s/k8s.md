# Kubernetes

# Links

* [Reference - Kubernetes](https://kubernetes.io/docs/reference/)

# Kubernetes Setup Using Ansible and Vagrant

* [Kubernetes Setup Using Ansible and Vagrant - Kubernetes](https://kubernetes.io/blog/2019/03/15/kubernetes-setup-using-ansible-and-vagrant/)

# Concepts

* [Concepts - Kubernetes](https://kubernetes.io/docs/concepts/)
* [Kubernetes Design Principles: Understand the Why - Saad Ali, Google - YouTube](https://www.youtube.com/watch?v=ZuIQurh_kDk)(**Good**)

# Services

* [Kubernetes - Services Explained - YouTube](https://www.youtube.com/watch?v=5lzUpDtmWgM&list=PL2We04F3Y_43dAehLMT5GxJhtk3mJtkl5&index=10)

# Ingress

* [Kubernetes Ingress Explained For Beginners - YouTube](https://www.youtube.com/watch?v=VicH6KojwCI&feature=youtu.be)

# Volumes

## Type of Volumes

* configMap
* [Configure a Pod to Use a ConfigMap - Kubernetes](https://kubernetes.io/docs/tasks/configure-pod-container/configure-pod-configmap/)(**good**)

```
kubectl create configmap <map-name> <data-source>
kubectl create configmap game-config --from-file=configure-pod-container/configmap/
kubectl create configmap game-config-3 --from-file=<my-key-name>=<path-to-file>
kubectl create configmap special-config --from-literal=special.how=very --from-literal=special.type=charm
```


# Kubernetes Basics

* [Learn Kubernetes Basics - Kubernetes](https://kubernetes.io/docs/tutorials/kubernetes-basics/)
* [Deployments - Kubernetes](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/)



# Code

* [community/contributors/devel at master · kubernetes/community · GitHub](https://github.com/kubernetes/community/tree/master/contributors/devel#readme)


# Cheat Sheet

* [kubectl Cheat Sheet - Kubernetes](https://kubernetes.io/docs/reference/kubectl/cheatsheet/)


# Install and Set Up kubectl

* [Install and Set Up kubectl - Kubernetes](https://kubernetes.io/docs/tasks/tools/install-kubectl/#install-kubectl-on-macos)
* [Configure Access to Multiple Clusters - Kubernetes](https://kubernetes.io/docs/tasks/access-application-cluster/configure-access-multiple-clusters/)

```
source <(kubectl completion bash) # setup autocomplete in bash into the current shell, bash-completion package should be installed first.
echo "source <(kubectl completion bash)" >> ~/.bashrc # add autocomplete permanently to your bash shell.

alias k=kubectl
complete -F __start_kubectl k


source <(kubectl completion zsh)  # setup autocomplete in zsh into the current shell
echo "if [ $commands[kubectl] ]; then source <(kubectl completion zsh); fi" >> ~/.zshrc # add autocomplete permanently to your zsh shell
```

# Install minikube

```
curl -Lo minikube https://storage.googleapis.com/minikube/releases/latest/minikube-darwin-amd64 \
  && chmod +x minikube

sudo mv minikube /usr/local/bin

minikube start
minikube stop
minikube delete

minikube dashboard
// or
kubectl proxy
// access the Dashboard over this url
http://127.0.0.1:8001/api/v1/namespaces/kube-system/services/kubernetes-dashboard:/proxy/#!/overview?namespace=default 
```














