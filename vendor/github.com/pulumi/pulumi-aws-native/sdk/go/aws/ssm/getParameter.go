// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ssm

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The “AWS::SSM::Parameter“ resource creates an SSM parameter in SYSlong Parameter Store.
//
//	 To create an SSM parameter, you must have the IAMlong (IAM) permissions ``ssm:PutParameter`` and ``ssm:AddTagsToResource``. On stack creation, CFNlong adds the following three tags to the parameter: ``aws:cloudformation:stack-name``, ``aws:cloudformation:logical-id``, and ``aws:cloudformation:stack-id``, in addition to any custom tags you specify.
//	To add, update, or remove tags during stack update, you must have IAM permissions for both ``ssm:AddTagsToResource`` and ``ssm:RemoveTagsFromResource``. For more information, see [Managing Access Using Policies](https://docs.aws.amazon.com/systems-manager/latest/userguide/security-iam.html#security_iam_access-manage) in the *User Guide*.
//	 For information about valid values for parameters, see [About requirements and constraints for parameter names](https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-paramstore-su-create.html#sysman-parameter-name-constraints) in the *User Guide* and [PutParameter](https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_PutParameter.html) in the *API Reference*.
func LookupParameter(ctx *pulumi.Context, args *LookupParameterArgs, opts ...pulumi.InvokeOption) (*LookupParameterResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupParameterResult
	err := ctx.Invoke("aws-native:ssm:getParameter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupParameterArgs struct {
	// The name of the parameter.
	//   The maximum length constraint listed below includes capacity for additional system attributes that aren't part of the name. The maximum length for a parameter name, including the full length of the parameter Amazon Resource Name (ARN), is 1011 characters. For example, the length of the following parameter name is 65 characters, not 20 characters: ``arn:aws:ssm:us-east-2:111222333444:parameter/ExampleParameterName``
	Name string `pulumi:"name"`
}

type LookupParameterResult struct {
	// The data type of the parameter, such as ``text`` or ``aws:ec2:image``. The default is ``text``.
	DataType *ParameterDataType `pulumi:"dataType"`
	// The type of parameter.
	Type *ParameterType `pulumi:"type"`
	// The parameter value.
	//   If type is ``StringList``, the system returns a comma-separated string with no spaces between commas in the ``Value`` field.
	Value *string `pulumi:"value"`
}

func LookupParameterOutput(ctx *pulumi.Context, args LookupParameterOutputArgs, opts ...pulumi.InvokeOption) LookupParameterResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupParameterResultOutput, error) {
			args := v.(LookupParameterArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:ssm:getParameter", args, LookupParameterResultOutput{}, options).(LookupParameterResultOutput), nil
		}).(LookupParameterResultOutput)
}

type LookupParameterOutputArgs struct {
	// The name of the parameter.
	//   The maximum length constraint listed below includes capacity for additional system attributes that aren't part of the name. The maximum length for a parameter name, including the full length of the parameter Amazon Resource Name (ARN), is 1011 characters. For example, the length of the following parameter name is 65 characters, not 20 characters: ``arn:aws:ssm:us-east-2:111222333444:parameter/ExampleParameterName``
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupParameterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupParameterArgs)(nil)).Elem()
}

type LookupParameterResultOutput struct{ *pulumi.OutputState }

func (LookupParameterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupParameterResult)(nil)).Elem()
}

func (o LookupParameterResultOutput) ToLookupParameterResultOutput() LookupParameterResultOutput {
	return o
}

func (o LookupParameterResultOutput) ToLookupParameterResultOutputWithContext(ctx context.Context) LookupParameterResultOutput {
	return o
}

// The data type of the parameter, such as “text“ or “aws:ec2:image“. The default is “text“.
func (o LookupParameterResultOutput) DataType() ParameterDataTypePtrOutput {
	return o.ApplyT(func(v LookupParameterResult) *ParameterDataType { return v.DataType }).(ParameterDataTypePtrOutput)
}

// The type of parameter.
func (o LookupParameterResultOutput) Type() ParameterTypePtrOutput {
	return o.ApplyT(func(v LookupParameterResult) *ParameterType { return v.Type }).(ParameterTypePtrOutput)
}

// The parameter value.
//
//	If type is ``StringList``, the system returns a comma-separated string with no spaces between commas in the ``Value`` field.
func (o LookupParameterResultOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupParameterResult) *string { return v.Value }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupParameterResultOutput{})
}
