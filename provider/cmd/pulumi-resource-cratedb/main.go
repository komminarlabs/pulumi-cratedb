//go:generate go run ./generate.go

package main

import (
	"context"
	_ "embed"

	cratedb "github.com/komminarlabs/pulumi-cratedb/provider"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfbridge"
)

//go:embed schema.json
var pulumiSchema []byte

func main() {
	tfbridge.Main(context.Background(), "cratedb", cratedb.Provider(),
		tfbridge.ProviderMetadata{PackageSchema: pulumiSchema})
}
