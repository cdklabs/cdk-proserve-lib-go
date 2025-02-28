package components

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Information about the table to provision.
// Experimental.
type DynamoDbProvisionTable_TableProps struct {
	// Name of the partition key for the table.
	// Experimental.
	PartitionKeyName *string `field:"required" json:"partitionKeyName" yaml:"partitionKeyName"`
	// CDK representation of the table itself.
	// Experimental.
	Resource awsdynamodb.ITable `field:"required" json:"resource" yaml:"resource"`
	// Optional existing encryption key associated with the table.
	// Experimental.
	Encryption awskms.IKey `field:"optional" json:"encryption" yaml:"encryption"`
	// Name of the sort key for the table if applicable.
	// Experimental.
	SortKeyName *string `field:"optional" json:"sortKeyName" yaml:"sortKeyName"`
}

