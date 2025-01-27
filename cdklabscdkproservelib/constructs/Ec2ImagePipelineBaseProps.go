package constructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/interfaces"
)

// Base properties for EC2 Image Pipeline configuration.
// Experimental.
type Ec2ImagePipelineBaseProps struct {
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
}

