// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package authorization

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This operation lists all the policy set definition versions for all policy set definitions at the management group scope.
//
// Uses Azure REST API version 2023-04-01.
//
// Other available API versions: 2024-05-01, 2025-01-01, 2025-03-01.
func ListPolicySetDefinitionVersionAllAtManagementGroup(ctx *pulumi.Context, args *ListPolicySetDefinitionVersionAllAtManagementGroupArgs, opts ...pulumi.InvokeOption) (*ListPolicySetDefinitionVersionAllAtManagementGroupResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListPolicySetDefinitionVersionAllAtManagementGroupResult
	err := ctx.Invoke("azure-native:authorization:listPolicySetDefinitionVersionAllAtManagementGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListPolicySetDefinitionVersionAllAtManagementGroupArgs struct {
	// The name of the management group. The name is case insensitive.
	ManagementGroupName string `pulumi:"managementGroupName"`
}

// List of policy set definition versions.
type ListPolicySetDefinitionVersionAllAtManagementGroupResult struct {
	// The URL to use for getting the next set of results.
	NextLink *string `pulumi:"nextLink"`
	// An array of policy set definition versions.
	Value []PolicySetDefinitionVersionResponse `pulumi:"value"`
}

func ListPolicySetDefinitionVersionAllAtManagementGroupOutput(ctx *pulumi.Context, args ListPolicySetDefinitionVersionAllAtManagementGroupOutputArgs, opts ...pulumi.InvokeOption) ListPolicySetDefinitionVersionAllAtManagementGroupResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListPolicySetDefinitionVersionAllAtManagementGroupResultOutput, error) {
			args := v.(ListPolicySetDefinitionVersionAllAtManagementGroupArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:authorization:listPolicySetDefinitionVersionAllAtManagementGroup", args, ListPolicySetDefinitionVersionAllAtManagementGroupResultOutput{}, options).(ListPolicySetDefinitionVersionAllAtManagementGroupResultOutput), nil
		}).(ListPolicySetDefinitionVersionAllAtManagementGroupResultOutput)
}

type ListPolicySetDefinitionVersionAllAtManagementGroupOutputArgs struct {
	// The name of the management group. The name is case insensitive.
	ManagementGroupName pulumi.StringInput `pulumi:"managementGroupName"`
}

func (ListPolicySetDefinitionVersionAllAtManagementGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListPolicySetDefinitionVersionAllAtManagementGroupArgs)(nil)).Elem()
}

// List of policy set definition versions.
type ListPolicySetDefinitionVersionAllAtManagementGroupResultOutput struct{ *pulumi.OutputState }

func (ListPolicySetDefinitionVersionAllAtManagementGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListPolicySetDefinitionVersionAllAtManagementGroupResult)(nil)).Elem()
}

func (o ListPolicySetDefinitionVersionAllAtManagementGroupResultOutput) ToListPolicySetDefinitionVersionAllAtManagementGroupResultOutput() ListPolicySetDefinitionVersionAllAtManagementGroupResultOutput {
	return o
}

func (o ListPolicySetDefinitionVersionAllAtManagementGroupResultOutput) ToListPolicySetDefinitionVersionAllAtManagementGroupResultOutputWithContext(ctx context.Context) ListPolicySetDefinitionVersionAllAtManagementGroupResultOutput {
	return o
}

// The URL to use for getting the next set of results.
func (o ListPolicySetDefinitionVersionAllAtManagementGroupResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListPolicySetDefinitionVersionAllAtManagementGroupResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

// An array of policy set definition versions.
func (o ListPolicySetDefinitionVersionAllAtManagementGroupResultOutput) Value() PolicySetDefinitionVersionResponseArrayOutput {
	return o.ApplyT(func(v ListPolicySetDefinitionVersionAllAtManagementGroupResult) []PolicySetDefinitionVersionResponse {
		return v.Value
	}).(PolicySetDefinitionVersionResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListPolicySetDefinitionVersionAllAtManagementGroupResultOutput{})
}
