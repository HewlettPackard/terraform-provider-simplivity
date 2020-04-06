package simplivity

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceSimplivityVirtualMachineBackup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSimplivityVirtualMachineBackupRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceSimplivityVirtualMachineBackupRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Config).Client
	name := d.Get("name").(string)

	log.Printf("[DEBUG] Reading backup: %s", name)
	vm, err := client.Backups.GetByName(name)
	if err != nil {
		return nil
	}

	d.SetId(name)
	d.Set("name", vm.Name)

	return nil
}
