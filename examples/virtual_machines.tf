provider "simplivity" {
   ovc_ip = "10.30.8.245"
   username = "sijeesh.kattumunda@demo.local"
   password = "Sijenov@2019"
}

resource "simplivity_vm" "example_vm" {
  name = "test_1"
  power_state = "off"
}

resource "simplivity_vm_clone" "example_vm_clone" {
  name = "test_1"
  new_vm_name = "terraform_test_clone"
}

resource "simplivity_vm" "example_vm_power" {
  name = simplivity_vm_clone.example_vm_clone.new_vm_name
  power_state = "on"
}
