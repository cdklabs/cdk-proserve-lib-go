package constructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm"
)

// Properties for the admin user password specific to when the credential is stored in AWS Systems Manager Parameter Store.
// Experimental.
type OpenSearchAdminUser_PasswordParameterProps struct {
	// Optional encryption key that protects the secret.
	// Experimental.
	Encryption awskms.IKey `field:"optional" json:"encryption" yaml:"encryption"`
	// Reference to the AWS Systems Manager Parameter Store parameter that contains the admin credential.
	// Experimental.
	Parameter awsssm.IParameter `field:"required" json:"parameter" yaml:"parameter"`
}

