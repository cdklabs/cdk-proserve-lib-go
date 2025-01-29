package components

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm"
)

// Properties for a server certificate element when it is stored in AWS Systems Manager Parameter Store.
// Experimental.
type IamServerCertificate_ParameterProps struct {
	// Reference to the AWS Systems Manager Parameter Store parameter that contains the data.
	// Experimental.
	Parameter awsssm.IParameter `field:"required" json:"parameter" yaml:"parameter"`
	// Optional encryption key that protects the secret.
	// Experimental.
	Encryption awskms.IKey `field:"optional" json:"encryption" yaml:"encryption"`
}

