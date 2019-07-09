# [Vagrant by HashiCorp](https://www.vagrantup.com/)

# Links

* [Getting Started - Vagrant by HashiCorp](https://www.vagrantup.com/intro/getting-started/index.html)
* [Discover Vagrant Boxes - Vagrant Cloud](https://app.vagrantup.com/boxes/search)

# Basis

```
vagrant up
vagrant ssh core-01 -- -A

vagrant box update
vagrant plugin update vagrant-ignition

vagrant destroy -f
```
# Ubuntu

* [Vagrant box ubuntu/disco64 - Vagrant Cloud](https://app.vagrantup.com/ubuntu/boxes/disco64)

```
vagrant init ubuntu/disco64
vagrant up
```

* [Vagrant box ubuntu/bionic64 - Vagrant Cloud](https://app.vagrantup.com/ubuntu/boxes/bionic64)

```
vagrant init ubuntu/bionic64
vagrant up
```

# centos7

* [centos - Vagrant Cloud](https://app.vagrantup.com/centos)
* [Vagrant box centos/7 - Vagrant Cloud](https://app.vagrantup.com/centos/boxes/7)

```
vagrant init centos/7
vagrant up
```

# Forwarded Ports 

* [Forwarded Ports - Networking - Vagrant by HashiCorp](https://www.vagrantup.com/docs/networking/forwarded_ports.html)

This will allow accessing port 80 on the guest via port 8080 on the host.

```
Vagrant.configure("2") do |config|
  config.vm.network "forwarded_port", guest: 80, host: 8080
  config.vm.network "forwarded_port", guest: 2003, host: 12003, protocol: "tcp"
  config.vm.network "forwarded_port", guest: 2003, host: 12003, protocol: "udp"
end
```

# Bridged Networking

```
Vagrant.configure("2") do |config|
  config.vm.network "public_network"
end
```

# Multi-Machine

* [Multi-Machine - Vagrant by HashiCorp](https://www.vagrantup.com/docs/multi-machine/)

```
```
config.vm.box = "ubuntu/disco64"
config.vm.provision :shell, path: "ss.sh"


config.vm.define "airport" do |airport|
	config.vm.network "public_network"
	airport.vm.network "private_network", ip: "192.168.33.2"
end

config.vm.define "vps" do |vps|
	vps.vm.network "private_network", ip: "192.168.33.1"
end
```
```
