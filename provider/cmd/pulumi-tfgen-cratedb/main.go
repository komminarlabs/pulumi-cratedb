package main

import (
	"github.com/pulumi/pulumi-terraform-bridge/pf/tfgen"

	cratedb "github.com/komminarlabs/pulumi-cratedb/provider"
)

func main() {
	tfgen.Main("cratedb", cratedb.Provider())
}
