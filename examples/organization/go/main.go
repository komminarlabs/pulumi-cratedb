package main

import (
	"github.com/komminarlabs/pulumi-cratedb/sdk/go/cratedb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		defaultOrg, err := cratedb.NewOrganization(ctx, "default", &cratedb.OrganizationArgs{
			Name: pulumi.String("default"),
		})
		if err != nil {
			return err
		}

		ctx.Export("defaultOrgName", defaultOrg.Name)
		return nil
	})
}
