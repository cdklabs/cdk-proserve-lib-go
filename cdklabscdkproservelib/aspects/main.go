package aspects

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdklabs/cdk-proserve-lib.aspects.ApplyRemovalPolicy",
		reflect.TypeOf((*ApplyRemovalPolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "visit", GoMethod: "Visit"},
		},
		func() interface{} {
			j := jsiiProxy_ApplyRemovalPolicy{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIAspect)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.aspects.ApplyRemovalPolicyProps",
		reflect.TypeOf((*ApplyRemovalPolicyProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-proserve-lib.aspects.CreateLambdaLogGroup",
		reflect.TypeOf((*CreateLambdaLogGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "visit", GoMethod: "Visit"},
		},
		func() interface{} {
			j := jsiiProxy_CreateLambdaLogGroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIAspect)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-proserve-lib.aspects.SecureSageMakerNotebook",
		reflect.TypeOf((*SecureSageMakerNotebook)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "visit", GoMethod: "Visit"},
		},
		func() interface{} {
			j := jsiiProxy_SecureSageMakerNotebook{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIAspect)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.aspects.SecureSageMakerNotebookProps",
		reflect.TypeOf((*SecureSageMakerNotebookProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-proserve-lib.aspects.SetLogRetention",
		reflect.TypeOf((*SetLogRetention)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "visit", GoMethod: "Visit"},
		},
		func() interface{} {
			j := jsiiProxy_SetLogRetention{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIAspect)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.aspects.SetLogRetentionProps",
		reflect.TypeOf((*SetLogRetentionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-proserve-lib.aspects.SqsRequireSsl",
		reflect.TypeOf((*SqsRequireSsl)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "visit", GoMethod: "Visit"},
		},
		func() interface{} {
			j := jsiiProxy_SqsRequireSsl{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIAspect)
			return &j
		},
	)
}
