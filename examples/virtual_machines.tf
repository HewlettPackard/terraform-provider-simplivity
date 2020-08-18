provider "simplivity" {
   ovc_ip = "10.20.4.245"
   username = "Sijeesh.Kattumunda@demo.local"
   password = "sijeaug@2020"
}

resource "simplivity_vm" "example_vm" {
  name = "test_1"
  power_state = "off"
}

resource "simplivity_vm_clone" "example_vm_clone" {
  name = "test_1"
  new_vm_name = "terraform_test_clone"
}

resource "simplivity_vm_move" "example_vm_move" {
  name = "terraform_test_clone"
  new_vm_name = "terraform_test_move"
  datastore_name = ""
}

resource "simplivity_vm_backup" "example_vm_backup" {
  name = "test_1_backup"
  vm_name = "test_1"
  omnistack_cluster_name = "RemoteCluster"
}

resource "simplivity_vm_backup" "example_vm_backup_local" {
  name = "test_1_back"
  vm_name = "test_1"
}

data "simplivity_vm" "example_vm_data" {
  name = "test_1"
}

data "simplivity_vm_backup" "example_vm_backup_data" {
  name = "test_1_backup1"
}
