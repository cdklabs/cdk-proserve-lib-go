package components

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdklabs/cdk-proserve-lib.constructs.DynamoDbProvisionTable",
		reflect.TypeOf((*DynamoDbProvisionTable)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DynamoDbProvisionTable{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.constructs.DynamoDbProvisionTable.TableProps",
		reflect.TypeOf((*DynamoDbProvisionTable_TableProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.constructs.DynamoDbProvisionTableProps",
		reflect.TypeOf((*DynamoDbProvisionTableProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-proserve-lib.constructs.Ec2ImageBuilderGetImage",
		reflect.TypeOf((*Ec2ImageBuilderGetImage)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "ami", GoGetter: "Ami"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2ImageBuilderGetImage{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.constructs.Ec2ImageBuilderGetImageProps",
		reflect.TypeOf((*Ec2ImageBuilderGetImageProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-proserve-lib.constructs.Ec2ImageBuilderStart",
		reflect.TypeOf((*Ec2ImageBuilderStart)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "imageBuildVersionArn", GoGetter: "ImageBuildVersionArn"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2ImageBuilderStart{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.constructs.Ec2ImageBuilderStart.WaitForCompletionProps",
		reflect.TypeOf((*Ec2ImageBuilderStart_WaitForCompletionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.constructs.Ec2ImageBuilderStartProps",
		reflect.TypeOf((*Ec2ImageBuilderStartProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-proserve-lib.constructs.Ec2ImagePipeline",
		reflect.TypeOf((*Ec2ImagePipeline)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "imagePipelineArn", GoGetter: "ImagePipelineArn"},
			_jsii_.MemberProperty{JsiiProperty: "imagePipelineTopic", GoGetter: "ImagePipelineTopic"},
			_jsii_.MemberProperty{JsiiProperty: "latestAmi", GoGetter: "LatestAmi"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2ImagePipeline{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.constructs.Ec2ImagePipeline.BuildConfigurationProps",
		reflect.TypeOf((*Ec2ImagePipeline_BuildConfigurationProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@cdklabs/cdk-proserve-lib.constructs.Ec2ImagePipeline.Component",
		reflect.TypeOf((*Ec2ImagePipeline_Component)(nil)).Elem(),
		map[string]interface{}{
			"AMAZON_CLOUDWATCH_AGENT_LINUX": Ec2ImagePipeline_Component_AMAZON_CLOUDWATCH_AGENT_LINUX,
			"AMAZON_CLOUDWATCH_AGENT_WINDOWS": Ec2ImagePipeline_Component_AMAZON_CLOUDWATCH_AGENT_WINDOWS,
			"AMAZON_CORRETTO_11_APT_GENERIC": Ec2ImagePipeline_Component_AMAZON_CORRETTO_11_APT_GENERIC,
			"AMAZON_CORRETTO_11_HEADLESS": Ec2ImagePipeline_Component_AMAZON_CORRETTO_11_HEADLESS,
			"AMAZON_CORRETTO_11_RPM_GENERIC": Ec2ImagePipeline_Component_AMAZON_CORRETTO_11_RPM_GENERIC,
			"AMAZON_CORRETTO_11_WINDOWS": Ec2ImagePipeline_Component_AMAZON_CORRETTO_11_WINDOWS,
			"AMAZON_CORRETTO_11": Ec2ImagePipeline_Component_AMAZON_CORRETTO_11,
			"AMAZON_CORRETTO_17_HEADLESS": Ec2ImagePipeline_Component_AMAZON_CORRETTO_17_HEADLESS,
			"AMAZON_CORRETTO_17_JDK": Ec2ImagePipeline_Component_AMAZON_CORRETTO_17_JDK,
			"AMAZON_CORRETTO_17_JRE": Ec2ImagePipeline_Component_AMAZON_CORRETTO_17_JRE,
			"AMAZON_CORRETTO_17_WINDOWS": Ec2ImagePipeline_Component_AMAZON_CORRETTO_17_WINDOWS,
			"AMAZON_CORRETTO_21_HEADLESS": Ec2ImagePipeline_Component_AMAZON_CORRETTO_21_HEADLESS,
			"AMAZON_CORRETTO_21_JDK": Ec2ImagePipeline_Component_AMAZON_CORRETTO_21_JDK,
			"AMAZON_CORRETTO_21_JRE": Ec2ImagePipeline_Component_AMAZON_CORRETTO_21_JRE,
			"AMAZON_CORRETTO_21_WINDOWS": Ec2ImagePipeline_Component_AMAZON_CORRETTO_21_WINDOWS,
			"AMAZON_CORRETTO_8_APT_GENERIC": Ec2ImagePipeline_Component_AMAZON_CORRETTO_8_APT_GENERIC,
			"AMAZON_CORRETTO_8_JDK": Ec2ImagePipeline_Component_AMAZON_CORRETTO_8_JDK,
			"AMAZON_CORRETTO_8_JRE": Ec2ImagePipeline_Component_AMAZON_CORRETTO_8_JRE,
			"AMAZON_CORRETTO_8_RPM_GENERIC": Ec2ImagePipeline_Component_AMAZON_CORRETTO_8_RPM_GENERIC,
			"AMAZON_CORRETTO_8_WINDOWS": Ec2ImagePipeline_Component_AMAZON_CORRETTO_8_WINDOWS,
			"AMAZON_KINESIS_AGENT_WINDOWS": Ec2ImagePipeline_Component_AMAZON_KINESIS_AGENT_WINDOWS,
			"ANACONDA_WINDOWS": Ec2ImagePipeline_Component_ANACONDA_WINDOWS,
			"APACHE_TOMCAT_9_LINUX": Ec2ImagePipeline_Component_APACHE_TOMCAT_9_LINUX,
			"APT_REPOSITORY_TEST_LINUX": Ec2ImagePipeline_Component_APT_REPOSITORY_TEST_LINUX,
			"AWS_CLI_VERSION_2_LINUX": Ec2ImagePipeline_Component_AWS_CLI_VERSION_2_LINUX,
			"AWS_CLI_VERSION_2_WINDOWS": Ec2ImagePipeline_Component_AWS_CLI_VERSION_2_WINDOWS,
			"AWS_CODEDEPLOY_AGENT_LINUX": Ec2ImagePipeline_Component_AWS_CODEDEPLOY_AGENT_LINUX,
			"AWS_CODEDEPLOY_AGENT_WINDOWS": Ec2ImagePipeline_Component_AWS_CODEDEPLOY_AGENT_WINDOWS,
			"AWS_VSS_COMPONENTS_WINDOWS": Ec2ImagePipeline_Component_AWS_VSS_COMPONENTS_WINDOWS,
			"CHOCOLATEY": Ec2ImagePipeline_Component_CHOCOLATEY,
			"CHRONY_TIME_CONFIGURATION_TEST": Ec2ImagePipeline_Component_CHRONY_TIME_CONFIGURATION_TEST,
			"DCV_SERVER_LINUX": Ec2ImagePipeline_Component_DCV_SERVER_LINUX,
			"DCV_SERVER_WINDOWS": Ec2ImagePipeline_Component_DCV_SERVER_WINDOWS,
			"DISTRIBUTOR_PACKAGE_WINDOWS": Ec2ImagePipeline_Component_DISTRIBUTOR_PACKAGE_WINDOWS,
			"DOCKER_CE_CENTOS": Ec2ImagePipeline_Component_DOCKER_CE_CENTOS,
			"DOCKER_CE_LINUX": Ec2ImagePipeline_Component_DOCKER_CE_LINUX,
			"DOCKER_CE_UBUNTU": Ec2ImagePipeline_Component_DOCKER_CE_UBUNTU,
			"DOTNET_DESKTOP_RUNTIME_LTS_WINDOWS": Ec2ImagePipeline_Component_DOTNET_DESKTOP_RUNTIME_LTS_WINDOWS,
			"DOTNET_HOSTING_BUNDLE_LTS_WINDOWS": Ec2ImagePipeline_Component_DOTNET_HOSTING_BUNDLE_LTS_WINDOWS,
			"DOTNET_RUNTIME_LTS_LINUX": Ec2ImagePipeline_Component_DOTNET_RUNTIME_LTS_LINUX,
			"DOTNET_RUNTIME_LTS_WINDOWS": Ec2ImagePipeline_Component_DOTNET_RUNTIME_LTS_WINDOWS,
			"DOTNET_SDK_LTS_LINUX": Ec2ImagePipeline_Component_DOTNET_SDK_LTS_LINUX,
			"DOTNET_SDK_LTS_WINDOWS": Ec2ImagePipeline_Component_DOTNET_SDK_LTS_WINDOWS,
			"EBS_VOLUME_USAGE_TEST_LINUX": Ec2ImagePipeline_Component_EBS_VOLUME_USAGE_TEST_LINUX,
			"EBS_VOLUME_USAGE_TEST_WINDOWS": Ec2ImagePipeline_Component_EBS_VOLUME_USAGE_TEST_WINDOWS,
			"EC2_NETWORK_ROUTE_TEST_WINDOWS": Ec2ImagePipeline_Component_EC2_NETWORK_ROUTE_TEST_WINDOWS,
			"EC2LAUNCH_V2_WINDOWS": Ec2ImagePipeline_Component_EC2LAUNCH_V2_WINDOWS,
			"ECS_OPTIMIZED_AMI_WINDOWS": Ec2ImagePipeline_Component_ECS_OPTIMIZED_AMI_WINDOWS,
			"EKS_OPTIMIZED_AMI_WINDOWS": Ec2ImagePipeline_Component_EKS_OPTIMIZED_AMI_WINDOWS,
			"ENI_ATTACHMENT_TEST_LINUX": Ec2ImagePipeline_Component_ENI_ATTACHMENT_TEST_LINUX,
			"ENI_ATTACHMENT_TEST_WINDOWS": Ec2ImagePipeline_Component_ENI_ATTACHMENT_TEST_WINDOWS,
			"GO_STABLE_LINUX": Ec2ImagePipeline_Component_GO_STABLE_LINUX,
			"GO_STABLE_WINDOWS": Ec2ImagePipeline_Component_GO_STABLE_WINDOWS,
			"HELLO_WORLD_LINUX": Ec2ImagePipeline_Component_HELLO_WORLD_LINUX,
			"HELLO_WORLD_WINDOWS": Ec2ImagePipeline_Component_HELLO_WORLD_WINDOWS,
			"INSPECTOR_TEST_LINUX": Ec2ImagePipeline_Component_INSPECTOR_TEST_LINUX,
			"INSPECTOR_TEST_WINDOWS": Ec2ImagePipeline_Component_INSPECTOR_TEST_WINDOWS,
			"INSTALL_PACKAGE_FROM_REPOSITORY": Ec2ImagePipeline_Component_INSTALL_PACKAGE_FROM_REPOSITORY,
			"MARIADB_LINUX": Ec2ImagePipeline_Component_MARIADB_LINUX,
			"MATE_DE_LINUX": Ec2ImagePipeline_Component_MATE_DE_LINUX,
			"MONO_LINUX": Ec2ImagePipeline_Component_MONO_LINUX,
			"PHP_8_2_LINUX": Ec2ImagePipeline_Component_PHP_8_2_LINUX,
			"POWERSHELL_LTS_LINUX": Ec2ImagePipeline_Component_POWERSHELL_LTS_LINUX,
			"POWERSHELL_LTS_WINDOWS": Ec2ImagePipeline_Component_POWERSHELL_LTS_WINDOWS,
			"POWERSHELL_SNAP": Ec2ImagePipeline_Component_POWERSHELL_SNAP,
			"POWERSHELL_YUM": Ec2ImagePipeline_Component_POWERSHELL_YUM,
			"PUTTY": Ec2ImagePipeline_Component_PUTTY,
			"PYTHON_3_LINUX": Ec2ImagePipeline_Component_PYTHON_3_LINUX,
			"PYTHON_3_WINDOWS": Ec2ImagePipeline_Component_PYTHON_3_WINDOWS,
			"REBOOT_LINUX": Ec2ImagePipeline_Component_REBOOT_LINUX,
			"REBOOT_TEST_LINUX": Ec2ImagePipeline_Component_REBOOT_TEST_LINUX,
			"REBOOT_TEST_WINDOWS": Ec2ImagePipeline_Component_REBOOT_TEST_WINDOWS,
			"REBOOT_WINDOWS": Ec2ImagePipeline_Component_REBOOT_WINDOWS,
			"SCAP_COMPLIANCE_CHECKER_LINUX": Ec2ImagePipeline_Component_SCAP_COMPLIANCE_CHECKER_LINUX,
			"SCAP_COMPLIANCE_CHECKER_WINDOWS": Ec2ImagePipeline_Component_SCAP_COMPLIANCE_CHECKER_WINDOWS,
			"SIMPLE_BOOT_TEST_LINUX": Ec2ImagePipeline_Component_SIMPLE_BOOT_TEST_LINUX,
			"SIMPLE_BOOT_TEST_WINDOWS": Ec2ImagePipeline_Component_SIMPLE_BOOT_TEST_WINDOWS,
			"STIG_BUILD_LINUX_HIGH": Ec2ImagePipeline_Component_STIG_BUILD_LINUX_HIGH,
			"STIG_BUILD_LINUX_LOW": Ec2ImagePipeline_Component_STIG_BUILD_LINUX_LOW,
			"STIG_BUILD_LINUX_MEDIUM": Ec2ImagePipeline_Component_STIG_BUILD_LINUX_MEDIUM,
			"STIG_BUILD_WINDOWS_HIGH": Ec2ImagePipeline_Component_STIG_BUILD_WINDOWS_HIGH,
			"STIG_BUILD_WINDOWS_LOW": Ec2ImagePipeline_Component_STIG_BUILD_WINDOWS_LOW,
			"STIG_BUILD_WINDOWS_MEDIUM": Ec2ImagePipeline_Component_STIG_BUILD_WINDOWS_MEDIUM,
			"UPDATE_LINUX_KERNEL_5": Ec2ImagePipeline_Component_UPDATE_LINUX_KERNEL_5,
			"UPDATE_LINUX_KERNEL_ML": Ec2ImagePipeline_Component_UPDATE_LINUX_KERNEL_ML,
			"UPDATE_LINUX": Ec2ImagePipeline_Component_UPDATE_LINUX,
			"UPDATE_WINDOWS": Ec2ImagePipeline_Component_UPDATE_WINDOWS,
			"VALIDATE_SINGLE_SSH_PUBLIC_KEY_TEST_LINUX": Ec2ImagePipeline_Component_VALIDATE_SINGLE_SSH_PUBLIC_KEY_TEST_LINUX,
			"VALIDATE_SSH_HOST_KEY_GENERATION_LINUX": Ec2ImagePipeline_Component_VALIDATE_SSH_HOST_KEY_GENERATION_LINUX,
			"VALIDATE_SSH_PUBLIC_KEY_LINUX": Ec2ImagePipeline_Component_VALIDATE_SSH_PUBLIC_KEY_LINUX,
			"WINDOWS_ACTIVATION_TEST": Ec2ImagePipeline_Component_WINDOWS_ACTIVATION_TEST,
			"WINDOWS_IS_READY_WITH_PASSWORD_GENERATION_TEST": Ec2ImagePipeline_Component_WINDOWS_IS_READY_WITH_PASSWORD_GENERATION_TEST,
			"WINDOWS_SERVER_IIS": Ec2ImagePipeline_Component_WINDOWS_SERVER_IIS,
			"YUM_REPOSITORY_TEST_LINUX": Ec2ImagePipeline_Component_YUM_REPOSITORY_TEST_LINUX,
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.constructs.Ec2ImagePipeline.VpcConfigurationProps",
		reflect.TypeOf((*Ec2ImagePipeline_VpcConfigurationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.constructs.Ec2ImagePipelineBaseProps",
		reflect.TypeOf((*Ec2ImagePipelineBaseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.constructs.Ec2ImagePipelineProps",
		reflect.TypeOf((*Ec2ImagePipelineProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-proserve-lib.constructs.FriendlyEmbrace",
		reflect.TypeOf((*FriendlyEmbrace)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "onEventHandler", GoGetter: "OnEventHandler"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_FriendlyEmbrace{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.constructs.FriendlyEmbraceProps",
		reflect.TypeOf((*FriendlyEmbraceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-proserve-lib.constructs.IamServerCertificate",
		reflect.TypeOf((*IamServerCertificate)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_IamServerCertificate{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.constructs.IamServerCertificate.ParameterProps",
		reflect.TypeOf((*IamServerCertificate_ParameterProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.constructs.IamServerCertificate.SecretProps",
		reflect.TypeOf((*IamServerCertificate_SecretProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.constructs.IamServerCertificateProps",
		reflect.TypeOf((*IamServerCertificateProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-proserve-lib.constructs.NetworkFirewall",
		reflect.TypeOf((*NetworkFirewall)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "firewall", GoGetter: "Firewall"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_NetworkFirewall{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"@cdklabs/cdk-proserve-lib.constructs.NetworkFirewall.LogType",
		reflect.TypeOf((*NetworkFirewall_LogType)(nil)).Elem(),
		map[string]interface{}{
			"TLS": NetworkFirewall_LogType_TLS,
			"FLOW": NetworkFirewall_LogType_FLOW,
			"ALERT": NetworkFirewall_LogType_ALERT,
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.constructs.NetworkFirewall.LoggingConfiguration",
		reflect.TypeOf((*NetworkFirewall_LoggingConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.constructs.NetworkFirewall.NetworkFirewallVpcRouteProps",
		reflect.TypeOf((*NetworkFirewall_NetworkFirewallVpcRouteProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-proserve-lib.constructs.NetworkFirewallEndpoints",
		reflect.TypeOf((*NetworkFirewallEndpoints)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getEndpointId", GoMethod: "GetEndpointId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_NetworkFirewallEndpoints{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.constructs.NetworkFirewallEndpointsProps",
		reflect.TypeOf((*NetworkFirewallEndpointsProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.constructs.NetworkFirewallProps",
		reflect.TypeOf((*NetworkFirewallProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-proserve-lib.constructs.OpenSearchAdminUser",
		reflect.TypeOf((*OpenSearchAdminUser)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_OpenSearchAdminUser{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.constructs.OpenSearchAdminUser.PasswordParameterProps",
		reflect.TypeOf((*OpenSearchAdminUser_PasswordParameterProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.constructs.OpenSearchAdminUser.PasswordSecretProps",
		reflect.TypeOf((*OpenSearchAdminUser_PasswordSecretProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.constructs.OpenSearchAdminUserProps",
		reflect.TypeOf((*OpenSearchAdminUserProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-proserve-lib.constructs.WebApplicationFirewall",
		reflect.TypeOf((*WebApplicationFirewall)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "associate", GoMethod: "Associate"},
			_jsii_.MemberProperty{JsiiProperty: "logGroup", GoGetter: "LogGroup"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "webAcl", GoGetter: "WebAcl"},
		},
		func() interface{} {
			j := jsiiProxy_WebApplicationFirewall{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"@cdklabs/cdk-proserve-lib.constructs.WebApplicationFirewall.AwsManagedRuleGroup",
		reflect.TypeOf((*WebApplicationFirewall_AwsManagedRuleGroup)(nil)).Elem(),
		map[string]interface{}{
			"COMMON_RULE_SET": WebApplicationFirewall_AwsManagedRuleGroup_COMMON_RULE_SET,
			"ADMIN_PROTECTION_RULE_SET": WebApplicationFirewall_AwsManagedRuleGroup_ADMIN_PROTECTION_RULE_SET,
			"KNOWN_BAD_INPUTS_RULE_SET": WebApplicationFirewall_AwsManagedRuleGroup_KNOWN_BAD_INPUTS_RULE_SET,
			"SQL_DATABASE_RULE_SET": WebApplicationFirewall_AwsManagedRuleGroup_SQL_DATABASE_RULE_SET,
			"LINUX_RULE_SET": WebApplicationFirewall_AwsManagedRuleGroup_LINUX_RULE_SET,
			"UNIX_RULE_SET": WebApplicationFirewall_AwsManagedRuleGroup_UNIX_RULE_SET,
			"WINDOWS_RULE_SET": WebApplicationFirewall_AwsManagedRuleGroup_WINDOWS_RULE_SET,
			"PHP_RULE_SET": WebApplicationFirewall_AwsManagedRuleGroup_PHP_RULE_SET,
			"WORD_PRESS_RULE_SET": WebApplicationFirewall_AwsManagedRuleGroup_WORD_PRESS_RULE_SET,
			"AMAZON_IP_REPUTATION_LIST": WebApplicationFirewall_AwsManagedRuleGroup_AMAZON_IP_REPUTATION_LIST,
			"ANONYMOUS_IP_LIST": WebApplicationFirewall_AwsManagedRuleGroup_ANONYMOUS_IP_LIST,
			"BOT_CONTROL_RULE_SET": WebApplicationFirewall_AwsManagedRuleGroup_BOT_CONTROL_RULE_SET,
			"ATP_RULE_SET": WebApplicationFirewall_AwsManagedRuleGroup_ATP_RULE_SET,
			"ACFP_RULE_SET": WebApplicationFirewall_AwsManagedRuleGroup_ACFP_RULE_SET,
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.constructs.WebApplicationFirewall.AwsManagedRuleGroupConfig",
		reflect.TypeOf((*WebApplicationFirewall_AwsManagedRuleGroupConfig)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@cdklabs/cdk-proserve-lib.constructs.WebApplicationFirewall.OverrideAction",
		reflect.TypeOf((*WebApplicationFirewall_OverrideAction)(nil)).Elem(),
		map[string]interface{}{
			"ALLOW": WebApplicationFirewall_OverrideAction_ALLOW,
			"BLOCK": WebApplicationFirewall_OverrideAction_BLOCK,
			"COUNT": WebApplicationFirewall_OverrideAction_COUNT,
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.constructs.WebApplicationFirewall.OverrideConfig",
		reflect.TypeOf((*WebApplicationFirewall_OverrideConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.constructs.WebApplicationFirewall.WebApplicationFirewallLoggingConfig",
		reflect.TypeOf((*WebApplicationFirewall_WebApplicationFirewallLoggingConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-proserve-lib.constructs.WebApplicationFirewallProps",
		reflect.TypeOf((*WebApplicationFirewallProps)(nil)).Elem(),
	)
}
