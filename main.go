package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/sgarlick987/terraform-provider-kong/kong"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: kong.Provider,
	})
}
