# kubectl

# Links

* [Organizing Cluster Access Using kubeconfig Files - Kubernetes](https://kubernetes.io/docs/concepts/configuration/organize-cluster-access-kubeconfig/)
* [Mastering the KUBECONFIG file](https://ahmet.im/blog/mastering-kubeconfig/)

# kubeconfig

# KUBECONFIG

If the `KUBECONFIG` environment variable does exist, kubectl uses an effective
configuration that is the result of merging the files listed in the
`KUBECONFIG` environment variable.


# Completion

* [Install and Set Up kubectl - Kubernetes](https://kubernetes.io/docs/tasks/tools/install-kubectl/#kubectl-autocompletion-2)

```
// bash
echo 'source <(kubectl completion bash)' >>~/.bashrc
echo 'alias k=kubectl' >>~/.bashrc
echo 'complete -F __start_kubectl k' >>~/.bashrc

// zsh
echo 'source <(kubectl completion zsh)' >> ~/.zshrc
echo 'alias k=kubectl' >>~/.zshrc
echo 'complete -F __start_kubectl k' >>~/.zshrc

// zsh need these lines in .zshrc, or you will get an error like complete:13: command not found: compdef
autoload -Uz compinit
compinit
```


