// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Snapshot resource.
//
// Uses Azure REST API version 2022-07-02. In version 1.x of the Azure Native provider, it used API version 2020-12-01.
//
// Other available API versions: 2023-01-02, 2023-04-02, 2023-10-02, 2024-03-02.
type Snapshot struct {
	pulumi.CustomResourceState

	// Percentage complete for the background copy when a resource is created via the CopyStart operation.
	CompletionPercent pulumi.Float64PtrOutput `pulumi:"completionPercent"`
	// Indicates the error details if the background copy of a resource created via the CopyStart operation fails.
	CopyCompletionError CopyCompletionErrorResponsePtrOutput `pulumi:"copyCompletionError"`
	// Disk source information. CreationData information cannot be changed after the disk has been created.
	CreationData CreationDataResponseOutput `pulumi:"creationData"`
	// Additional authentication requirements when exporting or uploading to a disk or snapshot.
	DataAccessAuthMode pulumi.StringPtrOutput `pulumi:"dataAccessAuthMode"`
	// ARM id of the DiskAccess resource for using private endpoints on disks.
	DiskAccessId pulumi.StringPtrOutput `pulumi:"diskAccessId"`
	// The size of the disk in bytes. This field is read only.
	DiskSizeBytes pulumi.Float64Output `pulumi:"diskSizeBytes"`
	// If creationData.createOption is Empty, this field is mandatory and it indicates the size of the disk to create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only allowed if the disk is not attached to a running VM, and can only increase the disk's size.
	DiskSizeGB pulumi.IntPtrOutput `pulumi:"diskSizeGB"`
	// The state of the snapshot.
	DiskState pulumi.StringOutput `pulumi:"diskState"`
	// Encryption property can be used to encrypt data at rest with customer managed keys or platform managed keys.
	Encryption EncryptionResponsePtrOutput `pulumi:"encryption"`
	// Encryption settings collection used be Azure Disk Encryption, can contain multiple encryption settings per disk or snapshot.
	EncryptionSettingsCollection EncryptionSettingsCollectionResponsePtrOutput `pulumi:"encryptionSettingsCollection"`
	// The extended location where the snapshot will be created. Extended location cannot be changed.
	ExtendedLocation ExtendedLocationResponsePtrOutput `pulumi:"extendedLocation"`
	// The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
	HyperVGeneration pulumi.StringPtrOutput `pulumi:"hyperVGeneration"`
	// Whether a snapshot is incremental. Incremental snapshots on the same disk occupy less space than full snapshots and can be diffed.
	Incremental pulumi.BoolPtrOutput `pulumi:"incremental"`
	// Incremental snapshots for a disk share an incremental snapshot family id. The Get Page Range Diff API can only be called on incremental snapshots with the same family id.
	IncrementalSnapshotFamilyId pulumi.StringOutput `pulumi:"incrementalSnapshotFamilyId"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Unused. Always Null.
	ManagedBy pulumi.StringOutput `pulumi:"managedBy"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Policy for accessing the disk via network.
	NetworkAccessPolicy pulumi.StringPtrOutput `pulumi:"networkAccessPolicy"`
	// The Operating System type.
	OsType pulumi.StringPtrOutput `pulumi:"osType"`
	// The disk provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Policy for controlling export on the disk.
	PublicNetworkAccess pulumi.StringPtrOutput `pulumi:"publicNetworkAccess"`
	// Purchase plan information for the image from which the source disk for the snapshot was originally created.
	PurchasePlan PurchasePlanResponsePtrOutput `pulumi:"purchasePlan"`
	// Contains the security related information for the resource.
	SecurityProfile DiskSecurityProfileResponsePtrOutput `pulumi:"securityProfile"`
	// The snapshots sku name. Can be Standard_LRS, Premium_LRS, or Standard_ZRS. This is an optional parameter for incremental snapshot and the default behavior is the SKU will be set to the same sku as the previous snapshot
	Sku SnapshotSkuResponsePtrOutput `pulumi:"sku"`
	// List of supported capabilities for the image from which the source disk from the snapshot was originally created.
	SupportedCapabilities SupportedCapabilitiesResponsePtrOutput `pulumi:"supportedCapabilities"`
	// Indicates the OS on a snapshot supports hibernation.
	SupportsHibernation pulumi.BoolPtrOutput `pulumi:"supportsHibernation"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The time when the snapshot was created.
	TimeCreated pulumi.StringOutput `pulumi:"timeCreated"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// Unique Guid identifying the resource.
	UniqueId pulumi.StringOutput `pulumi:"uniqueId"`
}

