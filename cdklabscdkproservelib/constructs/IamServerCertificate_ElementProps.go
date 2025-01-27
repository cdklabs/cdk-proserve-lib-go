package constructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Properties for a server certificate element regardless of where it is stored.
// Experimental.
type IamServerCertificate_ElementProps struct {
	// Optional encryption key that protects the secret.
	// Experimental.
	Encryption awskms.IKey `field:"optional" json:"encryption" yaml:"encryption"`
}

