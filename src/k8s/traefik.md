# [Traefik - The Cloud Native Edge Router](https://traefik.io/)

# Links

* [Kubernetes - Traefik](https://docs.traefik.io/user-guide/kubernetes/)

# Install

* [charts/stable/traefik at master · helm/charts · GitHub](https://github.com/helm/charts/tree/master/stable/traefik)

value:

```
$ cat value.yaml
rbac:
  enabled: true
dashboard:
  enabled: true
  domain: traefik.l
```

install:

```
$ helm install --namespace kube-system --values value.yaml stable/traefik
NAME:   agile-lobster
LAST DEPLOYED: Fri Aug 30 23:56:23 2019
NAMESPACE: kube-system
STATUS: DEPLOYED

RESOURCES:
==> v1/ClusterRole
NAME                   AGE
agile-lobster-traefik  3s

==> v1/ClusterRoleBinding
NAME                   AGE
agile-lobster-traefik  3s

==> v1/ConfigMap
NAME                        DATA  AGE
agile-lobster-traefik       1     3s
agile-lobster-traefik-test  1     3s

==> v1/Deployment
NAME                   READY  UP-TO-DATE  AVAILABLE  AGE
agile-lobster-traefik  0/1    0           0          1s

==> v1/Service
NAME                             TYPE          CLUSTER-IP     EXTERNAL-IP  PORT(S)                     AGE
agile-lobster-traefik            LoadBalancer  10.96.162.166  <pending>    80:31899/TCP,443:30909/TCP  1s
agile-lobster-traefik-dashboard  ClusterIP     10.102.35.190  <none>       80/TCP                      3s

==> v1/ServiceAccount
NAME                   SECRETS  AGE
agile-lobster-traefik  0        3s

==> v1beta1/Ingress
NAME                             HOSTS      ADDRESS  PORTS  AGE
agile-lobster-traefik-dashboard  traefik.l  80       1s


NOTES:

1. Get Traefik's load balancer IP/hostname:

     NOTE: It may take a few minutes for this to become available.

     You can watch the status by running:

         $ kubectl get svc agile-lobster-traefik --namespace kube-system -w

     Once 'EXTERNAL-IP' is no longer '<pending>':

         $ kubectl describe svc agile-lobster-traefik --namespace kube-system | grep Ingress | awk '{print $3}'

2. Configure DNS records corresponding to Kubernetes ingress resources to point to the load balancer IP/hostname found in step 1

```
