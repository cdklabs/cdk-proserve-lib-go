package components

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsopensearchservice"
	"github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/types"
)

// Properties for the OpenSearchProvisionDomain construct.
// Experimental.
type OpenSearchProvisionDomainProps struct {
	// Amazon OpenSearch Service domain to provision.
	// Experimental.
	Domain awsopensearchservice.IDomain `field:"required" json:"domain" yaml:"domain"`
	// AWS IAM Role that is configured as an administrative user of the Amazon OpenSearch Service domain.
	// Experimental.
	DomainAdmin awsiam.IRole `field:"required" json:"domainAdmin" yaml:"domainAdmin"`
	// Type of the managed Amazon OpenSearch Service domain.
	// Experimental.
	DomainType *string `field:"required" json:"domainType" yaml:"domainType"`
	// Path on the local disk to the files that will be used to provision the Amazon OpenSearch Service domain.
	// Experimental.
	ProvisioningConfigurationPath *string `field:"required" json:"provisioningConfigurationPath" yaml:"provisioningConfigurationPath"`
	// If specified, defines which destructive operations the Custom Resource will handle.
	//
	// If this is not specified, then the Custom Resource will only modify the domain on a CREATE call from AWS
	// CloudFormation.
	// Experimental.
	AllowDestructiveOperations types.DestructiveOperation `field:"optional" json:"allowDestructiveOperations" yaml:"allowDestructiveOperations"`
	// Additional settings to configure on the Amazon OpenSearch Service domain cluster itself.
	//
	// These settings will be sent as a JSON request to the /_cluster/settings API on OpenSearch.
	//
	// Additional details can be found
	// [here](https://docs.opensearch.org/docs/latest/api-reference/cluster-api/cluster-settings/)
	// Experimental.
	ClusterSettings *map[string]interface{} `field:"optional" json:"clusterSettings" yaml:"clusterSettings"`
	// Allows mapping of a role in an Amazon OpenSearch Service domain to multiple backend roles (like IAM Role ARNs, LDAP DNs, etc.).
	//
	// The key is the role name in OpenSearch and the value is a list of entities to map to that role (e.g. local
	// database users or AWS IAM role ARNs).
	// Experimental.
	DynamicRoleMappings *map[string]*[]*string `field:"optional" json:"dynamicRoleMappings" yaml:"dynamicRoleMappings"`
	// Encryption key for protecting the framework resources.
	// Experimental.
	Encryption awskms.IKey `field:"optional" json:"encryption" yaml:"encryption"`
	// Optional Lambda configuration settings.
	// Experimental.
	LambdaConfiguration *types.LambdaConfiguration `field:"optional" json:"lambdaConfiguration" yaml:"lambdaConfiguration"`
}

