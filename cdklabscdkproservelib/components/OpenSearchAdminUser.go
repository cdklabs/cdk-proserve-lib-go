package components

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/components/internal"
)

// Manages an admin user for an Amazon OpenSearch domain.
//
// This construct creates a Lambda-backed custom resource that adds an admin user to the specified OpenSearch domain.
// It uses the provided SSM parameter for the username, a provided SSM parameter or Secrets Manager secret for the
// password, and sets up the necessary IAM permissions for the Lambda function to interact with the OpenSearch domain
// and SSM parameter(s) and/or secret.
//
// The construct also handles encryption for the Lambda function's environment variables and dead letter queue,
// using either a provided KMS key or an AWS managed key.
//
// Example:
//   import { Key } from 'aws-cdk-lib/aws-kms';
//   import { Domain } from 'aws-cdk-lib/aws-opensearchservice';
//   import { Secret } from 'aws-cdk-lib/aws-secretsmanager';
//   import { OpenSearchAdminUser } from '@cdklabs/cdk-proserve-lib/constructs';
//
//   const keyArn = 'arn:aws:kms:us-east-1:111111111111:key/sample-key-id';
//   const key = Key.fromKeyArn(this, 'Encryption', keyArn);
//
//   const adminCredential = StringParameter.fromSecureStringParameterAttributes(this, 'AdminCredential', {
//        parameterName: 'sample-parameter',
//        encryptionKey: key
//   });
//
//   const domainKeyArn = 'arn:aws:kms:us-east-1:111111111111:key/sample-domain-key-id';
//   const domainKey = Key.fromKeyArn(this, 'DomainEncryption', domainKeyArn);
//   const domain = Domain.fromDomainEndpoint(this, 'Domain', 'vpc-testdomain.us-east-1.es.amazonaws.com');
//
//   const adminUser = new OpenSearchAdminUser(this, 'AdminUser', {
//        credential: {
//            secret: adminCredential,
//            encryption: key
//        },
//        domain: domain,
//        domainKey: domainKey,
//        username: 'admin'
//   });
//
// Experimental.
type OpenSearchAdminUser interface {
	constructs.Construct
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OpenSearchAdminUser
type jsiiProxy_OpenSearchAdminUser struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_OpenSearchAdminUser) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Constructor.
// Experimental.
func NewOpenSearchAdminUser(scope constructs.Construct, id *string, props *OpenSearchAdminUserProps) OpenSearchAdminUser {
	_init_.Initialize()

	if err := validateNewOpenSearchAdminUserParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_OpenSearchAdminUser{}

	_jsii_.Create(
		"@cdklabs/cdk-proserve-lib.constructs.OpenSearchAdminUser",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructor.
// Experimental.
func NewOpenSearchAdminUser_Override(o OpenSearchAdminUser, scope constructs.Construct, id *string, props *OpenSearchAdminUserProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-proserve-lib.constructs.OpenSearchAdminUser",
		[]interface{}{scope, id, props},
		o,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func OpenSearchAdminUser_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOpenSearchAdminUser_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-proserve-lib.constructs.OpenSearchAdminUser",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenSearchAdminUser) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

