package components

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/types"
)

// Properties for the DynamoDbProvisionTable construct.
// Experimental.
type DynamoDbProvisionTableProps struct {
	// Items to provision within the DynamoDB table.
	// Experimental.
	Items *[]*map[string]interface{} `field:"required" json:"items" yaml:"items"`
	// Table to provision.
	// Experimental.
	Table *DynamoDbProvisionTable_TableProps `field:"required" json:"table" yaml:"table"`
	// Encryption key for protecting the framework resources.
	// Experimental.
	Encryption awskms.IKey `field:"optional" json:"encryption" yaml:"encryption"`
	// Optional Lambda configuration settings.
	// Experimental.
	LambdaConfiguration *types.LambdaConfiguration `field:"optional" json:"lambdaConfiguration" yaml:"lambdaConfiguration"`
}

