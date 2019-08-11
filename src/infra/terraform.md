# Terraform

Infrastructure as Code (IAC)

# Links

* [Providers - Terraform by HashiCorp](https://www.terraform.io/docs/providers/index.html)(majority)
* [Terraform Curriculum - HashiCorp Learn](https://learn.hashicorp.com/terraform/#getting-started)
* [Build Infrastructure | Terraform - HashiCorp Learn](https://learn.hashicorp.com/terraform/getting-started/build)
* [Documentation - Terraform by HashiCorp](https://www.terraform.io/docs/index.html)

# Installation

[Download Terraform - Terraform by HashiCorp](https://www.terraform.io/downloads.html)

```
go get github.com/hashicorp/terraform
```

# First Steps

* `mkdir foo && touch example.tf`
* `terraform init`  
  `Usage: terraform init [options] [DIR]`  
* `terraform apply`
* `terraform show`  
	Terraform also wrote some data into the `terraform.tfstate` file.   
* `terraform destroy`  

# Shell Tab-completion

```
terraform -install-autocomplete
terraform -uninstall-autocomplete
```

# CLI Configuration File

* [CLI Configuration - Terraform by HashiCorp](https://www.terraform.io/docs/commands/cli-config.html)

`.terraformrc` or `terraform.rc`

The CLI configuration file configures per-user settings for CLI behaviors,
which apply across all Terraform working directories. 

# Plugins

* [Command: init - Terraform by HashiCorp](https://www.terraform.io/docs/commands/init.html)

For providers distributed by HashiCorp, init will automatically download and
install plugins if necessary. Plugins can also be manually installed in the
user plugins directory, located at `~/.terraform.d/plugins`

##  Provider Plugin Cache 

```
// config in .terraformrc

plugin_cache_dir = "$HOME/.terraform.d/plugin-cache"
```

This directory must already exist before Terraform will cache plugins;
Terraform will not create the directory itself.

Still need to do `cp plugin-cache/* plugins/` manually to make `terraform plan`
work correctly, kind of bug.

# Providers

## Linode Provider 

* [Provider: Linode - Terraform by HashiCorp](https://www.terraform.io/docs/providers/linode/index.html)
* [Linode: A Beginner's Guide to Terraform](https://www.linode.com/docs/applications/configuration-management/beginners-guide-to-terraform/)
* [Linode API Documentation](https://developers.linode.com/api/v4/)
* [Use Terraform to Provision Linode Environments](https://www.linode.com/docs/applications/configuration-management/how-to-build-your-infrastructure-using-terraform-and-linode/)
* [GitHub - LinodeContent/terraform-provider-linode: A terraform plugin for linode](https://github.com/LinodeContent/terraform-provider-linode)

## API

* [images](https://api.linode.com/v4/images)
* [types](https://api.linode.com/v4/linode/types)
* [regions](https://api.linode.com/v4/regions)

# Configuration Language

#  Resource Syntax 

[Resources - Configuration Language - Terraform by HashiCorp](https://www.terraform.io/docs/configuration/resources.html)

A resource block declares a resource of a given type ("aws_instance") with
a given local name ("web"). 

# Input Variables

For all files which match `terraform.tfvars` or `*.auto.tfvars` present in the
current directory, Terraform automatically loads them to populate variables. 

We **don't recommend** saving usernames and password to version control, but
you can create a local secret variables file and use `-var-file` to load it.

##  Lists

```
# implicitly by using brackets [...]
variable "cidrs" { default = [] }

# explicitly
variable "cidrs" { type = "list" }

cidrs = [ "10.0.0.0/16", "10.1.0.0/16" ]
```

## Maps

# Modules

Think of modules as similar to functions in programming languages.

