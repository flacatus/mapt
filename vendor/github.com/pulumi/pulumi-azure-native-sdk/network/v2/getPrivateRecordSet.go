// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a record set.
//
// Uses Azure REST API version 2020-06-01.
//
// Other available API versions: 2024-06-01.
func LookupPrivateRecordSet(ctx *pulumi.Context, args *LookupPrivateRecordSetArgs, opts ...pulumi.InvokeOption) (*LookupPrivateRecordSetResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupPrivateRecordSetResult
	err := ctx.Invoke("azure-native:network:getPrivateRecordSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateRecordSetArgs struct {
	// The name of the Private DNS zone (without a terminating dot).
	PrivateZoneName string `pulumi:"privateZoneName"`
	// The type of DNS record in this record set.
	RecordType string `pulumi:"recordType"`
	// The name of the record set, relative to the name of the zone.
	RelativeRecordSetName string `pulumi:"relativeRecordSetName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Describes a DNS record set (a collection of DNS records with the same name and type) in a Private DNS zone.
type LookupPrivateRecordSetResult struct {
	// The list of A records in the record set.
	ARecords []ARecordResponse `pulumi:"aRecords"`
	// The list of AAAA records in the record set.
	AaaaRecords []AaaaRecordResponse `pulumi:"aaaaRecords"`
	// The CNAME record in the record set.
	CnameRecord *CnameRecordResponse `pulumi:"cnameRecord"`
	// The ETag of the record set.
	Etag *string `pulumi:"etag"`
	// Fully qualified domain name of the record set.
	Fqdn string `pulumi:"fqdn"`
	// Fully qualified resource Id for the resource. Example - '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateDnsZoneName}'.
	Id string `pulumi:"id"`
	// Is the record set auto-registered in the Private DNS zone through a virtual network link?
	IsAutoRegistered bool `pulumi:"isAutoRegistered"`
	// The metadata attached to the record set.
	Metadata map[string]string `pulumi:"metadata"`
	// The list of MX records in the record set.
	MxRecords []MxRecordResponse `pulumi:"mxRecords"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The list of PTR records in the record set.
	PtrRecords []PtrRecordResponse `pulumi:"ptrRecords"`
	// The SOA record in the record set.
	SoaRecord *SoaRecordResponse `pulumi:"soaRecord"`
	// The list of SRV records in the record set.
	SrvRecords []SrvRecordResponse `pulumi:"srvRecords"`
	// The TTL (time-to-live) of the records in the record set.
	Ttl *float64 `pulumi:"ttl"`
	// The list of TXT records in the record set.
	TxtRecords []TxtRecordResponse `pulumi:"txtRecords"`
	// The type of the resource. Example - 'Microsoft.Network/privateDnsZones'.
	Type string `pulumi:"type"`
}

func LookupPrivateRecordSetOutput(ctx *pulumi.Context, args LookupPrivateRecordSetOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateRecordSetResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupPrivateRecordSetResultOutput, error) {
			args := v.(LookupPrivateRecordSetArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:network:getPrivateRecordSet", args, LookupPrivateRecordSetResultOutput{}, options).(LookupPrivateRecordSetResultOutput), nil
		}).(LookupPrivateRecordSetResultOutput)
}

