Terraform Provider for HPE SimpliVity
==================

- Website: https://www.terraform.io
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)

Maintainers
-----------

This provider plugin is maintained by:

* [HPE](http://hpe.com/simplivity)

Requirements
------------

- [Terraform](https://www.terraform.io/downloads.html) 0.12+


## Usage Example

```
# Configure HPE SimpliVity Provider
provider "simplivity" {
   ovc_ip = "<ovc_ip>"
   username = "<vcenter_username>"
   password = "<vcenter_password>"
}

# Power off VM
resource "simplivity_vm" "example_vm" {
  name = "vm_name"
  power_state = "off"
}

# Create a clone of the VM
resource "simplivity_vm_clone" "example_vm_clone" {
  name = "vm_name"
  new_vm_name = "vm_clone_name"
}

# Create a backup of the VM
resource "simplivity_vm_backup" "example_vm_backup" {
  name = "vm_backup_name"
  vm_name = "vm_name"
  omnistack_cluster_name = "cluster_name"
}
```

Building the provider
---------------------

Clone repository to: `$GOPATH/src/github.com/hashicorp/terraform-provider-simplivity`

```sh
$ mkdir -p $GOPATH/src/github.com/terraform-providers; cd $GOPATH/src/github.com/terraform-providers
$ git clone git@github.com:hashicorp/terraform-provider-simplivity
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/hashicorp/terraform-provider-simplivity
$ make build
```

Developing the provider
---------------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org)
installed on your machine. You can use [goenv](https://github.com/syndbg/goenv)
to manage your Go version. You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH),
as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`.

```sh
$ make build
```

### License

This project is licensed under the Apache 2.0 license.

## Version and changes

To view history and notes for this version, view the [Changelog](CHANGELOG.md)
