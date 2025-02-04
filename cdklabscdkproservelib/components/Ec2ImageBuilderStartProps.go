package components

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/types"
)

// Properties for the EC2 Image Builder Start custom resource.
// Experimental.
type Ec2ImageBuilderStartProps struct {
	// The ARN of the Image Builder pipeline to start.
	// Experimental.
	PipelineArn *string `field:"required" json:"pipelineArn" yaml:"pipelineArn"`
	// Optional KMS Encryption Key to use for encrypting resources.
	// Experimental.
	Encryption awskms.IKey `field:"optional" json:"encryption" yaml:"encryption"`
	// An optional user-generated hash value that will determine if the construct will start the build pipeline.
	//
	// If this is not set, the pipeline
	// will only start once on initial deployment. By setting this, you can for
	// example start a new build if your build instructions have changed and
	// then wait for the pipeline to complete again.
	//
	// This hash should be a short
	// string, ideally ~7 characters or less. It will be set as the Physical ID
	// of the Custom Resource and also used to append to Waiter function
	// Physical IDs.
	// Experimental.
	Hash *string `field:"optional" json:"hash" yaml:"hash"`
	// Optional Lambda configuration settings.
	// Experimental.
	LambdaConfiguration *types.LambdaConfiguration `field:"optional" json:"lambdaConfiguration" yaml:"lambdaConfiguration"`
	// Set these properties to wait for the Image Build to complete.
	//
	// This is
	// useful if you need the AMI before your next infrastructure step.
	// Experimental.
	WaitForCompletion *Ec2ImageBuilderStart_WaitForCompletionProps `field:"optional" json:"waitForCompletion" yaml:"waitForCompletion"`
}

