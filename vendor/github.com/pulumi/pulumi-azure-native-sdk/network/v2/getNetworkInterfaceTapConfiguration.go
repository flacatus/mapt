// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the specified tap configuration on a network interface.
//
// Uses Azure REST API version 2023-02-01.
//
// Other available API versions: 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-05-01.
func LookupNetworkInterfaceTapConfiguration(ctx *pulumi.Context, args *LookupNetworkInterfaceTapConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupNetworkInterfaceTapConfigurationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupNetworkInterfaceTapConfigurationResult
	err := ctx.Invoke("azure-native:network:getNetworkInterfaceTapConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupNetworkInterfaceTapConfigurationArgs struct {
	// The name of the network interface.
	NetworkInterfaceName string `pulumi:"networkInterfaceName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the tap configuration.
	TapConfigurationName string `pulumi:"tapConfigurationName"`
}

// Tap configuration in a Network Interface.
type LookupNetworkInterfaceTapConfigurationResult struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// The provisioning state of the network interface tap configuration resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Sub Resource type.
	Type string `pulumi:"type"`
	// The reference to the Virtual Network Tap resource.
	VirtualNetworkTap *VirtualNetworkTapResponse `pulumi:"virtualNetworkTap"`
}

// Defaults sets the appropriate defaults for LookupNetworkInterfaceTapConfigurationResult
func (val *LookupNetworkInterfaceTapConfigurationResult) Defaults() *LookupNetworkInterfaceTapConfigurationResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.VirtualNetworkTap = tmp.VirtualNetworkTap.Defaults()

	return &tmp
}
func LookupNetworkInterfaceTapConfigurationOutput(ctx *pulumi.Context, args LookupNetworkInterfaceTapConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupNetworkInterfaceTapConfigurationResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupNetworkInterfaceTapConfigurationResultOutput, error) {
			args := v.(LookupNetworkInterfaceTapConfigurationArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:network:getNetworkInterfaceTapConfiguration", args, LookupNetworkInterfaceTapConfigurationResultOutput{}, options).(LookupNetworkInterfaceTapConfigurationResultOutput), nil
		}).(LookupNetworkInterfaceTapConfigurationResultOutput)
}

type LookupNetworkInterfaceTapConfigurationOutputArgs struct {
	// The name of the network interface.
	NetworkInterfaceName pulumi.StringInput `pulumi:"networkInterfaceName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the tap configuration.
	TapConfigurationName pulumi.StringInput `pulumi:"tapConfigurationName"`
}

func (LookupNetworkInterfaceTapConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkInterfaceTapConfigurationArgs)(nil)).Elem()
}

// Tap configuration in a Network Interface.
type LookupNetworkInterfaceTapConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupNetworkInterfaceTapConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkInterfaceTapConfigurationResult)(nil)).Elem()
}

func (o LookupNetworkInterfaceTapConfigurationResultOutput) ToLookupNetworkInterfaceTapConfigurationResultOutput() LookupNetworkInterfaceTapConfigurationResultOutput {
	return o
}

func (o LookupNetworkInterfaceTapConfigurationResultOutput) ToLookupNetworkInterfaceTapConfigurationResultOutputWithContext(ctx context.Context) LookupNetworkInterfaceTapConfigurationResultOutput {
	return o
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupNetworkInterfaceTapConfigurationResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceTapConfigurationResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupNetworkInterfaceTapConfigurationResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceTapConfigurationResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The name of the resource that is unique within a resource group. This name can be used to access the resource.
func (o LookupNetworkInterfaceTapConfigurationResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceTapConfigurationResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The provisioning state of the network interface tap configuration resource.
func (o LookupNetworkInterfaceTapConfigurationResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceTapConfigurationResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Sub Resource type.
func (o LookupNetworkInterfaceTapConfigurationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceTapConfigurationResult) string { return v.Type }).(pulumi.StringOutput)
}

// The reference to the Virtual Network Tap resource.
func (o LookupNetworkInterfaceTapConfigurationResultOutput) VirtualNetworkTap() VirtualNetworkTapResponsePtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceTapConfigurationResult) *VirtualNetworkTapResponse {
		return v.VirtualNetworkTap
	}).(VirtualNetworkTapResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkInterfaceTapConfigurationResultOutput{})
}
