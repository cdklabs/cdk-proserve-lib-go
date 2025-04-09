package types

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.types.AwsCustomResourceLambdaConfiguration",
		reflect.TypeOf((*AwsCustomResourceLambdaConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-proserve-lib.types.AwsManagedPolicy",
		reflect.TypeOf((*AwsManagedPolicy)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_AwsManagedPolicy{}
		},
	)
	_jsii_.RegisterEnum(
		"@cdklabs/cdk-proserve-lib.types.DestructiveOperation",
		reflect.TypeOf((*DestructiveOperation)(nil)).Elem(),
		map[string]interface{}{
			"UPDATE": DestructiveOperation_UPDATE,
			"DELETE": DestructiveOperation_DELETE,
			"ALL": DestructiveOperation_ALL,
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-proserve-lib.types.Ec2InstanceType",
		reflect.TypeOf((*Ec2InstanceType)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_Ec2InstanceType{}
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.types.LambdaConfiguration",
		reflect.TypeOf((*LambdaConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-proserve-lib.types.SageMakerNotebookInstanceType",
		reflect.TypeOf((*SageMakerNotebookInstanceType)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_SageMakerNotebookInstanceType{}
		},
	)
}
