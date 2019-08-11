# Kubernetes

# Links

* [Reference - Kubernetes](https://kubernetes.io/docs/reference/)
* [GitHub - kubernetes/minikube: Run Kubernetes locally](https://github.com/kubernetes/minikube)

# minikube

## Install

1. [Install and Set Up kubectl - Kubernetes](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
1. [Install Minikube - Kubernetes](https://kubernetes.io/docs/tasks/tools/install-minikube/)
1. [Running Kubernetes Locally via Minikube - Kubernetes](https://kubernetes.io/docs/setup/minikube/)

install kubectl

```
// macOS
brew install kubernetes-cli
brew install bash-completion@2

export BASH_COMPLETION_COMPAT_DIR=/usr/local/etc/bash_completion.d
[[ -r /usr/local/etc/profile.d/bash_completion.sh ]] && . /usr/local/etc/profile.d/bash_completion.sh

// Homebrew automatically doing this
echo 'source <(kubectl completion bash)' >>~/.bashrc
kubectl completion bash >/usr/local/etc/bash_completion.d/kubectl
```

install minikube

```
// macOS
brew cask install minikube

minikube start
```

# Code

* [community/contributors/devel at master · kubernetes/community · GitHub](https://github.com/kubernetes/community/tree/master/contributors/devel#readme)


