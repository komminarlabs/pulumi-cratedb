// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cratedb

import (
	"context"
	"reflect"

	"github.com/komminarlabs/pulumi-cratedb/sdk/go/cratedb/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// To retrieve an organization.
func LookupOrganization(ctx *pulumi.Context, args *LookupOrganizationArgs, opts ...pulumi.InvokeOption) (*LookupOrganizationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupOrganizationResult
	err := ctx.Invoke("cratedb:index/getOrganization:getOrganization", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOrganization.
type LookupOrganizationArgs struct {
	// The id of the organization.
	Id string `pulumi:"id"`
}

// A collection of values returned by getOrganization.
type LookupOrganizationResult struct {
	// The DublinCore of the organization.
	Dc GetOrganizationDc `pulumi:"dc"`
	// The notification email used in the organization.
	Email string `pulumi:"email"`
	// The id of the organization.
	Id string `pulumi:"id"`
	// The name of the organization.
	Name string `pulumi:"name"`
	// Whether notifications enabled for the organization.
	NotificationsEnabled bool `pulumi:"notificationsEnabled"`
	// The support plan type used in the organization.
	PlanType int `pulumi:"planType"`
	// The project count in the organization.
	ProjectCount int `pulumi:"projectCount"`
	// The role FQN.
	RoleFqn string `pulumi:"roleFqn"`
}

func LookupOrganizationOutput(ctx *pulumi.Context, args LookupOrganizationOutputArgs, opts ...pulumi.InvokeOption) LookupOrganizationResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupOrganizationResultOutput, error) {
			args := v.(LookupOrganizationArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("cratedb:index/getOrganization:getOrganization", args, LookupOrganizationResultOutput{}, options).(LookupOrganizationResultOutput), nil
		}).(LookupOrganizationResultOutput)
}

// A collection of arguments for invoking getOrganization.
type LookupOrganizationOutputArgs struct {
	// The id of the organization.
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupOrganizationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrganizationArgs)(nil)).Elem()
}

// A collection of values returned by getOrganization.
type LookupOrganizationResultOutput struct{ *pulumi.OutputState }

func (LookupOrganizationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrganizationResult)(nil)).Elem()
}

func (o LookupOrganizationResultOutput) ToLookupOrganizationResultOutput() LookupOrganizationResultOutput {
	return o
}

func (o LookupOrganizationResultOutput) ToLookupOrganizationResultOutputWithContext(ctx context.Context) LookupOrganizationResultOutput {
	return o
}

// The DublinCore of the organization.
func (o LookupOrganizationResultOutput) Dc() GetOrganizationDcOutput {
	return o.ApplyT(func(v LookupOrganizationResult) GetOrganizationDc { return v.Dc }).(GetOrganizationDcOutput)
}

// The notification email used in the organization.
func (o LookupOrganizationResultOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationResult) string { return v.Email }).(pulumi.StringOutput)
}

// The id of the organization.
func (o LookupOrganizationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the organization.
func (o LookupOrganizationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationResult) string { return v.Name }).(pulumi.StringOutput)
}

// Whether notifications enabled for the organization.
func (o LookupOrganizationResultOutput) NotificationsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupOrganizationResult) bool { return v.NotificationsEnabled }).(pulumi.BoolOutput)
}

// The support plan type used in the organization.
func (o LookupOrganizationResultOutput) PlanType() pulumi.IntOutput {
	return o.ApplyT(func(v LookupOrganizationResult) int { return v.PlanType }).(pulumi.IntOutput)
}

// The project count in the organization.
func (o LookupOrganizationResultOutput) ProjectCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupOrganizationResult) int { return v.ProjectCount }).(pulumi.IntOutput)
}

// The role FQN.
func (o LookupOrganizationResultOutput) RoleFqn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationResult) string { return v.RoleFqn }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOrganizationResultOutput{})
}
