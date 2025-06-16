package main

import (
	"terraform-provider-nocloudiso/iso"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: iso.Provider,
	})
}

