package simplivity

import (
	"log"

	"github.com/HewlettPackard/simplivity-go/ovc"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSimplivityVirtualMachineBackup() *schema.Resource {
	return &schema.Resource{
		Create: resourceSimplivityVirtualMachineBackupCreate,
		Read:   resourceSimplivityVirtualMachineBackupRead,
		Delete: resourceSimplivityVirtualMachineBackupDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"vm_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"omnistack_cluster_name": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"retention": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
				ForceNew: true,
			},
			"app_consistent": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				ForceNew: true,
			},
			"consistency_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "NONE",
				ForceNew: true,
			},
		},
	}
}

func resourceSimplivityVirtualMachineBackupCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Config).Client

	name := d.Get("name").(string)
	vm_name := d.Get("vm_name").(string)
	omnistack_cluster_name := d.Get("omnistack_cluster_name").(string)

	cluster, err := client.OmniStackClusters.GetByName(omnistack_cluster_name)
	if err != nil {
		return err
	}

	vm, err := client.VirtualMachines.GetByName(vm_name)
	if err != nil {
		return err
	}

	backupRequest := &ovc.CreateBackupRequest{
		Name:            name,
		ConsistencyType: d.Get("consistency_type").(string),
		AppConsistent:   d.Get("app_consistent").(bool),
		Retention:       d.Get("retention").(int),
		Destination:     cluster.Id}

	_, err = vm.CreateBackup(backupRequest, nil)
	if err != nil {
		return err
	}

	d.SetId(name)

	return resourceSimplivityVirtualMachineBackupRead(d, meta)
}

func resourceSimplivityVirtualMachineBackupRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Config).Client
	name := d.Id()

	log.Printf("[DEBUG] Reading backup details of a Virtual Machine: %s", name)
	backup, err := client.Backups.GetByName(name)
	if err != nil {
		d.SetId("")
		return err
	}

	d.SetId(name)
	d.Set("name", name)
	d.Set("app_consistent", backup.ApplicationConsistent)
	d.Set("cosistency_type", backup.ConsistencyType)
	d.Set("omnistack_cluster_name", backup.OmnistackClusterName)

	return nil
}

func resourceSimplivityVirtualMachineBackupDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Config).Client
	name := d.Id()

	log.Printf("[DEBUG] Reading backup details of a Virtual Machine: %s", name)
	backup, err := client.Backups.GetByName(name)
	if err != nil {
		return err
	}

	err = backup.Delete()
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