// NewSnapshot registers a new resource with the given unique name, arguments, and options.
func NewSnapshot(ctx *pulumi.Context,
	name string, args *SnapshotArgs, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CreationData == nil {
		return nil, errors.New("invalid value for required argument 'CreationData'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:compute/v20160430preview:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20170330:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20180401:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20180601:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20180930:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190301:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190701:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20191101:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200501:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200630:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200930:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20201201:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210401:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210801:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20211201:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220302:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220702:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20230102:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20230402:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20231002:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20240302:Snapshot"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Snapshot
	err := ctx.RegisterResource("azure-native:compute:Snapshot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSnapshot gets an existing Snapshot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnapshot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnapshotState, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	var resource Snapshot
	err := ctx.ReadResource("azure-native:compute:Snapshot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Snapshot resources.
type snapshotState struct {
}

type SnapshotState struct {
}

func (SnapshotState) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotState)(nil)).Elem()
}

type snapshotArgs struct {
	// Percentage complete for the background copy when a resource is created via the CopyStart operation.
	CompletionPercent *float64 `pulumi:"completionPercent"`
	// Indicates the error details if the background copy of a resource created via the CopyStart operation fails.
	CopyCompletionError *CopyCompletionError `pulumi:"copyCompletionError"`
	// Disk source information. CreationData information cannot be changed after the disk has been created.
	CreationData CreationData `pulumi:"creationData"`
	// Additional authentication requirements when exporting or uploading to a disk or snapshot.
	DataAccessAuthMode *string `pulumi:"dataAccessAuthMode"`
	// ARM id of the DiskAccess resource for using private endpoints on disks.
	DiskAccessId *string `pulumi:"diskAccessId"`
	// If creationData.createOption is Empty, this field is mandatory and it indicates the size of the disk to create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only allowed if the disk is not attached to a running VM, and can only increase the disk's size.
	DiskSizeGB *int `pulumi:"diskSizeGB"`
	// Encryption property can be used to encrypt data at rest with customer managed keys or platform managed keys.
	Encryption *Encryption `pulumi:"encryption"`
	// Encryption settings collection used be Azure Disk Encryption, can contain multiple encryption settings per disk or snapshot.
	EncryptionSettingsCollection *EncryptionSettingsCollection `pulumi:"encryptionSettingsCollection"`
	// The extended location where the snapshot will be created. Extended location cannot be changed.
	ExtendedLocation *ExtendedLocation `pulumi:"extendedLocation"`
	// The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
	HyperVGeneration *string `pulumi:"hyperVGeneration"`
	// Whether a snapshot is incremental. Incremental snapshots on the same disk occupy less space than full snapshots and can be diffed.
	Incremental *bool `pulumi:"incremental"`
	// Resource location
	Location *string `pulumi:"location"`
	// Policy for accessing the disk via network.
	NetworkAccessPolicy *string `pulumi:"networkAccessPolicy"`
	// The Operating System type.
	OsType *OperatingSystemTypes `pulumi:"osType"`
	// Policy for controlling export on the disk.
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// Purchase plan information for the image from which the source disk for the snapshot was originally created.
	PurchasePlan *PurchasePlan `pulumi:"purchasePlan"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Contains the security related information for the resource.
	SecurityProfile *DiskSecurityProfile `pulumi:"securityProfile"`
	// The snapshots sku name. Can be Standard_LRS, Premium_LRS, or Standard_ZRS. This is an optional parameter for incremental snapshot and the default behavior is the SKU will be set to the same sku as the previous snapshot
	Sku *SnapshotSku `pulumi:"sku"`
	// The name of the snapshot that is being created. The name can't be changed after the snapshot is created. Supported characters for the name are a-z, A-Z, 0-9, _ and -. The max name length is 80 characters.
	SnapshotName *string `pulumi:"snapshotName"`
	// List of supported capabilities for the image from which the source disk from the snapshot was originally created.
	SupportedCapabilities *SupportedCapabilities `pulumi:"supportedCapabilities"`
	// Indicates the OS on a snapshot supports hibernation.
	SupportsHibernation *bool `pulumi:"supportsHibernation"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Snapshot resource.
type SnapshotArgs struct {
	// Percentage complete for the background copy when a resource is created via the CopyStart operation.
	CompletionPercent pulumi.Float64PtrInput
	// Indicates the error details if the background copy of a resource created via the CopyStart operation fails.
	CopyCompletionError CopyCompletionErrorPtrInput
	// Disk source information. CreationData information cannot be changed after the disk has been created.
	CreationData CreationDataInput
	// Additional authentication requirements when exporting or uploading to a disk or snapshot.
	DataAccessAuthMode pulumi.StringPtrInput
	// ARM id of the DiskAccess resource for using private endpoints on disks.
	DiskAccessId pulumi.StringPtrInput
	// If creationData.createOption is Empty, this field is mandatory and it indicates the size of the disk to create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only allowed if the disk is not attached to a running VM, and can only increase the disk's size.
	DiskSizeGB pulumi.IntPtrInput
	// Encryption property can be used to encrypt data at rest with customer managed keys or platform managed keys.
	Encryption EncryptionPtrInput
	// Encryption settings collection used be Azure Disk Encryption, can contain multiple encryption settings per disk or snapshot.
	EncryptionSettingsCollection EncryptionSettingsCollectionPtrInput
	// The extended location where the snapshot will be created. Extended location cannot be changed.
	ExtendedLocation ExtendedLocationPtrInput
	// The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
	HyperVGeneration pulumi.StringPtrInput
	// Whether a snapshot is incremental. Incremental snapshots on the same disk occupy less space than full snapshots and can be diffed.
	Incremental pulumi.BoolPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// Policy for accessing the disk via network.
	NetworkAccessPolicy pulumi.StringPtrInput
	// The Operating System type.
	OsType OperatingSystemTypesPtrInput
	// Policy for controlling export on the disk.
	PublicNetworkAccess pulumi.StringPtrInput
	// Purchase plan information for the image from which the source disk for the snapshot was originally created.
	PurchasePlan PurchasePlanPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Contains the security related information for the resource.
	SecurityProfile DiskSecurityProfilePtrInput
	// The snapshots sku name. Can be Standard_LRS, Premium_LRS, or Standard_ZRS. This is an optional parameter for incremental snapshot and the default behavior is the SKU will be set to the same sku as the previous snapshot
	Sku SnapshotSkuPtrInput
	// The name of the snapshot that is being created. The name can't be changed after the snapshot is created. Supported characters for the name are a-z, A-Z, 0-9, _ and -. The max name length is 80 characters.
	SnapshotName pulumi.StringPtrInput
	// List of supported capabilities for the image from which the source disk from the snapshot was originally created.
	SupportedCapabilities SupportedCapabilitiesPtrInput
	// Indicates the OS on a snapshot supports hibernation.
	SupportsHibernation pulumi.BoolPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
}

func (SnapshotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotArgs)(nil)).Elem()
}

type SnapshotInput interface {
	pulumi.Input

	ToSnapshotOutput() SnapshotOutput
	ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput
}

func (*Snapshot) ElementType() reflect.Type {
	return reflect.TypeOf((**Snapshot)(nil)).Elem()
}

func (i *Snapshot) ToSnapshotOutput() SnapshotOutput {
	return i.ToSnapshotOutputWithContext(context.Background())
}

func (i *Snapshot) ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotOutput)
}

