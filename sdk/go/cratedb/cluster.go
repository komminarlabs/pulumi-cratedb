// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cratedb

import (
	"context"
	"reflect"

	"errors"
	"github.com/komminarlabs/pulumi-cratedb/sdk/go/cratedb/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates and manages a cluster.
type Cluster struct {
	pulumi.CustomResourceState

	// The allow custom storage flag.
	AllowCustomStorage pulumi.BoolOutput `pulumi:"allowCustomStorage"`
	// The allow suspend flag.
	AllowSuspend pulumi.BoolOutput `pulumi:"allowSuspend"`
	// The backup schedule.
	BackupSchedule pulumi.StringOutput `pulumi:"backupSchedule"`
	// The channel of the cluster. Default is `stable`.
	Channel pulumi.StringOutput `pulumi:"channel"`
	// The CrateDB version of the cluster.
	CrateVersion pulumi.StringOutput `pulumi:"crateVersion"`
	// The DublinCore of the cluster.
	Dc ClusterDcOutput `pulumi:"dc"`
	// The deletion protected flag.
	DeletionProtected pulumi.BoolOutput `pulumi:"deletionProtected"`
	// The external IP address.
	ExternalIp pulumi.StringOutput `pulumi:"externalIp"`
	// The Fully Qualified Domain Name.
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// The garbage collection available flag.
	GcAvailable pulumi.BoolOutput `pulumi:"gcAvailable"`
	// The hardware specs of the cluster.
	HardwareSpecs ClusterHardwareSpecsOutput `pulumi:"hardwareSpecs"`
	// The health of the cluster.
	Health ClusterHealthOutput `pulumi:"health"`
	// The IP whitelist of the cluster.
	IpWhitelists ClusterIpWhitelistArrayOutput `pulumi:"ipWhitelists"`
	// The name of the cluster.
	Name pulumi.StringOutput `pulumi:"name"`
	// The number of nodes in the cluster.
	NumNodes pulumi.Float64Output `pulumi:"numNodes"`
	// The organization id of the cluster.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// The origin of the cluster.
	Origin pulumi.StringOutput `pulumi:"origin"`
	// The password of the cluster.
	Password pulumi.StringOutput `pulumi:"password"`
	// The product name of the cluster.
	ProductName pulumi.StringOutput `pulumi:"productName"`
	// The product tier of the cluster.
	ProductTier pulumi.StringOutput `pulumi:"productTier"`
	// The product unit of the cluster. Default is `0`.
	ProductUnit pulumi.Float64Output `pulumi:"productUnit"`
	// The project id of the cluster.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The subscription id of the cluster.
	SubscriptionId pulumi.StringOutput `pulumi:"subscriptionId"`
	// The suspended flag.
	Suspended pulumi.BoolOutput `pulumi:"suspended"`
	// The URL of the cluster.
	Url pulumi.StringOutput `pulumi:"url"`
	// The username of the cluster.
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CrateVersion == nil {
		return nil, errors.New("invalid value for required argument 'CrateVersion'")
	}
	if args.OrganizationId == nil {
		return nil, errors.New("invalid value for required argument 'OrganizationId'")
	}
	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	if args.ProductName == nil {
		return nil, errors.New("invalid value for required argument 'ProductName'")
	}
	if args.ProductTier == nil {
		return nil, errors.New("invalid value for required argument 'ProductTier'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.SubscriptionId == nil {
		return nil, errors.New("invalid value for required argument 'SubscriptionId'")
	}
	if args.Username == nil {
		return nil, errors.New("invalid value for required argument 'Username'")
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Cluster
	err := ctx.RegisterResource("cratedb:index/cluster:Cluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCluster gets an existing Cluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterState, opts ...pulumi.ResourceOption) (*Cluster, error) {
	var resource Cluster
	err := ctx.ReadResource("cratedb:index/cluster:Cluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cluster resources.
type clusterState struct {
	// The allow custom storage flag.
	AllowCustomStorage *bool `pulumi:"allowCustomStorage"`
	// The allow suspend flag.
	AllowSuspend *bool `pulumi:"allowSuspend"`
	// The backup schedule.
	BackupSchedule *string `pulumi:"backupSchedule"`
	// The channel of the cluster. Default is `stable`.
	Channel *string `pulumi:"channel"`
	// The CrateDB version of the cluster.
	CrateVersion *string `pulumi:"crateVersion"`
	// The DublinCore of the cluster.
	Dc *ClusterDc `pulumi:"dc"`
	// The deletion protected flag.
	DeletionProtected *bool `pulumi:"deletionProtected"`
	// The external IP address.
	ExternalIp *string `pulumi:"externalIp"`
	// The Fully Qualified Domain Name.
	Fqdn *string `pulumi:"fqdn"`
	// The garbage collection available flag.
	GcAvailable *bool `pulumi:"gcAvailable"`
	// The hardware specs of the cluster.
	HardwareSpecs *ClusterHardwareSpecs `pulumi:"hardwareSpecs"`
	// The health of the cluster.
	Health *ClusterHealth `pulumi:"health"`
	// The IP whitelist of the cluster.
	IpWhitelists []ClusterIpWhitelist `pulumi:"ipWhitelists"`
	// The name of the cluster.
	Name *string `pulumi:"name"`
	// The number of nodes in the cluster.
	NumNodes *float64 `pulumi:"numNodes"`
	// The organization id of the cluster.
	OrganizationId *string `pulumi:"organizationId"`
	// The origin of the cluster.
	Origin *string `pulumi:"origin"`
	// The password of the cluster.
	Password *string `pulumi:"password"`
	// The product name of the cluster.
	ProductName *string `pulumi:"productName"`
	// The product tier of the cluster.
	ProductTier *string `pulumi:"productTier"`
	// The product unit of the cluster. Default is `0`.
	ProductUnit *float64 `pulumi:"productUnit"`
	// The project id of the cluster.
	ProjectId *string `pulumi:"projectId"`
	// The subscription id of the cluster.
	SubscriptionId *string `pulumi:"subscriptionId"`
	// The suspended flag.
	Suspended *bool `pulumi:"suspended"`
	// The URL of the cluster.
	Url *string `pulumi:"url"`
	// The username of the cluster.
	Username *string `pulumi:"username"`
}

type ClusterState struct {
	// The allow custom storage flag.
	AllowCustomStorage pulumi.BoolPtrInput
	// The allow suspend flag.
	AllowSuspend pulumi.BoolPtrInput
	// The backup schedule.
	BackupSchedule pulumi.StringPtrInput
	// The channel of the cluster. Default is `stable`.
	Channel pulumi.StringPtrInput
	// The CrateDB version of the cluster.
	CrateVersion pulumi.StringPtrInput
	// The DublinCore of the cluster.
	Dc ClusterDcPtrInput
	// The deletion protected flag.
	DeletionProtected pulumi.BoolPtrInput
	// The external IP address.
	ExternalIp pulumi.StringPtrInput
	// The Fully Qualified Domain Name.
	Fqdn pulumi.StringPtrInput
	// The garbage collection available flag.
	GcAvailable pulumi.BoolPtrInput
	// The hardware specs of the cluster.
	HardwareSpecs ClusterHardwareSpecsPtrInput
	// The health of the cluster.
	Health ClusterHealthPtrInput
	// The IP whitelist of the cluster.
	IpWhitelists ClusterIpWhitelistArrayInput
	// The name of the cluster.
	Name pulumi.StringPtrInput
	// The number of nodes in the cluster.
	NumNodes pulumi.Float64PtrInput
	// The organization id of the cluster.
	OrganizationId pulumi.StringPtrInput
	// The origin of the cluster.
	Origin pulumi.StringPtrInput
	// The password of the cluster.
	Password pulumi.StringPtrInput
	// The product name of the cluster.
	ProductName pulumi.StringPtrInput
	// The product tier of the cluster.
	ProductTier pulumi.StringPtrInput
	// The product unit of the cluster. Default is `0`.
	ProductUnit pulumi.Float64PtrInput
	// The project id of the cluster.
	ProjectId pulumi.StringPtrInput
	// The subscription id of the cluster.
	SubscriptionId pulumi.StringPtrInput
	// The suspended flag.
	Suspended pulumi.BoolPtrInput
	// The URL of the cluster.
	Url pulumi.StringPtrInput
	// The username of the cluster.
	Username pulumi.StringPtrInput
}

func (ClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterState)(nil)).Elem()
}

type clusterArgs struct {
	// The channel of the cluster. Default is `stable`.
	Channel *string `pulumi:"channel"`
	// The CrateDB version of the cluster.
	CrateVersion string `pulumi:"crateVersion"`
	// The hardware specs of the cluster.
	HardwareSpecs *ClusterHardwareSpecs `pulumi:"hardwareSpecs"`
	// The name of the cluster.
	Name *string `pulumi:"name"`
	// The organization id of the cluster.
	OrganizationId string `pulumi:"organizationId"`
	// The password of the cluster.
	Password string `pulumi:"password"`
	// The product name of the cluster.
	ProductName string `pulumi:"productName"`
	// The product tier of the cluster.
	ProductTier string `pulumi:"productTier"`
	// The product unit of the cluster. Default is `0`.
	ProductUnit *float64 `pulumi:"productUnit"`
	// The project id of the cluster.
	ProjectId string `pulumi:"projectId"`
	// The subscription id of the cluster.
	SubscriptionId string `pulumi:"subscriptionId"`
	// The username of the cluster.
	Username string `pulumi:"username"`
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	// The channel of the cluster. Default is `stable`.
	Channel pulumi.StringPtrInput
	// The CrateDB version of the cluster.
	CrateVersion pulumi.StringInput
	// The hardware specs of the cluster.
	HardwareSpecs ClusterHardwareSpecsPtrInput
	// The name of the cluster.
	Name pulumi.StringPtrInput
	// The organization id of the cluster.
	OrganizationId pulumi.StringInput
	// The password of the cluster.
	Password pulumi.StringInput
	// The product name of the cluster.
	ProductName pulumi.StringInput
	// The product tier of the cluster.
	ProductTier pulumi.StringInput
	// The product unit of the cluster. Default is `0`.
	ProductUnit pulumi.Float64PtrInput
	// The project id of the cluster.
	ProjectId pulumi.StringInput
	// The subscription id of the cluster.
	SubscriptionId pulumi.StringInput
	// The username of the cluster.
	Username pulumi.StringInput
}

func (ClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterArgs)(nil)).Elem()
}

type ClusterInput interface {
	pulumi.Input

	ToClusterOutput() ClusterOutput
	ToClusterOutputWithContext(ctx context.Context) ClusterOutput
}

func (*Cluster) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil)).Elem()
}

