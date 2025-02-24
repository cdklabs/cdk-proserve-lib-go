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
}
