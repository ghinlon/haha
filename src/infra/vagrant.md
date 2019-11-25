# [Vagrant by HashiCorp](https://www.vagrantup.com/)

# Links

* [Getting Started - Vagrant by HashiCorp](https://www.vagrantup.com/intro/getting-started/index.html)
* [Discover Vagrant Boxes - Vagrant Cloud](https://app.vagrantup.com/boxes/search)

# Install

* [Download - Vagrant by HashiCorp](https://www.vagrantup.com/downloads.html)

```
brew cask install vagrant
```

# Basis

```
vagrant up
vagrant ssh core-01 -- -A

vagrant plugin update vagrant-ignition

vagrant destroy -f
```

# boxes

* [Box Versioning - Vagrant by HashiCorp](https://www.vagrantup.com/docs/boxes/versioning.html)

```
vagrant box list
vagrant box update
```

# Archlinux

* [Vagrant box archlinux/archlinux - Vagrant Cloud](https://app.vagrantup.com/archlinux/boxes/archlinux)

```
vagrant init archlinux/archlinux
vagrant up
```

# Ubuntu

* [Vagrant box ubuntu/disco64 - Vagrant Cloud](https://app.vagrantup.com/ubuntu/boxes/disco64)

```
// ubuntu 19
vagrant init ubuntu/disco64
vagrant up
```

* [Vagrant box ubuntu/bionic64 - Vagrant Cloud](https://app.vagrantup.com/ubuntu/boxes/bionic64)

```
// ubuntu 18
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

This will bring up 4 machine:

```
% cat Vagrantfile
IMAGE_NAME = "centos/7"
M = 2
N = 2

Vagrant.configure("2") do |config|

    config.vm.provider "virtualbox" do |v|
        v.memory = 1024
        v.cpus = 2
    end

   (1..M).each do |i|
      config.vm.define "master#{i}" do |master|
        master.vm.box = IMAGE_NAME
        master.vm.network "private_network", ip: "10.1.1.#{i+10}"
        master.vm.hostname = "master#{i}"
      end
        end

    (1..N).each do |i|
      config.vm.define "node#{i}" do |node|
        node.vm.box = IMAGE_NAME
        node.vm.network "private_network", ip: "10.1.1.#{i + 12}"
        node.vm.hostname = "node#{i}"
      end
    end
end
```

# How to copy file from a Vagrant machine to local host

* [linux - How to copy file from a Vagrant machine to local host - Stack Overflow](https://stackoverflow.com/a/46203304)

```
vagrant ssh-config > config.txt
scp -F config.txt default:/path/to/file .
```

# Providers

## VirtualBox

* [VirtualBox Provider - Vagrant by HashiCorp](https://www.vagrantup.com/docs/virtualbox/)

There are some convenience shortcuts for memory and CPU settings:

```
config.vm.provider "virtualbox" do |v|
  v.memory = 1024
  v.cpus = 2
end
```


# Plugins

## vagrant-disksize

* [GitHub - sprotheroe/vagrant-disksize: Vagrant plugin to resize disks in VirtualBox](https://github.com/sprotheroe/vagrant-disksize)

Installation

```
vagrant plugin install vagrant-disksize
```

Usage

```
Vagrant.configure('2') do |config|
  config.vm.box = 'ubuntu/xenial64'
  config.disksize.size = '50GB'
end
```

