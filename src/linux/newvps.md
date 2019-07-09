# NewVPS

# Links

* [Get Docker CE for Ubuntu | Docker Documentation](https://docs.docker.com/install/linux/docker-ce/ubuntu/#install-using-the-repository)

# First Step

**CentOS 7**

1. `useradd -ms /bin/bash xx && passwd xx`
2. `vim /etc/sudoers`
3. `vim /etc/ssh/sshd_config` 
	```
	# egrep '^Port|^PermitRootLogin|^P.*Auth' /etc/ssh/sshd_config

	Port 26271
	PermitRootLogin no
	PubkeyAuthentication yes
	PasswordAuthentication no
	```  
4. `yum install policycoreutils-python`; `semanage port -a -t ssh_port_t -p tcp 26271`
5. `firewall-cmd --zone=public --add-port=26271/tcp --permanent` then **must** `firewall-cmd --reload`
6. `sshcopykey.sh xx@host`
7. `systemctl restart sshd`
8. **Don't logout.** `ssh xx@host` in other tab to make sure it's OK.

# install docker on Ubuntu 18.04

**run screen first**

```
sudo apt-get update

sudo apt-get install \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg-agent \
    software-properties-common

curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -

// x86_64
sudo add-apt-repository \
   "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
   $(lsb_release -cs) \
   stable"
sudo apt-get update

// install
sudo apt-get install docker-ce docker-ce-cli containerd.io

// test
sudo docker run hello-world

// permission
// MUST logout once after.
sudo usermod -a -G docker $USER

// run
docker run \
        -p 443:8388 \
        -e PASSWORD="passwd" \
        -e METHOD="aes-128-gcm" \
        -e ARGS=-v \
        -d \
        --restart always \
        shadowsocks/shadowsocks-libev
```

# UPDATE GCE VM CPU

[GCM教學：如何變更VM執行個體的CPU數量、記憶體大小或硬碟大小？ – 電癮院](https://mrtang.tw/blog/post/how-to-change-a-machine-type-on-google-compute-engine)

* stop the VM, then u can edit it.
* f1-micro's speed is about 3Mbps, it's shared CPU(0.2vCPU~0.5vCPU), ~ $5/m
* g1-small's speed is ~ 60Mbps, it's 1 vCPU. ~ $15/m