func (i *Cluster) ToClusterOutput() ClusterOutput {
	return i.ToClusterOutputWithContext(context.Background())
}

func (i *Cluster) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterOutput)
}

// ClusterArrayInput is an input type that accepts ClusterArray and ClusterArrayOutput values.
// You can construct a concrete instance of `ClusterArrayInput` via:
//
//	ClusterArray{ ClusterArgs{...} }
type ClusterArrayInput interface {
	pulumi.Input

	ToClusterArrayOutput() ClusterArrayOutput
	ToClusterArrayOutputWithContext(context.Context) ClusterArrayOutput
}

type ClusterArray []ClusterInput

func (ClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Cluster)(nil)).Elem()
}

func (i ClusterArray) ToClusterArrayOutput() ClusterArrayOutput {
	return i.ToClusterArrayOutputWithContext(context.Background())
}

func (i ClusterArray) ToClusterArrayOutputWithContext(ctx context.Context) ClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterArrayOutput)
}

// ClusterMapInput is an input type that accepts ClusterMap and ClusterMapOutput values.
// You can construct a concrete instance of `ClusterMapInput` via:
//
//	ClusterMap{ "key": ClusterArgs{...} }
type ClusterMapInput interface {
	pulumi.Input

	ToClusterMapOutput() ClusterMapOutput
	ToClusterMapOutputWithContext(context.Context) ClusterMapOutput
}

