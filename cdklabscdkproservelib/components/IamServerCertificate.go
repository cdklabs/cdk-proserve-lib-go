package components

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/components/internal"
)

// A construct that creates a Custom Resource to manage an AWS Identity and Access Management Server Certificate for use in regions/partitions where AWS Certificate Manager is not available.
//
// This construct allows you to create an IAM Server Certificate using a certificate and private key stored in either
// AWS Systems Manager Parameter Store or AWS Secrets Manager. It uses a Custom Resource to manage the lifecycle of the
// server certificate.
//
// The construct also handles encryption for the framework resources using either a provided KMS key or an
// AWS managed key.
//
// Example:
//   import { Key } from 'aws-cdk-lib/aws-kms';
//   import { Secret } from 'aws-cdk-lib/aws-secretsmanager';
//   import { StringParameter } from 'aws-cdk-lib/aws-ssm';
//   import { IamServerCertificate } from '@cdklabs/cdk-proserve-lib/constructs';
//
//   const keyArn = 'arn:aws:kms:us-east-1:111111111111:key/sample-key-id';
//   const key = Key.fromKeyArn(this, 'Encryption', keyArn);
//
//   const certificateData = StringParameter.fromSecureStringParameterAttributes(this, 'CertificateData', {
//        parameterName: 'sample-parameter',
//        encryptionKey: key
//   });
//
//   const privateKeyData = Secret.fromSecretAttributes(this, 'PrivateKeySecret', {
//        encryptionKey: key,
//        secretCompleteArn: 'arn:aws:secretsmanager:us-east-1:111111111111:secret:PrivateKeySecret-aBc123'
//   });
//
//   const certificate = new IamServerCertificate(this, 'ServerCertificate', {
//        certificate: {
//            parameter: certificateData,
//            encryption: key
//        },
//        privateKey: {
//            secret: privateKeyData,
//            encryption: key
//        },
//        prefix: 'myapp'
//   });
//
// Experimental.
type IamServerCertificate interface {
	constructs.Construct
	// ARN for the created AWS IAM Server Certificate.
	// Experimental.
	Arn() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IamServerCertificate
type jsiiProxy_IamServerCertificate struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_IamServerCertificate) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamServerCertificate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Creates a new AWS IAM Server Certificate.
// Experimental.
func NewIamServerCertificate(scope constructs.Construct, id *string, props *IamServerCertificateProps) IamServerCertificate {
	_init_.Initialize()

	if err := validateNewIamServerCertificateParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_IamServerCertificate{}

	_jsii_.Create(
		"@cdklabs/cdk-proserve-lib.constructs.IamServerCertificate",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Creates a new AWS IAM Server Certificate.
// Experimental.
func NewIamServerCertificate_Override(i IamServerCertificate, scope constructs.Construct, id *string, props *IamServerCertificateProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-proserve-lib.constructs.IamServerCertificate",
		[]interface{}{scope, id, props},
		i,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func IamServerCertificate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIamServerCertificate_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-proserve-lib.constructs.IamServerCertificate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamServerCertificate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

