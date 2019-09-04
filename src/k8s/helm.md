# [Helm](https://helm.sh/)

# Links

# Install

```
curl -L https://git.io/get_helm.sh | bash
helm init

// example
$ helm repo update
$ helm install stable/traefik
```

# Issue

*. Error: Get https://10.96.0.1:443/api/v1/namespaces/kube-system/configmaps?labelSelector=OWNER%!D(MISSING)TILLER: dial tcp 10.96.0.1:443: i/o timeout

**Answers**:

It's a calico's error. pod_network_cidr can't be same as node's IP.


* Helm: Error: no available release name found

Answer:

[kubernetes - Helm: Error: no available release name found - Stack Overflow](https://stackoverflow.com/questions/43499971/helm-error-no-available-release-name-found/45306258#45306258)

```
kubectl create serviceaccount --namespace kube-system tiller
kubectl create clusterrolebinding tiller-cluster-rule --clusterrole=cluster-admin --serviceaccount=kube-system:tiller
kubectl patch deploy --namespace kube-system tiller-deploy -p '{"spec":{"template":{"spec":{"serviceAccount":"tiller"}}}}'
```

# Plugins

* [Helm Plugins](https://helm.sh/docs/related/#helm-plugins)


