package aspects

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/types"
)

// Properties for configuring the RDS Oracle MultiTenant Aspect.
// Experimental.
type RdsOracleMultiTenantProps struct {
	// Optional KMS key for encrypting Lambda environment variables and CloudWatch log group.
	//
	// If not provided, AWS managed keys will be used for encryption.
	// The Lambda function will be granted encrypt/decrypt permissions on this key.
	// Default: - AWS managed keys are used.
	//
	// Experimental.
	Encryption awskms.IKey `field:"optional" json:"encryption" yaml:"encryption"`
	// Optional Lambda configuration settings for the custom resource handler.
	//
	// Allows customization of VPC settings, security groups, log retention, and other
	// Lambda function properties. Useful when the RDS instance is in a private VPC
	// or when specific networking requirements exist.
	// See: {@link LambdaConfiguration } for available options.
	//
	// Default: - Lambda function uses default settings with no VPC configuration.
	//
	// Experimental.
	LambdaConfiguration *types.LambdaConfiguration `field:"optional" json:"lambdaConfiguration" yaml:"lambdaConfiguration"`
}