type ClusterMap map[string]ClusterInput

func (ClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Cluster)(nil)).Elem()
}

func (i ClusterMap) ToClusterMapOutput() ClusterMapOutput {
	return i.ToClusterMapOutputWithContext(context.Background())
}

func (i ClusterMap) ToClusterMapOutputWithContext(ctx context.Context) ClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterMapOutput)
}

type ClusterOutput struct{ *pulumi.OutputState }

func (ClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil)).Elem()
}

func (o ClusterOutput) ToClusterOutput() ClusterOutput {
	return o
}

func (o ClusterOutput) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return o
}

// The allow custom storage flag.
func (o ClusterOutput) AllowCustomStorage() pulumi.BoolOutput {
	return o.ApplyT(func(v *Cluster) pulumi.BoolOutput { return v.AllowCustomStorage }).(pulumi.BoolOutput)
}

// The allow suspend flag.
func (o ClusterOutput) AllowSuspend() pulumi.BoolOutput {
	return o.ApplyT(func(v *Cluster) pulumi.BoolOutput { return v.AllowSuspend }).(pulumi.BoolOutput)
}

// The backup schedule.
func (o ClusterOutput) BackupSchedule() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.BackupSchedule }).(pulumi.StringOutput)
}

// The channel of the cluster. Default is `stable`.
func (o ClusterOutput) Channel() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Channel }).(pulumi.StringOutput)
}

// The CrateDB version of the cluster.
func (o ClusterOutput) CrateVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.CrateVersion }).(pulumi.StringOutput)
}

// The DublinCore of the cluster.
func (o ClusterOutput) Dc() ClusterDcOutput {
	return o.ApplyT(func(v *Cluster) ClusterDcOutput { return v.Dc }).(ClusterDcOutput)
}

