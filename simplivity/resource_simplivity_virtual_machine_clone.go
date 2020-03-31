package simplivity

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSimplivityVirtualMachineClone() *schema.Resource {
	return &schema.Resource{
		Create: resourceSimplivityVirtualMachineCloneCreate,
		Read:   resourceSimplivityVirtualMachineRead,
		Update: resourceSimplivityVirtualMachineUpdate,
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
			"app_consistent": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func resourceSimplivityVirtualMachineCloneCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Config).Client

	new_vm_name := d.Get("new_vm_name").(string)
	name := d.Get("name").(string)
	app_consistent := d.Get("app_consistent").(bool)

	vm, err := client.VirtualMachines.GetByName(name)
	if err != nil {
		return err
	}

	_, err = vm.Clone(new_vm_name, app_consistent)
	if err != nil {
		return err
	}

	d.SetId(new_vm_name)

	return resourceSimplivityVirtualMachineRead(d, meta)
}
