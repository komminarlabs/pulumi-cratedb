package cratedb

import (
	"fmt"
	"path"

	// Allow embedding bridge-metadata.json in the provider.
	_ "embed"

	cratedbshim "github.com/komminarlabs/terraform-provider-cratedb/shim"

	pf "github.com/pulumi/pulumi-terraform-bridge/pf/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"

	// Import custom shim
	"github.com/komminarlabs/pulumi-cratedb/provider/pkg/version"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "cratedb"
	// modules:
	mainMod = "index" // the cratedb module
)

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(resource.PropertyMap, shim.ResourceConfig) error {
	return nil
}

//go:embed cmd/pulumi-resource-cratedb/bridge-metadata.json
var metadata []byte

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	prov := tfbridge.ProviderInfo{
		// Instantiate the Terraform provider
		P:                    pf.ShimProvider(cratedbshim.NewProvider()),
		Name:                 "cratedb",
		DisplayName:          "CrateDB",
		Publisher:            "komminarlabs",
		Version:              version.Version,
		LogoURL:              "https://avatars.githubusercontent.com/u/4048232?s=48&v=4",
		PluginDownloadURL:    "github://api.github.com/komminarlabs",
		Description:          "A Pulumi package for creating and managing CrateDB resources.",
		Keywords:             []string{"pulumi", "cratedb", "category/database"},
		License:              "Apache-2.0",
		Homepage:             "https://www.cratedb.com",
		Repository:           "https://github.com/komminarlabs/pulumi-cratedb",
		GitHubOrg:            "komminarlabs",
		MetadataInfo:         tfbridge.NewProviderMetadata(metadata),
		PreConfigureCallback: preConfigureCallback,
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

	prov.MustComputeTokens(tokens.SingleModule("cratedb_", mainMod,
		tokens.MakeStandard(mainPkg)))
	prov.MustApplyAutoAliases()
	prov.SetAutonaming(255, "-")

	return prov
}
