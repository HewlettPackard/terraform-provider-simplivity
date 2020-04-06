package simplivity

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceSimplivityVirtualMachine() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSimplivityVirtualMachineRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceSimplivityVirtualMachineRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Config).Client
	name := d.Get("name").(string)

	log.Printf("[DEBUG] Reading Virtual Machine: %s", name)
	vm, err := client.VirtualMachines.GetByName(name)
	if err != nil {
		return nil
	}

	d.SetId(name)
	d.Set("name", vm.Name)

	return nil
}
