package constructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Properties for the admin user password regardless of where it is stored.
// Experimental.
type OpenSearchAdminUser_PasswordCommonProps struct {
	// Optional encryption key that protects the secret.
	// Experimental.
	Encryption awskms.IKey `field:"optional" json:"encryption" yaml:"encryption"`
}

