package aspects


// Experimental.
type SecurityCompliance_Suppressions struct {
	// Suppressions to add for CDK Nag on CDK generated policies.
	//
	// If enabled
	// this will add a stack suppression for `AwsSolutions-IAM5` on the actions
	// that CDK commonly generates when using `.grant(...)` methods.
	// Experimental.
	CdkCommonGrants *string `field:"optional" json:"cdkCommonGrants" yaml:"cdkCommonGrants"`
	// Suppressions to add for CDK Nag on CDK generated resources.
	//
	// If enabled
	// this will suppress `AwsSolutions-IAM5` on the policies that are
	// created by CDK Generated Lambda functions, as well as other CDK generated
	// resources such as Log Groups and Step Functions that support CDK
	// generated custom resources. This only applies to resources that are
	// created by the underlying CDK.
	//
	// - Policy suppression: AwsSolutions-IAM5
	// - Log Group suppression: NIST.800.53.R5-CloudWatchLogGroupEncrypted
	// - Step Function suppression: AwsSolutions-SF1.
	// Experimental.
	CdkGeneratedResources *string `field:"optional" json:"cdkGeneratedResources" yaml:"cdkGeneratedResources"`
	// Adds a stack suppression for `NIST.800.53.R5-IAMNoInlinePolicy`. CDK commonly uses inline policies when adding permissions.
	// Experimental.
	IamNoInlinePolicies *string `field:"optional" json:"iamNoInlinePolicies" yaml:"iamNoInlinePolicies"`
	// Adds a stack suppression for `NIST.800.53.R5-LambdaDLQ`.
	// Experimental.
	LambdaNoDlq *string `field:"optional" json:"lambdaNoDlq" yaml:"lambdaNoDlq"`
	// Adds a stack suppression for `NIST.800.53.R5-LambdaInsideVPC`.
	// Experimental.
	LambdaNotInVpc *string `field:"optional" json:"lambdaNotInVpc" yaml:"lambdaNotInVpc"`
	// Adds a stack suppression for `NIST.800.53.R5-S3BucketReplicationEnabled`.
	// Experimental.
	S3BucketReplication *string `field:"optional" json:"s3BucketReplication" yaml:"s3BucketReplication"`
}

