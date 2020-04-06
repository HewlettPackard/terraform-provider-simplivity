provider "simplivity" {
   ovc_ip = "10.30.8.245"
   username = "sijeesh.kattumunda@demo.local"
   password = "Sijenov@2019"
}

#resource "simplivity_vm" "example_vm" {
#  name = "test_1"
#  power_state = "off"
#}

#resource "simplivity_vm_clone" "example_vm_clone" {
#  name = "test_1"
#  new_vm_name = "terraform_test_clone"
#}

#resource "simplivity_vm_move" "example_vm_move" {
#  name = "test_1"
#  new_vm_name = "terraform_test_move"
#  datastore_name = "SVT_Montreal01"
#}

#resource "simplivity_vm" "example_vm_power" {
#  name = simplivity_vm_clone.example_vm_clone.new_vm_name
#  power_state = "on"
#}

#resource "simplivity_vm_backup" "example_vm_backup" {
#  name = "test_1_backup"
#  vm_name = "test_1"
#  omnistack_cluster_name = "RemoteCluster"
#}

data "simplivity_vm" "example_vm_data" {
  name = "test_1"
}

data "simplivity_vm_backup" "example_vm_backup_data" {
  name = "test_1_backup"
}

resource "simplivity_vm_backup" "example_vm_backup" {
  name = data.simplivity_vm.example_vm_data.name
  vm_name = "test_1"
}
