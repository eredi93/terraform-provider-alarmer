package main

import (
	"github.com/hashicorp/terraform/plugin"

	"github.com/eredi93/terraform-provider-alarmer/alarmer"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: alarmer.Provider,
	})
}
