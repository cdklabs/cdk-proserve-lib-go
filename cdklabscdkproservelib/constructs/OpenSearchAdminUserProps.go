package constructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsopensearchservice"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm"
	"github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/interfaces"
)

// Properties for the OpenSearchAdminUser construct.
// Experimental.
type OpenSearchAdminUserProps struct {
	// The SSM parameter or Secret containing the password for the OpenSearch admin user.
	// Experimental.
	Credential interface{} `field:"required" json:"credential" yaml:"credential"`
	// The OpenSearch domain to which the admin user will be added.
	// Experimental.
	Domain awsopensearchservice.IDomain `field:"required" json:"domain" yaml:"domain"`
	// The SSM parameter containing the username for the OpenSearch admin user.
	// Experimental.
	Username awsssm.IParameter `field:"required" json:"username" yaml:"username"`
	// Optional.
	//
	// The KMS key used to encrypt the OpenSearch domain.
	// If provided, the construct will grant the necessary permissions to use this key.
	// Experimental.
	DomainKey awskms.IKey `field:"optional" json:"domainKey" yaml:"domainKey"`
	// Optional.
	//
	// The KMS key used to encrypt the worker resources (e.g., Lambda function environment variables).
	// If provided, this key will be used for encryption; otherwise, an AWS managed key will be used.
	// Experimental.
	Encryption awskms.IKey `field:"optional" json:"encryption" yaml:"encryption"`
	// Optional Lambda configuration settings.
	// Experimental.
	LambdaConfiguration *interfaces.LambdaConfiguration `field:"optional" json:"lambdaConfiguration" yaml:"lambdaConfiguration"`
}

