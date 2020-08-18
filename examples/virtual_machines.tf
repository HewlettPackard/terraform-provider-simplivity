provider "simplivity" {
   ovc_ip = "<ovc_ip>"
   username = "<vcenter_username>"
   password = "<vcenter_password>"
}

resource "simplivity_vm" "example_vm" {
  name = "vm_name"
  power_state = "off"
}

resource "simplivity_vm_clone" "example_vm_clone" {
  name = "vm_name"
  new_vm_name = "vm_clone_name"
}

resource "simplivity_vm_move" "example_vm_move" {
  name = "vm_name"
  new_vm_name = "vm_moved_name"
  datastore_name = ""
}

resource "simplivity_vm_backup" "example_vm_backup" {
  name = "backup_name"
  vm_name = "vm_name"
  omnistack_cluster_name = "cluster_name"
}

resource "simplivity_vm_backup" "example_vm_backup_local" {
  name = "backup_name"
  vm_name = "vm_name"
}

data "simplivity_vm" "example_vm_data" {
  name = "vm_name"
}

data "simplivity_vm_backup" "example_vm_backup_data" {
  name = "vm_backup_name"
}
