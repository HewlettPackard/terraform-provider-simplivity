package simplivity

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSimplivityVirtualMachineMove() *schema.Resource {
	return &schema.Resource{
		Create: resourceSimplivityVirtualMachineMoveCreate,
		Read:   resourceSimplivityVirtualMachineRead,
		Update: resourceSimplivityVirtualMachineCreateOrUpdate,
		Delete: resourceSimplivityVirtualMachineDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"new_vm_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"datastore_name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceSimplivityVirtualMachineMoveCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Config).Client

	name := d.Get("name").(string)
	new_vm_name := d.Get("new_vm_name").(string)
	datastore_name := d.Get("datastore_name").(string)

	vm, err := client.VirtualMachines.GetByName(name)
	if err != nil {
		return err
	}

	datastore, err := client.Datastores.GetByName(datastore_name)
	if err != nil {
		return err
	}

	_, err = vm.Move(new_vm_name, datastore)
	if err != nil {
		return err
	}

	d.Set("name", new_vm_name)

	d.SetId(new_vm_name)

	return resourceSimplivityVirtualMachineRead(d, meta)
}
