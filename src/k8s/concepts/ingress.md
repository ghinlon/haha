# [Ingress - Kubernetes](https://kubernetes.io/docs/concepts/services-networking/ingress/)

# Links

* [Kubernetes Ingress 101: NodePort, Load Balancers, and Ingress Controllers](https://blog.getambassador.io/kubernetes-ingress-nodeport-load-balancers-and-ingress-controllers-6e29f1c44f2d?gi=33271a43bd46) 
* [kubernetes - Ingress vs Load Balancer - Stack Overflow](https://stackoverflow.com/questions/45079988/ingress-vs-load-balancer)

# Does an Ingress has a port ?

I know after all, it must be a silly question.  
  Wed Sep 25 00:56:51 CST 2019

* `Ingress` is just routers.
* `LoadBalancer` service type automatically deploys an external load balancer to
  your `NodePort`.
* An `Ingress Controller` is responsible for reading the `Ingress Resource`
  information and processing that data accordingly. 


**So, the question is, what makes it possible that both `foo.com` and `bar.com`
go in cluster through the same port, port `80` ?**

**what I know is, that different services can't bind to the same `NodePort`**


# The Point Of LoadBalancer

``` 
Regardless of your ingress strategy, you probably will need to start with
an external load balancer. This load balancer will then route traffic to
a Kubernetes service (or ingress) on your cluster that will perform
service-specific routing. In this set up, your load balancer provides a stable
endpoint (IP address) for external traffic to access.

Both ingress controllers and Kubernetes services require an external load
balancer, and, as previously discussed, NodePorts are not designed to be
directly used for production.  
```

# Can I just use multiple DNS record to every node ?

* [Single A records, multiple IP addresses - DNS - Spiceworks](https://community.spiceworks.com/topic/336568-single-a-records-multiple-ip-addresses?page=1#entry-2201467)

```
You can have multiple A records for the same FQDN. That's how DNS round-robin works.

for example - nslookup www.google.com

Non-authoritative answer:
Name:    www.google.com
Addresses:  2a00:1450:400b:c02::63
          74.125.24.147
          74.125.24.99
          74.125.24.103
          74.125.24.104
          74.125.24.105
          74.125.24.106
```

# Conclusion

At this point, I think I don't need a LoadBalancer.

	I think it's wrong.  
	Fri Sep 27 01:17:53 CST 2019







