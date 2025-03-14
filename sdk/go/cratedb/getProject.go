// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cratedb

import (
	"context"
	"reflect"

	"github.com/komminarlabs/pulumi-cratedb/sdk/go/cratedb/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// To retrieve a project.
func LookupProject(ctx *pulumi.Context, args *LookupProjectArgs, opts ...pulumi.InvokeOption) (*LookupProjectResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupProjectResult
	err := ctx.Invoke("cratedb:index/getProject:getProject", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProject.
type LookupProjectArgs struct {
	// The id of the project.
	Id string `pulumi:"id"`
}

// A collection of values returned by getProject.
type LookupProjectResult struct {
	// The DublinCore of the project.
	Dc GetProjectDc `pulumi:"dc"`
	// The id of the project.
	Id string `pulumi:"id"`
	// The name of the project.
	Name string `pulumi:"name"`
	// The organization id of the project.
	OrganizationId string `pulumi:"organizationId"`
	// The region of the project.
	Region string `pulumi:"region"`
}

func LookupProjectOutput(ctx *pulumi.Context, args LookupProjectOutputArgs, opts ...pulumi.InvokeOption) LookupProjectResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupProjectResultOutput, error) {
			args := v.(LookupProjectArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("cratedb:index/getProject:getProject", args, LookupProjectResultOutput{}, options).(LookupProjectResultOutput), nil
		}).(LookupProjectResultOutput)
}

// A collection of arguments for invoking getProject.
type LookupProjectOutputArgs struct {
	// The id of the project.
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupProjectOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectArgs)(nil)).Elem()
}

// A collection of values returned by getProject.
type LookupProjectResultOutput struct{ *pulumi.OutputState }

func (LookupProjectResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectResult)(nil)).Elem()
}

func (o LookupProjectResultOutput) ToLookupProjectResultOutput() LookupProjectResultOutput {
	return o
}

func (o LookupProjectResultOutput) ToLookupProjectResultOutputWithContext(ctx context.Context) LookupProjectResultOutput {
	return o
}

// The DublinCore of the project.
func (o LookupProjectResultOutput) Dc() GetProjectDcOutput {
	return o.ApplyT(func(v LookupProjectResult) GetProjectDc { return v.Dc }).(GetProjectDcOutput)
}

// The id of the project.
func (o LookupProjectResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the project.
func (o LookupProjectResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Name }).(pulumi.StringOutput)
}

// The organization id of the project.
func (o LookupProjectResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

// The region of the project.
func (o LookupProjectResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Region }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProjectResultOutput{})
}