type LookupPrivateRecordSetOutputArgs struct {
	// The name of the Private DNS zone (without a terminating dot).
	PrivateZoneName pulumi.StringInput `pulumi:"privateZoneName"`
	// The type of DNS record in this record set.
	RecordType pulumi.StringInput `pulumi:"recordType"`
	// The name of the record set, relative to the name of the zone.
	RelativeRecordSetName pulumi.StringInput `pulumi:"relativeRecordSetName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPrivateRecordSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateRecordSetArgs)(nil)).Elem()
}

// Describes a DNS record set (a collection of DNS records with the same name and type) in a Private DNS zone.
type LookupPrivateRecordSetResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateRecordSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateRecordSetResult)(nil)).Elem()
}

func (o LookupPrivateRecordSetResultOutput) ToLookupPrivateRecordSetResultOutput() LookupPrivateRecordSetResultOutput {
	return o
}

func (o LookupPrivateRecordSetResultOutput) ToLookupPrivateRecordSetResultOutputWithContext(ctx context.Context) LookupPrivateRecordSetResultOutput {
	return o
}

// The list of A records in the record set.
func (o LookupPrivateRecordSetResultOutput) ARecords() ARecordResponseArrayOutput {
	return o.ApplyT(func(v LookupPrivateRecordSetResult) []ARecordResponse { return v.ARecords }).(ARecordResponseArrayOutput)
}

// The list of AAAA records in the record set.
func (o LookupPrivateRecordSetResultOutput) AaaaRecords() AaaaRecordResponseArrayOutput {
	return o.ApplyT(func(v LookupPrivateRecordSetResult) []AaaaRecordResponse { return v.AaaaRecords }).(AaaaRecordResponseArrayOutput)
}

// The CNAME record in the record set.
func (o LookupPrivateRecordSetResultOutput) CnameRecord() CnameRecordResponsePtrOutput {
	return o.ApplyT(func(v LookupPrivateRecordSetResult) *CnameRecordResponse { return v.CnameRecord }).(CnameRecordResponsePtrOutput)
}

// The ETag of the record set.
func (o LookupPrivateRecordSetResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateRecordSetResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Fully qualified domain name of the record set.
func (o LookupPrivateRecordSetResultOutput) Fqdn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateRecordSetResult) string { return v.Fqdn }).(pulumi.StringOutput)
}

// Fully qualified resource Id for the resource. Example - '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateDnsZoneName}'.
func (o LookupPrivateRecordSetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateRecordSetResult) string { return v.Id }).(pulumi.StringOutput)
}

// Is the record set auto-registered in the Private DNS zone through a virtual network link?
func (o LookupPrivateRecordSetResultOutput) IsAutoRegistered() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupPrivateRecordSetResult) bool { return v.IsAutoRegistered }).(pulumi.BoolOutput)
}

// The metadata attached to the record set.
func (o LookupPrivateRecordSetResultOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPrivateRecordSetResult) map[string]string { return v.Metadata }).(pulumi.StringMapOutput)
}

// The list of MX records in the record set.
func (o LookupPrivateRecordSetResultOutput) MxRecords() MxRecordResponseArrayOutput {
	return o.ApplyT(func(v LookupPrivateRecordSetResult) []MxRecordResponse { return v.MxRecords }).(MxRecordResponseArrayOutput)
}

// The name of the resource
func (o LookupPrivateRecordSetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateRecordSetResult) string { return v.Name }).(pulumi.StringOutput)
}

// The list of PTR records in the record set.
func (o LookupPrivateRecordSetResultOutput) PtrRecords() PtrRecordResponseArrayOutput {
	return o.ApplyT(func(v LookupPrivateRecordSetResult) []PtrRecordResponse { return v.PtrRecords }).(PtrRecordResponseArrayOutput)
}

// The SOA record in the record set.
func (o LookupPrivateRecordSetResultOutput) SoaRecord() SoaRecordResponsePtrOutput {
	return o.ApplyT(func(v LookupPrivateRecordSetResult) *SoaRecordResponse { return v.SoaRecord }).(SoaRecordResponsePtrOutput)
}

// The list of SRV records in the record set.
func (o LookupPrivateRecordSetResultOutput) SrvRecords() SrvRecordResponseArrayOutput {
	return o.ApplyT(func(v LookupPrivateRecordSetResult) []SrvRecordResponse { return v.SrvRecords }).(SrvRecordResponseArrayOutput)
}

// The TTL (time-to-live) of the records in the record set.
func (o LookupPrivateRecordSetResultOutput) Ttl() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupPrivateRecordSetResult) *float64 { return v.Ttl }).(pulumi.Float64PtrOutput)
}

// The list of TXT records in the record set.
func (o LookupPrivateRecordSetResultOutput) TxtRecords() TxtRecordResponseArrayOutput {
	return o.ApplyT(func(v LookupPrivateRecordSetResult) []TxtRecordResponse { return v.TxtRecords }).(TxtRecordResponseArrayOutput)
}

// The type of the resource. Example - 'Microsoft.Network/privateDnsZones'.
func (o LookupPrivateRecordSetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateRecordSetResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateRecordSetResultOutput{})
}
