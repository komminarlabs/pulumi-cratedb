package cratedb

import (
	// Allow embedding bridge-metadata.json in the provider.
	_ "embed"
	"fmt"
	"path"

	// Import custom shim
	"github.com/komminarlabs/pulumi-cratedb/provider/pkg/version"
	cratedbshim "github.com/komminarlabs/terraform-provider-cratedb/shim"

	pfbridge "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "cratedb"
	// modules:
	mainMod = "index" // the cratedb module
)

//go:embed cmd/pulumi-resource-cratedb/bridge-metadata.json
var bridgeMetadata []byte

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	info := tfbridge.ProviderInfo{
		// Instantiate the Terraform provider
		P:                 pfbridge.ShimProvider(cratedbshim.NewProvider(version.Version)),
		Name:              "cratedb",
		DisplayName:       "CrateDB",
		Publisher:         "komminarlabs",
		Version:           version.Version,
		LogoURL:           "https://avatars.githubusercontent.com/u/4048232?s=48&v=4",
		PluginDownloadURL: "github://api.github.com/komminarlabs",
		Description:       "A Pulumi package for creating and managing CrateDB resources.",
		Keywords:          []string{"pulumi", "cratedb", "category/database"},
		License:           "Apache-2.0",
		Homepage:          "https://www.cratedb.com",
		Repository:        "https://github.com/komminarlabs/pulumi-cratedb",
		GitHubOrg:         "komminarlabs",
		MetadataInfo:      tfbridge.NewProviderMetadata(bridgeMetadata),
		Resources: map[string]*tfbridge.ResourceInfo{
			"cratedb_cluster":      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Cluster")},
			"cratedb_organization": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Organization")},
			"cratedb_project":      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Project")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"cratedb_cluster":       {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getCluster")},
			"cratedb_organization":  {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getOrganization")},
			"cratedb_organizations": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getOrganizations")},
			"cratedb_project":       {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getProject")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: "@komminarlabs/cratedb",
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0",
				"@types/mime": "^2.0.0",
			},
		},
		Python: &tfbridge.PythonInfo{
			PackageName: "komminarlabs_cratedb",
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: path.Join(
				fmt.Sprintf("github.com/komminarlabs/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RootNamespace: "KomminarLabs",
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
			Namespaces: map[string]string{
				"cratedb": "CrateDB",
			},
		},
	}

	info.MustComputeTokens(tokens.SingleModule("cratedb_", mainMod,
		tokens.MakeStandard(mainPkg)))
	info.MustApplyAutoAliases()
	info.SetAutonaming(255, "-")

	return info
}
