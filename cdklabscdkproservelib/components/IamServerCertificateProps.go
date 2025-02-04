package components

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/types"
)

// Properties for the IamServerCertificate construct.
// Experimental.
type IamServerCertificateProps struct {
	// AWS Systems Manager parameter or AWS Secrets Manager secret which contains the public certificate.
	// Experimental.
	Certificate interface{} `field:"required" json:"certificate" yaml:"certificate"`
	// Prefix to prepend to the AWS IAM Server Certificate name.
	// Experimental.
	Prefix *string `field:"required" json:"prefix" yaml:"prefix"`
	// AWS Systems Manager parameter or AWS Secrets Manager secret which contains the private key.
	// Experimental.
	PrivateKey interface{} `field:"required" json:"privateKey" yaml:"privateKey"`
	// Encryption key for protecting the framework resources.
	// Experimental.
	Encryption awskms.IKey `field:"optional" json:"encryption" yaml:"encryption"`
	// Optional Lambda configuration settings.
	// Experimental.
	LambdaConfiguration *types.LambdaConfiguration `field:"optional" json:"lambdaConfiguration" yaml:"lambdaConfiguration"`
}

