package components

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/components/internal"
)

// Manages provisioning a DynamoDB table.
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

