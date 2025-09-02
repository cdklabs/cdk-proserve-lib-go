package patterns

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdklabs/cdk-proserve-lib.patterns.ApiGatewayStaticHosting",
		reflect.TypeOf((*ApiGatewayStaticHosting)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "components", GoGetter: "Components"},
			_jsii_.MemberProperty{JsiiProperty: "customDomainNameAlias", GoGetter: "CustomDomainNameAlias"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "url", GoGetter: "Url"},
		},
		func() interface{} {
			j := jsiiProxy_ApiGatewayStaticHosting{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.patterns.ApiGatewayStaticHosting.Asset",
		reflect.TypeOf((*ApiGatewayStaticHosting_Asset)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.patterns.ApiGatewayStaticHosting.CustomDomainConfiguration",
		reflect.TypeOf((*ApiGatewayStaticHosting_CustomDomainConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.patterns.ApiGatewayStaticHosting.DefaultEndpointConfiguration",
		reflect.TypeOf((*ApiGatewayStaticHosting_DefaultEndpointConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.patterns.ApiGatewayStaticHosting.PatternComponents",
		reflect.TypeOf((*ApiGatewayStaticHosting_PatternComponents)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.patterns.ApiGatewayStaticHostingProps",
		reflect.TypeOf((*ApiGatewayStaticHostingProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-proserve-lib.patterns.Ec2LinuxImagePipeline",
		reflect.TypeOf((*Ec2LinuxImagePipeline)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "imagePipelineArn", GoGetter: "ImagePipelineArn"},
			_jsii_.MemberProperty{JsiiProperty: "imagePipelineTopic", GoGetter: "ImagePipelineTopic"},
			_jsii_.MemberProperty{JsiiProperty: "latestAmi", GoGetter: "LatestAmi"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2LinuxImagePipeline{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"@cdklabs/cdk-proserve-lib.patterns.Ec2LinuxImagePipeline.Feature",
		reflect.TypeOf((*Ec2LinuxImagePipeline_Feature)(nil)).Elem(),
		map[string]interface{}{
			"AWS_CLI": Ec2LinuxImagePipeline_Feature_AWS_CLI,
			"NICE_DCV": Ec2LinuxImagePipeline_Feature_NICE_DCV,
			"RETAIN_SSM_AGENT": Ec2LinuxImagePipeline_Feature_RETAIN_SSM_AGENT,
			"STIG": Ec2LinuxImagePipeline_Feature_STIG,
			"SCAP": Ec2LinuxImagePipeline_Feature_SCAP,
		},
	)
	_jsii_.RegisterEnum(
		"@cdklabs/cdk-proserve-lib.patterns.Ec2LinuxImagePipeline.OperatingSystem",
		reflect.TypeOf((*Ec2LinuxImagePipeline_OperatingSystem)(nil)).Elem(),
		map[string]interface{}{
			"RED_HAT_ENTERPRISE_LINUX_8_9": Ec2LinuxImagePipeline_OperatingSystem_RED_HAT_ENTERPRISE_LINUX_8_9,
			"AMAZON_LINUX_2": Ec2LinuxImagePipeline_OperatingSystem_AMAZON_LINUX_2,
			"AMAZON_LINUX_2023": Ec2LinuxImagePipeline_OperatingSystem_AMAZON_LINUX_2023,
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.patterns.Ec2LinuxImagePipelineProps",
		reflect.TypeOf((*Ec2LinuxImagePipelineProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-proserve-lib.patterns.KeycloakService",
		reflect.TypeOf((*KeycloakService)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "adminUser", GoGetter: "AdminUser"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KeycloakService{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.patterns.KeycloakService.ApplicationConfiguration",
		reflect.TypeOf((*KeycloakService_ApplicationConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.patterns.KeycloakService.ClusterConfiguration",
		reflect.TypeOf((*KeycloakService_ClusterConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.patterns.KeycloakService.ClusterRequestCountScalingConfiguration",
		reflect.TypeOf((*KeycloakService_ClusterRequestCountScalingConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.patterns.KeycloakService.ClusterScalingConfiguration",
		reflect.TypeOf((*KeycloakService_ClusterScalingConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-proserve-lib.patterns.KeycloakService.EngineVersion",
		reflect.TypeOf((*KeycloakService_EngineVersion)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "is", GoMethod: "Is"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_KeycloakService_EngineVersion{}
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.patterns.KeycloakService.FabricConfiguration",
		reflect.TypeOf((*KeycloakService_FabricConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.patterns.KeycloakService.HostnameConfiguration",
		reflect.TypeOf((*KeycloakService_HostnameConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.patterns.KeycloakService.InfrastructureConfiguration",
		reflect.TypeOf((*KeycloakService_InfrastructureConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.patterns.KeycloakService.KeycloakProps",
		reflect.TypeOf((*KeycloakService_KeycloakProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.patterns.KeycloakService.ManagementConfiguration",
		reflect.TypeOf((*KeycloakService_ManagementConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.patterns.KeycloakService.RemovalPolicies",
		reflect.TypeOf((*KeycloakService_RemovalPolicies)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.patterns.KeycloakService.ServerlessDatabaseConfiguration",
		reflect.TypeOf((*KeycloakService_ServerlessDatabaseConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.patterns.KeycloakService.TaskSizingConfiguration",
		reflect.TypeOf((*KeycloakService_TaskSizingConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.patterns.KeycloakService.TraditionalDatabaseConfiguration",
		reflect.TypeOf((*KeycloakService_TraditionalDatabaseConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.patterns.KeycloakServiceProps",
		reflect.TypeOf((*KeycloakServiceProps)(nil)).Elem(),
	)
}
