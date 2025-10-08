package aspects

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.aspects.AlarmConfig",
		reflect.TypeOf((*AlarmConfig)(nil)).Elem(),
	)
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
		"@cdklabs/cdk-proserve-lib.aspects.Ec2AutomatedShutdown",
		reflect.TypeOf((*Ec2AutomatedShutdown)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "visit", GoMethod: "Visit"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2AutomatedShutdown{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIAspect)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"@cdklabs/cdk-proserve-lib.aspects.Ec2AutomatedShutdown.Ec2MetricName",
		reflect.TypeOf((*Ec2AutomatedShutdown_Ec2MetricName)(nil)).Elem(),
		map[string]interface{}{
			"CPU_UTILIZATION": Ec2AutomatedShutdown_Ec2MetricName_CPU_UTILIZATION,
			"DISK_READ_OPS": Ec2AutomatedShutdown_Ec2MetricName_DISK_READ_OPS,
			"DISK_WRITE_OPS": Ec2AutomatedShutdown_Ec2MetricName_DISK_WRITE_OPS,
			"DISK_READ_BYTES": Ec2AutomatedShutdown_Ec2MetricName_DISK_READ_BYTES,
			"DISK_WRITE_BYTES": Ec2AutomatedShutdown_Ec2MetricName_DISK_WRITE_BYTES,
			"NETWORK_IN": Ec2AutomatedShutdown_Ec2MetricName_NETWORK_IN,
			"NETWORK_OUT": Ec2AutomatedShutdown_Ec2MetricName_NETWORK_OUT,
			"NETWORK_PACKETS_IN": Ec2AutomatedShutdown_Ec2MetricName_NETWORK_PACKETS_IN,
			"NETWORK_PACKETS_OUT": Ec2AutomatedShutdown_Ec2MetricName_NETWORK_PACKETS_OUT,
			"STATUS_CHECK_FAILED": Ec2AutomatedShutdown_Ec2MetricName_STATUS_CHECK_FAILED,
			"STATUS_CHECK_FAILED_INSTANCE": Ec2AutomatedShutdown_Ec2MetricName_STATUS_CHECK_FAILED_INSTANCE,
			"STATUS_CHECK_FAILED_SYSTEM": Ec2AutomatedShutdown_Ec2MetricName_STATUS_CHECK_FAILED_SYSTEM,
			"METADATA_NO_TOKEN": Ec2AutomatedShutdown_Ec2MetricName_METADATA_NO_TOKEN,
			"CPU_CREDIT_USAGE": Ec2AutomatedShutdown_Ec2MetricName_CPU_CREDIT_USAGE,
			"CPU_CREDIT_BALANCE": Ec2AutomatedShutdown_Ec2MetricName_CPU_CREDIT_BALANCE,
			"CPU_SURPLUS_CREDIT_BALANCE": Ec2AutomatedShutdown_Ec2MetricName_CPU_SURPLUS_CREDIT_BALANCE,
			"CPU_SURPLUS_CREDITS_CHARGED": Ec2AutomatedShutdown_Ec2MetricName_CPU_SURPLUS_CREDITS_CHARGED,
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.aspects.Ec2AutomatedShutdownProps",
		reflect.TypeOf((*Ec2AutomatedShutdownProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-proserve-lib.aspects.RdsOracleMultiTenant",
		reflect.TypeOf((*RdsOracleMultiTenant)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberMethod{JsiiMethod: "visit", GoMethod: "Visit"},
		},
		func() interface{} {
			j := jsiiProxy_RdsOracleMultiTenant{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIAspect)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.aspects.RdsOracleMultiTenantProps",
		reflect.TypeOf((*RdsOracleMultiTenantProps)(nil)).Elem(),
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
		"@cdklabs/cdk-proserve-lib.aspects.SecurityCompliance",
		reflect.TypeOf((*SecurityCompliance)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "visit", GoMethod: "Visit"},
		},
		func() interface{} {
			j := jsiiProxy_SecurityCompliance{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIAspect)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.aspects.SecurityCompliance.ApiGatewaySettings",
		reflect.TypeOf((*SecurityCompliance_ApiGatewaySettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.aspects.SecurityCompliance.DisableableSetting",
		reflect.TypeOf((*SecurityCompliance_DisableableSetting)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.aspects.SecurityCompliance.DynamoDbSettings",
		reflect.TypeOf((*SecurityCompliance_DynamoDbSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.aspects.SecurityCompliance.EcsSettings",
		reflect.TypeOf((*SecurityCompliance_EcsSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.aspects.SecurityCompliance.LambdaSettings",
		reflect.TypeOf((*SecurityCompliance_LambdaSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.aspects.SecurityCompliance.ReservedConcurrentSettings",
		reflect.TypeOf((*SecurityCompliance_ReservedConcurrentSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.aspects.SecurityCompliance.S3Settings",
		reflect.TypeOf((*SecurityCompliance_S3Settings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.aspects.SecurityCompliance.ServerAccessLogsSettings",
		reflect.TypeOf((*SecurityCompliance_ServerAccessLogsSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.aspects.SecurityCompliance.Settings",
		reflect.TypeOf((*SecurityCompliance_Settings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.aspects.SecurityCompliance.StageMethodLogging",
		reflect.TypeOf((*SecurityCompliance_StageMethodLogging)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.aspects.SecurityCompliance.StepFunctionsSettings",
		reflect.TypeOf((*SecurityCompliance_StepFunctionsSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.aspects.SecurityCompliance.Suppressions",
		reflect.TypeOf((*SecurityCompliance_Suppressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.aspects.SecurityComplianceProps",
		reflect.TypeOf((*SecurityComplianceProps)(nil)).Elem(),
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
