# grafana

# Links

# Install

* [charts/stable/grafana at master · helm/charts · GitHub](https://github.com/helm/charts/tree/master/stable/grafana)


values:

```

```


install:

```
helm install -f value.yaml stable/grafana


NOTES:
1. Get your 'admin' user password by running:

   kubectl get secret --namespace default nomadic-horse-grafana -o jsonpath="{.data.admin-password}" | base64 --decode ; echo

2. The Grafana server can be accessed via port 80 on the following DNS name from within your cluster:

   nomadic-horse-grafana.default.svc.cluster.local

   Get the Grafana URL to visit by running these commands in the same shell:

     export POD_NAME=$(kubectl get pods --namespace default -l "app=grafana,release=nomadic-horse" -o jsonpath="{.items[0].metadata.name}")
     kubectl --namespace default port-forward $POD_NAME 3000

3. Login with the password from step 1 and the username: admin
#################################################################################
######   WARNING: Persistence is disabled!!! You will lose your data when   #####
######            the Grafana pod is terminated.                            #####
#################################################################################

```

