package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/robsonvn/terraform-provider-couchdb/couchdb"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: couchdb.Provider})
}
