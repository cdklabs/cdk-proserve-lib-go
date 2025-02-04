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
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.types.LambdaConfiguration",
		reflect.TypeOf((*LambdaConfiguration)(nil)).Elem(),
	)
}
