package main

import (
	cratedb "github.com/komminarlabs/pulumi-cratedb/provider"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfgen"
)

func main() {
	tfgen.Main("cratedb", cratedb.Provider())
}
