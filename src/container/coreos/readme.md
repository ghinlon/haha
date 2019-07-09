# [CoreOS](https://coreos.com/)

# Links

* [Container Linux Quick Start | Container Linux by CoreOS](https://coreos.com/os/docs/latest/quickstart.html)
* [CoreOS · GitHub](https://github.com/coreos/)
* [Vagrant](https://coreos.com/os/docs/latest/booting-on-vagrant.html)
* [stable.release.core-os.net/amd64-usr/current/](https://stable.release.core-os.net/amd64-usr/current/)

# Install with Vagrant and ignition

* [GitHub - coreos/coreos-vagrant: Minimal Vagrantfile for Container Linux](https://github.com/coreos/coreos-vagrant#provisioning-with-ignition-virtualbox-provider-default)
* [CL Config Dynamic Data](https://coreos.com/os/docs/latest/dynamic-data.html)

`Vagrant-Virtualbox` currently only support these Dynamic Data: `PRIVATE_IPV4`
and `HOSTNAME`


```
curl https://discovery.etcd.io/new\?size\=3
// replace <token> in the cl.conf file with the generated token from the curl command.
// replace <PUBLIC_IPV4> to <PRIVATE_IPV4> in the cl.conf
ct --platform=vagrant-virtualbox < cl.conf > config.ign

vagrant up
```

The Vagrantfile will parse a `config.rb` file containing a set of options used
to configure your CoreOS cluster. See `config.rb.sample` for more information.

# etcd

## Docker Forwarding

By setting the `$expose_docker_tcp` configuration value in `config.rb.sample`

then:

```
vagrant reload --provision
```

Then you can then use the docker command from your local shell by setting DOCKER_HOST:

`export DOCKER_HOST=tcp://localhost:2375`

test etcd:

```
 vagrant ssh core-01

core@core-01 ~ $ systemctl status etcd-member
core@core-01 ~ $ etcdctl set /message Hello
Hello
core@core-01 ~ $ etcdctl get /message
Hello
core@core-01 ~ $  etcdctl rm /message
PrevNode.Value: Hello
```

# Easy development/testing cluster

[CoreOS Container Linux Cluster Architectures | Linux Clustering](https://coreos.com/os/docs/latest/cluster-architectures.html)





