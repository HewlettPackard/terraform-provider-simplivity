package main

import (
	"github.com/HewlettPackard/terraform-provider-simplivity/simplivity"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: simplivity.Provider})
}
