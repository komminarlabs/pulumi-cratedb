// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cratedb

import (
	"context"
	"reflect"

	"github.com/komminarlabs/pulumi-cratedb/sdk/go/cratedb/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// To retrieve all organizations.
func GetOrganizations(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetOrganizationsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetOrganizationsResult
	err := ctx.Invoke("cratedb:index/getOrganizations:getOrganizations", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getOrganizations.
type GetOrganizationsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id            string                         `pulumi:"id"`
	Organizations []GetOrganizationsOrganization `pulumi:"organizations"`
}

func GetOrganizationsOutput(ctx *pulumi.Context, opts ...pulumi.InvokeOption) GetOrganizationsResultOutput {
	return pulumi.ToOutput(0).ApplyT(func(int) (GetOrganizationsResult, error) {
		r, err := GetOrganizations(ctx, opts...)
		var s GetOrganizationsResult
		if r != nil {
			s = *r
		}
		return s, err
	}).(GetOrganizationsResultOutput)
}

// A collection of values returned by getOrganizations.
type GetOrganizationsResultOutput struct{ *pulumi.OutputState }

func (GetOrganizationsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOrganizationsResult)(nil)).Elem()
}

func (o GetOrganizationsResultOutput) ToGetOrganizationsResultOutput() GetOrganizationsResultOutput {
	return o
}

func (o GetOrganizationsResultOutput) ToGetOrganizationsResultOutputWithContext(ctx context.Context) GetOrganizationsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetOrganizationsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetOrganizationsResultOutput) Organizations() GetOrganizationsOrganizationArrayOutput {
	return o.ApplyT(func(v GetOrganizationsResult) []GetOrganizationsOrganization { return v.Organizations }).(GetOrganizationsOrganizationArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetOrganizationsResultOutput{})
}
