package components

import (
	"github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/interfaces"
)

// Properties for the Ec2ImageBuilderGetImage construct.
// Experimental.
type Ec2ImageBuilderGetImageProps struct {
	// The ARN of the EC2 Image Builder image build version.
	// Experimental.
	ImageBuildVersionArn *string `field:"required" json:"imageBuildVersionArn" yaml:"imageBuildVersionArn"`
	// Optional Lambda configuration settings.
	// Experimental.
	LambdaConfiguration *interfaces.AwsCustomResourceLambdaConfiguration `field:"optional" json:"lambdaConfiguration" yaml:"lambdaConfiguration"`
}

