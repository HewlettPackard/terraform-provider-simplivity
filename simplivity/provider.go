package simplivity

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func Provider() terraform.ResourceProvider {
	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"ovc_ip": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("OVC_TOKEN", nil),
				Description: "IP of the Omnistack Virtual Controller",
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("VCENTER_USERNAME", nil),
				Description: "Vcenter username",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("VCENTER_PASSWORD", nil),
				Description: "Vcenter password",
			},
			"certificate_path": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CERTIFICATE_PATH", nil),
				Description: "SSL certificate path",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"simplivity_vm":        resourceSimplivityVirtualMachine(),
			"simplivity_vm_clone":  resourceSimplivityVirtualMachineClone(),
			"simplivity_vm_move":   resourceSimplivityVirtualMachineMove(),
			"simplivity_vm_backup": resourceSimplivityVirtualMachineBackup()},

		DataSourcesMap: map[string]*schema.Resource{},
	}

	p.ConfigureFunc = providerConfigure(p)

	return p
}

func providerConfigure(p *schema.Provider) schema.ConfigureFunc {
	return func(d *schema.ResourceData) (interface{}, error) {
		config := Config{
			OVCIP:           d.Get("ovc_ip").(string),
			Username:        d.Get("username").(string),
			Password:        d.Get("password").(string),
			CertificatePath: d.Get("certificate_path").(string),
		}

		err := config.SetClient()
		if err != nil {
			return nil, err
		}

		return &config, nil
	}
}
