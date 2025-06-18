package main

import (
	"terraform-provider-nocloudiso/nocloudiso"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: nocloudiso.Provider,
	})
}