type SnapshotOutput struct{ *pulumi.OutputState }

func (SnapshotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Snapshot)(nil)).Elem()
}

func (o SnapshotOutput) ToSnapshotOutput() SnapshotOutput {
	return o
}

func (o SnapshotOutput) ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput {
	return o
}

// Percentage complete for the background copy when a resource is created via the CopyStart operation.
func (o SnapshotOutput) CompletionPercent() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.Float64PtrOutput { return v.CompletionPercent }).(pulumi.Float64PtrOutput)
}

// Indicates the error details if the background copy of a resource created via the CopyStart operation fails.
func (o SnapshotOutput) CopyCompletionError() CopyCompletionErrorResponsePtrOutput {
	return o.ApplyT(func(v *Snapshot) CopyCompletionErrorResponsePtrOutput { return v.CopyCompletionError }).(CopyCompletionErrorResponsePtrOutput)
}

// Disk source information. CreationData information cannot be changed after the disk has been created.
func (o SnapshotOutput) CreationData() CreationDataResponseOutput {
	return o.ApplyT(func(v *Snapshot) CreationDataResponseOutput { return v.CreationData }).(CreationDataResponseOutput)
}

// Additional authentication requirements when exporting or uploading to a disk or snapshot.
func (o SnapshotOutput) DataAccessAuthMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringPtrOutput { return v.DataAccessAuthMode }).(pulumi.StringPtrOutput)
}

// ARM id of the DiskAccess resource for using private endpoints on disks.
func (o SnapshotOutput) DiskAccessId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringPtrOutput { return v.DiskAccessId }).(pulumi.StringPtrOutput)
}

// The size of the disk in bytes. This field is read only.
func (o SnapshotOutput) DiskSizeBytes() pulumi.Float64Output {
	return o.ApplyT(func(v *Snapshot) pulumi.Float64Output { return v.DiskSizeBytes }).(pulumi.Float64Output)
}

