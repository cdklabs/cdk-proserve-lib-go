package components

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/components/internal"
)

// Controls the contents of an Amazon DynamoDB table from Infrastructure as Code.
//
// This construct uses information about the key attributes of a table and a list of rows to populate the table upon
// creation, update the table upon update, and remove entries from the table upon delete.
//
// WARNING: This construct should only be used with tables that are created and fully managed via IaC. While the
// the construct will only manage rows within the table that it is aware of, there is no way to detect drift and thus
// it is possible to cause data loss within the table if it is managed outside of IaC as well.
//
// The construct also handles encryption for the framework resources using either a provided KMS key or an
// AWS managed key.
//
// Example:
//   import { DynamoDbProvisionTable } from '@cdklabs/cdk-proserve-lib/constructs';
//   import { Table } from 'aws-cdk-lib/aws-dynamodb';
//   import { Key } from 'aws-cdk-lib/aws-kms';
//
//   interface TableRow {
//       readonly uid: number;
//       readonly isActive: boolean;
//   }
//
//   const partitionKey: keyof TableRow = 'uid';
//
//   const rows: TableRow[] = [
//       {
//           isActive: true,
//           uid: 1
//       },
//       {
//           isActive: true,
//           uid: 2
//       },
//       {
//           isActive: false,
//           uid: 3
//       }
//   ];
//
//   const tableArn = 'arn:aws:dynamodb:us-west-1:111111111111:table/sample';
//   const table = Table.fromTableArn(this, 'Table', tableArn);
//
//   const keyArn = 'arn:aws:kms:us-east-1:111111111111:key/sample-key-id';
//   const key = Key.fromKeyArn(this, 'Encryption', keyArn);
//
//   new DynamoDbProvisionTable(this, 'ProvisionTable', {
//       items: rows,
//       table: {
//           partitionKeyName: partitionKey,
//           resource: table,
//           encryption: key
//       }
//   });
//
//   }
//
// Experimental.
type DynamoDbProvisionTable interface {
	constructs.Construct
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DynamoDbProvisionTable
type jsiiProxy_DynamoDbProvisionTable struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_DynamoDbProvisionTable) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Provisions an existing DynamoDB Table with user-specified data.
// Experimental.
func NewDynamoDbProvisionTable(scope constructs.Construct, id *string, props *DynamoDbProvisionTableProps) DynamoDbProvisionTable {
	_init_.Initialize()

	if err := validateNewDynamoDbProvisionTableParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_DynamoDbProvisionTable{}

	_jsii_.Create(
		"@cdklabs/cdk-proserve-lib.constructs.DynamoDbProvisionTable",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Provisions an existing DynamoDB Table with user-specified data.
// Experimental.
func NewDynamoDbProvisionTable_Override(d DynamoDbProvisionTable, scope constructs.Construct, id *string, props *DynamoDbProvisionTableProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-proserve-lib.constructs.DynamoDbProvisionTable",
		[]interface{}{scope, id, props},
		d,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func DynamoDbProvisionTable_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDynamoDbProvisionTable_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-proserve-lib.constructs.DynamoDbProvisionTable",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoDbProvisionTable) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

