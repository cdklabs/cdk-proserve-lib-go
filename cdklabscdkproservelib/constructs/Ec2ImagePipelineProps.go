package constructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsimagebuilder"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/interfaces"
)

// Properties for EC2 Image Pipeline, extending the base properties.
// Experimental.
type Ec2ImagePipelineProps struct {
	// Version of the image pipeline.
	//
	// This must be updated if you make
	// underlying changes to the pipeline configuration.
	// Experimental.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Configuration options for the build process.
	// Experimental.
	BuildConfiguration *Ec2ImagePipeline_BuildConfigurationProps `field:"optional" json:"buildConfiguration" yaml:"buildConfiguration"`
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
	LambdaConfiguration *interfaces.LambdaConfiguration `field:"optional" json:"lambdaConfiguration" yaml:"lambdaConfiguration"`
	// VPC configuration for the image pipeline.
	// Experimental.
	VpcConfiguration *Ec2ImagePipeline_VpcConfigurationProps `field:"optional" json:"vpcConfiguration" yaml:"vpcConfiguration"`
	// Block device mappings for the image.
	// Experimental.
	BlockDeviceMappings *[]*awsimagebuilder.CfnImageRecipe_InstanceBlockDeviceMappingProperty `field:"optional" json:"blockDeviceMappings" yaml:"blockDeviceMappings"`
	// Components to be included in the image pipeline.
	//
	// Can be either custom Ec2ImagePipeline.Component or AWS CDK imagebuilder.CfnComponent.
	// Experimental.
	Components *[]interface{} `field:"optional" json:"components" yaml:"components"`
	// The machine image to use as a base for the pipeline.
	// Default: AmazonLinux2023.
	//
	// Experimental.
	MachineImage awsec2.IMachineImage `field:"optional" json:"machineImage" yaml:"machineImage"`
}

