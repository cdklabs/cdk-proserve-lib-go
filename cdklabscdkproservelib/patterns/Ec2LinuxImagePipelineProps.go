package patterns

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsimagebuilder"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/components"
	"github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/types"
)

// Properties for creating a Linux STIG Image Pipeline.
// Experimental.
type Ec2LinuxImagePipelineProps struct {
	// Version of the image pipeline.
	//
	// This must be updated if you make
	// underlying changes to the pipeline configuration.
	// Experimental.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Configuration options for the build process.
	// Experimental.
	BuildConfiguration *components.Ec2ImagePipeline_BuildConfigurationProps `field:"optional" json:"buildConfiguration" yaml:"buildConfiguration"`
	// Description of the image pipeline.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// KMS key for encryption.
	// Experimental.
	Encryption awskms.IKey `field:"optional" json:"encryption" yaml:"encryption"`
	// Instance types for the Image Builder Pipeline.
	//
	// Default: [t3.medium]
	// Experimental.
	InstanceTypes *[]*string `field:"optional" json:"instanceTypes" yaml:"instanceTypes"`
	// Optional Lambda configuration settings.
	// Experimental.
	LambdaConfiguration *types.LambdaConfiguration `field:"optional" json:"lambdaConfiguration" yaml:"lambdaConfiguration"`
	// VPC configuration for the image pipeline.
	// Experimental.
	VpcConfiguration *components.Ec2ImagePipeline_VpcConfigurationProps `field:"optional" json:"vpcConfiguration" yaml:"vpcConfiguration"`
	// Additional components to install in the image.
	//
	// These components will be added after the default Linux components.
	// Use this to add custom components beyond the predefined features.
	// Experimental.
	ExtraComponents *[]interface{} `field:"optional" json:"extraComponents" yaml:"extraComponents"`
	// Additional EBS volume mappings to add to the image.
	//
	// These volumes will be added in addition to the root volume.
	// Use this to specify additional storage volumes that should be included in the AMI.
	// Experimental.
	ExtraDeviceMappings *[]*awsimagebuilder.CfnImageRecipe_InstanceBlockDeviceMappingProperty `field:"optional" json:"extraDeviceMappings" yaml:"extraDeviceMappings"`
	// A list of features to install on the image.
	//
	// Features are predefined sets of components and configurations.
	// Default: [AWS_CLI, RETAIN_SSM_AGENT].
	// Default: [Ec2LinuxImagePipeline.Feature.AWS_CLI, Ec2LinuxImagePipeline.Feature.RETAIN_SSM_AGENT]
	//
	// Experimental.
	Features *[]Ec2LinuxImagePipeline_Feature `field:"optional" json:"features" yaml:"features"`
	// The operating system to use for the image pipeline.
	//
	// Specifies which operating system version to use as the base image.
	// Default: AMAZON_LINUX_2023.
	// Default: Ec2LinuxImagePipeline.OperatingSystem.AMAZON_LINUX_2023
	//
	// Experimental.
	OperatingSystem Ec2LinuxImagePipeline_OperatingSystem `field:"optional" json:"operatingSystem" yaml:"operatingSystem"`
	// Size for the root volume in GB.
	//
	// Default: 10 GB.
	// Default: 10.
	//
	// Experimental.
	RootVolumeSize *float64 `field:"optional" json:"rootVolumeSize" yaml:"rootVolumeSize"`
}

