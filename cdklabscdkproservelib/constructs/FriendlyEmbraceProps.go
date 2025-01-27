package constructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/interfaces"
)

// Input metadata for the custom resource.
// Experimental.
type FriendlyEmbraceProps struct {
	// Optional S3 Bucket configuration settings.
	// Experimental.
	BucketConfiguration *awss3.BucketProps `field:"optional" json:"bucketConfiguration" yaml:"bucketConfiguration"`
	// Encryption key for protecting the Lambda environment.
	// Experimental.
	Encryption awskms.IKey `field:"optional" json:"encryption" yaml:"encryption"`
	// Whether or not stacks in error state should be fatal to CR completion.
	// Experimental.
	IgnoreInvalidStates *bool `field:"optional" json:"ignoreInvalidStates" yaml:"ignoreInvalidStates"`
	// Optional Lambda configuration settings.
	// Experimental.
	LambdaConfiguration *interfaces.LambdaConfiguration `field:"optional" json:"lambdaConfiguration" yaml:"lambdaConfiguration"`
}

