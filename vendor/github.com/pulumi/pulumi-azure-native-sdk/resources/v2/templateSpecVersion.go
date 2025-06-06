// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package resources

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Template Spec Version object.
//
// Uses Azure REST API version 2022-02-01. In version 1.x of the Azure Native provider, it used API version 2022-02-01.
//
// Other available API versions: 2019-06-01-preview.
type TemplateSpecVersion struct {
	pulumi.CustomResourceState

	// Template Spec version description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// An array of linked template artifacts.
	LinkedTemplates LinkedTemplateArtifactResponseArrayOutput `pulumi:"linkedTemplates"`
	// The location of the Template Spec Version. It must match the location of the parent Template Spec.
	Location pulumi.StringOutput `pulumi:"location"`
	// The main Azure Resource Manager template content.
	MainTemplate pulumi.AnyOutput `pulumi:"mainTemplate"`
	// The version metadata. Metadata is an open-ended object and is typically a collection of key-value pairs.
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// Name of this resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Type of this resource.
	Type pulumi.StringOutput `pulumi:"type"`
	// The Azure Resource Manager template UI definition content.
	UiFormDefinition pulumi.AnyOutput `pulumi:"uiFormDefinition"`
}

// NewTemplateSpecVersion registers a new resource with the given unique name, arguments, and options.
func NewTemplateSpecVersion(ctx *pulumi.Context,
	name string, args *TemplateSpecVersionArgs, opts ...pulumi.ResourceOption) (*TemplateSpecVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.TemplateSpecName == nil {
		return nil, errors.New("invalid value for required argument 'TemplateSpecName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:resources/v20190601preview:TemplateSpecVersion"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20210301preview:TemplateSpecVersion"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20210501:TemplateSpecVersion"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20220201:TemplateSpecVersion"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource TemplateSpecVersion
	err := ctx.RegisterResource("azure-native:resources:TemplateSpecVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTemplateSpecVersion gets an existing TemplateSpecVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTemplateSpecVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TemplateSpecVersionState, opts ...pulumi.ResourceOption) (*TemplateSpecVersion, error) {
	var resource TemplateSpecVersion
	err := ctx.ReadResource("azure-native:resources:TemplateSpecVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TemplateSpecVersion resources.
type templateSpecVersionState struct {
}

type TemplateSpecVersionState struct {
}

func (TemplateSpecVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*templateSpecVersionState)(nil)).Elem()
}

type templateSpecVersionArgs struct {
	// Template Spec version description.
	Description *string `pulumi:"description"`
	// An array of linked template artifacts.
	LinkedTemplates []LinkedTemplateArtifact `pulumi:"linkedTemplates"`
	// The location of the Template Spec Version. It must match the location of the parent Template Spec.
	Location *string `pulumi:"location"`
	// The main Azure Resource Manager template content.
	MainTemplate interface{} `pulumi:"mainTemplate"`
	// The version metadata. Metadata is an open-ended object and is typically a collection of key-value pairs.
	Metadata interface{} `pulumi:"metadata"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Name of the Template Spec.
	TemplateSpecName string `pulumi:"templateSpecName"`
	// The version of the Template Spec.
	TemplateSpecVersion *string `pulumi:"templateSpecVersion"`
	// The Azure Resource Manager template UI definition content.
	UiFormDefinition interface{} `pulumi:"uiFormDefinition"`
}

// The set of arguments for constructing a TemplateSpecVersion resource.
type TemplateSpecVersionArgs struct {
	// Template Spec version description.
	Description pulumi.StringPtrInput
	// An array of linked template artifacts.
	LinkedTemplates LinkedTemplateArtifactArrayInput
	// The location of the Template Spec Version. It must match the location of the parent Template Spec.
	Location pulumi.StringPtrInput
	// The main Azure Resource Manager template content.
	MainTemplate pulumi.Input
	// The version metadata. Metadata is an open-ended object and is typically a collection of key-value pairs.
	Metadata pulumi.Input
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Name of the Template Spec.
	TemplateSpecName pulumi.StringInput
	// The version of the Template Spec.
	TemplateSpecVersion pulumi.StringPtrInput
	// The Azure Resource Manager template UI definition content.
	UiFormDefinition pulumi.Input
}

func (TemplateSpecVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*templateSpecVersionArgs)(nil)).Elem()
}

type TemplateSpecVersionInput interface {
	pulumi.Input

	ToTemplateSpecVersionOutput() TemplateSpecVersionOutput
	ToTemplateSpecVersionOutputWithContext(ctx context.Context) TemplateSpecVersionOutput
}

func (*TemplateSpecVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**TemplateSpecVersion)(nil)).Elem()
}

func (i *TemplateSpecVersion) ToTemplateSpecVersionOutput() TemplateSpecVersionOutput {
	return i.ToTemplateSpecVersionOutputWithContext(context.Background())
}

func (i *TemplateSpecVersion) ToTemplateSpecVersionOutputWithContext(ctx context.Context) TemplateSpecVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateSpecVersionOutput)
}

type TemplateSpecVersionOutput struct{ *pulumi.OutputState }

func (TemplateSpecVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TemplateSpecVersion)(nil)).Elem()
}

func (o TemplateSpecVersionOutput) ToTemplateSpecVersionOutput() TemplateSpecVersionOutput {
	return o
}

func (o TemplateSpecVersionOutput) ToTemplateSpecVersionOutputWithContext(ctx context.Context) TemplateSpecVersionOutput {
	return o
}

// Template Spec version description.
func (o TemplateSpecVersionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TemplateSpecVersion) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// An array of linked template artifacts.
func (o TemplateSpecVersionOutput) LinkedTemplates() LinkedTemplateArtifactResponseArrayOutput {
	return o.ApplyT(func(v *TemplateSpecVersion) LinkedTemplateArtifactResponseArrayOutput { return v.LinkedTemplates }).(LinkedTemplateArtifactResponseArrayOutput)
}

// The location of the Template Spec Version. It must match the location of the parent Template Spec.
func (o TemplateSpecVersionOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *TemplateSpecVersion) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The main Azure Resource Manager template content.
func (o TemplateSpecVersionOutput) MainTemplate() pulumi.AnyOutput {
	return o.ApplyT(func(v *TemplateSpecVersion) pulumi.AnyOutput { return v.MainTemplate }).(pulumi.AnyOutput)
}

// The version metadata. Metadata is an open-ended object and is typically a collection of key-value pairs.
func (o TemplateSpecVersionOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v *TemplateSpecVersion) pulumi.AnyOutput { return v.Metadata }).(pulumi.AnyOutput)
}

// Name of this resource.
func (o TemplateSpecVersionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TemplateSpecVersion) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o TemplateSpecVersionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *TemplateSpecVersion) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o TemplateSpecVersionOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *TemplateSpecVersion) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Type of this resource.
func (o TemplateSpecVersionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *TemplateSpecVersion) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The Azure Resource Manager template UI definition content.
func (o TemplateSpecVersionOutput) UiFormDefinition() pulumi.AnyOutput {
	return o.ApplyT(func(v *TemplateSpecVersion) pulumi.AnyOutput { return v.UiFormDefinition }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterOutputType(TemplateSpecVersionOutput{})
}
