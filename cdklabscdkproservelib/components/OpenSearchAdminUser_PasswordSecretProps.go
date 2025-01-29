package components

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// Properties for the admin user password specific to when the credential is stored in AWS Secrets Manager.
// Experimental.
type OpenSearchAdminUser_PasswordSecretProps struct {
	// Reference to the AWS Secrets Manager secret that contains the admin credential.
	// Experimental.
	Secret awssecretsmanager.ISecret `field:"required" json:"secret" yaml:"secret"`
	// Optional encryption key that protects the secret.
	// Experimental.
	Encryption awskms.IKey `field:"optional" json:"encryption" yaml:"encryption"`
}

