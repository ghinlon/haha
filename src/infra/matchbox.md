# matchbox

# Links

* [GitHub - coreos/matchbox: Network boot and provision Container Linux clusters (e.g. etcd3, Kubernetes, more)](https://github.com/coreos/matchbox)

# matchbox concepts

* [Matchbox concepts](https://coreos.com/matchbox/docs/latest/matchbox.html)

![matchbox concepts](imgs/overview.png)  
Image from: [Matchbox concepts](https://coreos.com/matchbox/docs/latest/matchbox.html)

# Lifecycle of a physical machine

* [Lifecycle of a physical machine](https://coreos.com/matchbox/docs/latest/machine-lifecycle.html)

![Lifecycle of a physical machine](imgs/machine-lifecycle.png)  
Image from: [Lifecycle of a physical machine](https://coreos.com/matchbox/docs/latest/machine-lifecycle.html)

# Installation

[Releases · coreos/matchbox · GitHub](https://github.com/coreos/matchbox/releases)

This link is useful, very straightforward:

[matchbox/deployment.md at master · coreos/matchbox · GitHub](https://github.com/coreos/matchbox/blob/master/Documentation/deployment.md)

## Download Container Linux (optional)

[CoreOS Container Linux Release](https://coreos.com/releases/)

```
./scripts/get-coreos stable 2079.3.0 .     # note the "." 3rd argument
sudo cp -r coreos /var/lib/matchbox/assets

// verify
curl http://matchbox.example.com:8080/assets/coreos/2079.3.0/
```

# Network setup

[matchbox/network-setup.md at master · coreos/matchbox · GitHub](https://github.com/coreos/matchbox/blob/master/Documentation/network-setup.md)

# Getting Started

## terraform-provider-matchbox

* [Download Terraform - Terraform by HashiCorp](https://www.terraform.io/downloads.html)
* [Releases · poseidon/terraform-provider-matchbox · GitHub](https://github.com/poseidon/terraform-provider-matchbox/releases)

## Example simple-install

* [matchbox/getting-started.md at master · poseidon/matchbox · GitHub](https://github.com/poseidon/matchbox/blob/master/Documentation/getting-started.md#first-cluster)

After change change `terraform.tfvars.example` to `terraform.tfvars`, and
modified it. then run `terraform apply`, some files will be generated:

```
$ tree /var/lib/matchbox/
/var/lib/matchbox/
├── assets
├── groups
│   ├── default.json
│   └── node1.json
├── ignition
│   ├── coreos-install.yaml.tmpl
│   └── simple.yaml.tmpl
└── profiles
    ├── coreos-install.json
    └── simple.json

4 directories, 6 files
```

and these links will be accessable:

Matchbox serves configs to machines and respects query parameters, if you're
interested:


* iPXE default - [/ipxe](http://matchbox.example.com:8080/ipxe)
* Ignition default - [/ignition](http://matchbox.example.com:8080/ignition)
* Ignition post-install - [/ignition?os=installed](http://matchbox.example.com:8080/ignition?os=installed)
* GRUB default - [/grub](http://matchbox.example.com:8080/grub)

## terraform-provider-matchbox

**Provider**

```
// Configure the matchbox provider
provider "matchbox" {
  endpoint = "${var.matchbox_rpc_endpoint}"
  client_cert = "${file("~/.matchbox/client.crt")}"
  client_key = "${file("~/.matchbox/client.key")}"
  ca         = "${file("~/.matchbox/ca.crt")}"
}
```
**Resources**

```
resource "matchbox_profile" "coreos-install" {
  name = "coreos-install"
  kernel = "https://stable.release.core-os.net/amd64-usr/current/coreos_production_pxe.vmlinuz"
  initrd = [
    "https://stable.release.core-os.net/amd64-usr/current/coreos_production_pxe_image.cpio.gz"
  ]
  args = [
    "coreos.config.url=${var.matchbox_http_endpoint}/ignition?uuid=$${uuid}&mac=$${mac:hexhyp}",
    "coreos.first_boot=yes",
    "console=tty0",
    "console=ttyS0",
  ]
  container_linux_config = "${file("./cl/coreos-install.yaml.tmpl")}"
}
```

This group does not have a selector block, so any machines which network boot
from matchbox will match this group and be provisioned using the coreos-install
profile. 

```
resource "matchbox_group" "default" {
  name = "default"
  profile = "${matchbox_profile.coreos-install.name}"
  # no selector means all machines can be matched
  metadata {
    ignition_endpoint = "${var.matchbox_http_endpoint}/ignition"
    ssh_authorized_key = "${var.ssh_authorized_key}"
  }
}
```



