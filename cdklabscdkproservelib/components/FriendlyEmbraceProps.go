package components

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/types"
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
	LambdaConfiguration *types.LambdaConfiguration `field:"optional" json:"lambdaConfiguration" yaml:"lambdaConfiguration"`
	// Manually provide specific read-only permissions for resources in your CloudFormation templates to support instead of using the AWS managed policy [ReadOnlyAccess](https://docs.aws.amazon.com/aws-managed-policy/latest/reference/ReadOnlyAccess.html).
	//
	// This can be useful in environments where the caller wants to maintain tight control over the permissions granted
	// to the custom resource worker.
	// Experimental.
	ManualReadPermissions *[]awsiam.PolicyStatement `field:"optional" json:"manualReadPermissions" yaml:"manualReadPermissions"`
}