// If creationData.createOption is Empty, this field is mandatory and it indicates the size of the disk to create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only allowed if the disk is not attached to a running VM, and can only increase the disk's size.
func (o SnapshotOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.IntPtrOutput { return v.DiskSizeGB }).(pulumi.IntPtrOutput)
}

// The state of the snapshot.
func (o SnapshotOutput) DiskState() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.DiskState }).(pulumi.StringOutput)
}

// Encryption property can be used to encrypt data at rest with customer managed keys or platform managed keys.
func (o SnapshotOutput) Encryption() EncryptionResponsePtrOutput {
	return o.ApplyT(func(v *Snapshot) EncryptionResponsePtrOutput { return v.Encryption }).(EncryptionResponsePtrOutput)
}

// Encryption settings collection used be Azure Disk Encryption, can contain multiple encryption settings per disk or snapshot.
func (o SnapshotOutput) EncryptionSettingsCollection() EncryptionSettingsCollectionResponsePtrOutput {
	return o.ApplyT(func(v *Snapshot) EncryptionSettingsCollectionResponsePtrOutput { return v.EncryptionSettingsCollection }).(EncryptionSettingsCollectionResponsePtrOutput)
}

// The extended location where the snapshot will be created. Extended location cannot be changed.
func (o SnapshotOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *Snapshot) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
func (o SnapshotOutput) HyperVGeneration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringPtrOutput { return v.HyperVGeneration }).(pulumi.StringPtrOutput)
}

// Whether a snapshot is incremental. Incremental snapshots on the same disk occupy less space than full snapshots and can be diffed.
func (o SnapshotOutput) Incremental() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.BoolPtrOutput { return v.Incremental }).(pulumi.BoolPtrOutput)
}

// Incremental snapshots for a disk share an incremental snapshot family id. The Get Page Range Diff API can only be called on incremental snapshots with the same family id.
func (o SnapshotOutput) IncrementalSnapshotFamilyId() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.IncrementalSnapshotFamilyId }).(pulumi.StringOutput)
}

// Resource location
func (o SnapshotOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Unused. Always Null.
func (o SnapshotOutput) ManagedBy() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.ManagedBy }).(pulumi.StringOutput)
}

// Resource name
func (o SnapshotOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Policy for accessing the disk via network.
func (o SnapshotOutput) NetworkAccessPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringPtrOutput { return v.NetworkAccessPolicy }).(pulumi.StringPtrOutput)
}

// The Operating System type.
func (o SnapshotOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringPtrOutput { return v.OsType }).(pulumi.StringPtrOutput)
}

// The disk provisioning state.
func (o SnapshotOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Policy for controlling export on the disk.
func (o SnapshotOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringPtrOutput { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// Purchase plan information for the image from which the source disk for the snapshot was originally created.
func (o SnapshotOutput) PurchasePlan() PurchasePlanResponsePtrOutput {
	return o.ApplyT(func(v *Snapshot) PurchasePlanResponsePtrOutput { return v.PurchasePlan }).(PurchasePlanResponsePtrOutput)
}

// Contains the security related information for the resource.
func (o SnapshotOutput) SecurityProfile() DiskSecurityProfileResponsePtrOutput {
	return o.ApplyT(func(v *Snapshot) DiskSecurityProfileResponsePtrOutput { return v.SecurityProfile }).(DiskSecurityProfileResponsePtrOutput)
}

// The snapshots sku name. Can be Standard_LRS, Premium_LRS, or Standard_ZRS. This is an optional parameter for incremental snapshot and the default behavior is the SKU will be set to the same sku as the previous snapshot
func (o SnapshotOutput) Sku() SnapshotSkuResponsePtrOutput {
	return o.ApplyT(func(v *Snapshot) SnapshotSkuResponsePtrOutput { return v.Sku }).(SnapshotSkuResponsePtrOutput)
}

// List of supported capabilities for the image from which the source disk from the snapshot was originally created.
func (o SnapshotOutput) SupportedCapabilities() SupportedCapabilitiesResponsePtrOutput {
	return o.ApplyT(func(v *Snapshot) SupportedCapabilitiesResponsePtrOutput { return v.SupportedCapabilities }).(SupportedCapabilitiesResponsePtrOutput)
}

// Indicates the OS on a snapshot supports hibernation.
func (o SnapshotOutput) SupportsHibernation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.BoolPtrOutput { return v.SupportsHibernation }).(pulumi.BoolPtrOutput)
}

// Resource tags
func (o SnapshotOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The time when the snapshot was created.
func (o SnapshotOutput) TimeCreated() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.TimeCreated }).(pulumi.StringOutput)
}

// Resource type
func (o SnapshotOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Unique Guid identifying the resource.
func (o SnapshotOutput) UniqueId() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.UniqueId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SnapshotOutput{})
}