// The deletion protected flag.
func (o ClusterOutput) DeletionProtected() pulumi.BoolOutput {
	return o.ApplyT(func(v *Cluster) pulumi.BoolOutput { return v.DeletionProtected }).(pulumi.BoolOutput)
}

// The external IP address.
func (o ClusterOutput) ExternalIp() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.ExternalIp }).(pulumi.StringOutput)
}

// The Fully Qualified Domain Name.
func (o ClusterOutput) Fqdn() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Fqdn }).(pulumi.StringOutput)
}

// The garbage collection available flag.
func (o ClusterOutput) GcAvailable() pulumi.BoolOutput {
	return o.ApplyT(func(v *Cluster) pulumi.BoolOutput { return v.GcAvailable }).(pulumi.BoolOutput)
}

// The hardware specs of the cluster.
func (o ClusterOutput) HardwareSpecs() ClusterHardwareSpecsOutput {
	return o.ApplyT(func(v *Cluster) ClusterHardwareSpecsOutput { return v.HardwareSpecs }).(ClusterHardwareSpecsOutput)
}

// The health of the cluster.
func (o ClusterOutput) Health() ClusterHealthOutput {
	return o.ApplyT(func(v *Cluster) ClusterHealthOutput { return v.Health }).(ClusterHealthOutput)
}

// The IP whitelist of the cluster.
func (o ClusterOutput) IpWhitelists() ClusterIpWhitelistArrayOutput {
	return o.ApplyT(func(v *Cluster) ClusterIpWhitelistArrayOutput { return v.IpWhitelists }).(ClusterIpWhitelistArrayOutput)
}

// The name of the cluster.
func (o ClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The number of nodes in the cluster.
func (o ClusterOutput) NumNodes() pulumi.Float64Output {
	return o.ApplyT(func(v *Cluster) pulumi.Float64Output { return v.NumNodes }).(pulumi.Float64Output)
}

// The organization id of the cluster.
func (o ClusterOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// The origin of the cluster.
func (o ClusterOutput) Origin() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Origin }).(pulumi.StringOutput)
}

// The password of the cluster.
func (o ClusterOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

// The product name of the cluster.
func (o ClusterOutput) ProductName() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.ProductName }).(pulumi.StringOutput)
}

// The product tier of the cluster.
func (o ClusterOutput) ProductTier() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.ProductTier }).(pulumi.StringOutput)
}

// The product unit of the cluster. Default is `0`.
func (o ClusterOutput) ProductUnit() pulumi.Float64Output {
	return o.ApplyT(func(v *Cluster) pulumi.Float64Output { return v.ProductUnit }).(pulumi.Float64Output)
}

// The project id of the cluster.
func (o ClusterOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The subscription id of the cluster.
func (o ClusterOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.SubscriptionId }).(pulumi.StringOutput)
}

// The suspended flag.
func (o ClusterOutput) Suspended() pulumi.BoolOutput {
	return o.ApplyT(func(v *Cluster) pulumi.BoolOutput { return v.Suspended }).(pulumi.BoolOutput)
}

// The URL of the cluster.
func (o ClusterOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

// The username of the cluster.
func (o ClusterOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

type ClusterArrayOutput struct{ *pulumi.OutputState }

func (ClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Cluster)(nil)).Elem()
}

func (o ClusterArrayOutput) ToClusterArrayOutput() ClusterArrayOutput {
	return o
}

func (o ClusterArrayOutput) ToClusterArrayOutputWithContext(ctx context.Context) ClusterArrayOutput {
	return o
}

func (o ClusterArrayOutput) Index(i pulumi.IntInput) ClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Cluster {
		return vs[0].([]*Cluster)[vs[1].(int)]
	}).(ClusterOutput)
}

type ClusterMapOutput struct{ *pulumi.OutputState }

func (ClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Cluster)(nil)).Elem()
}

func (o ClusterMapOutput) ToClusterMapOutput() ClusterMapOutput {
	return o
}

func (o ClusterMapOutput) ToClusterMapOutputWithContext(ctx context.Context) ClusterMapOutput {
	return o
}

func (o ClusterMapOutput) MapIndex(k pulumi.StringInput) ClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Cluster {
		return vs[0].(map[string]*Cluster)[vs[1].(string)]
	}).(ClusterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterInput)(nil)).Elem(), &Cluster{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterArrayInput)(nil)).Elem(), ClusterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterMapInput)(nil)).Elem(), ClusterMap{})
	pulumi.RegisterOutputType(ClusterOutput{})
	pulumi.RegisterOutputType(ClusterArrayOutput{})
	pulumi.RegisterOutputType(ClusterMapOutput{})
}