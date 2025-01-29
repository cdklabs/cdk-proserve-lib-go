package components

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/components/internal"
)

// Retrieves an EC2 Image Builder image build version.
//
// This is useful for retrieving the AMI ID of an image that was built by an
// EC2 Image Builder pipeline.
//
// Example:
//   import { CfnOutput } from 'aws-cdk-lib';
//   import { Ec2ImageBuilderGetImage } from '@cdklabs/cdk-proserve-lib/constructs';
//
//   const image = new Ec2ImageBuilderGetImage(this, 'SomeImage', {
//     imageBuildVersionArn: 'arn:aws:imagebuilder:us-east-1:123456789012:image/some-image/0.0.1/1'
//   });
//   new CfnOutput(this, 'AmiId', { value: image.ami });
//
// Experimental.
type Ec2ImageBuilderGetImage interface {
	constructs.Construct
	// The AMI ID retrieved from the EC2 Image Builder image.
	// Experimental.
	Ami() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Ec2ImageBuilderGetImage
type jsiiProxy_Ec2ImageBuilderGetImage struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_Ec2ImageBuilderGetImage) Ami() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ami",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ImageBuilderGetImage) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Retrieves an EC2 Image Builder image build version.
// Experimental.
func NewEc2ImageBuilderGetImage(scope constructs.Construct, id *string, props *Ec2ImageBuilderGetImageProps) Ec2ImageBuilderGetImage {
	_init_.Initialize()

	if err := validateNewEc2ImageBuilderGetImageParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2ImageBuilderGetImage{}

	_jsii_.Create(
		"@cdklabs/cdk-proserve-lib.constructs.Ec2ImageBuilderGetImage",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Retrieves an EC2 Image Builder image build version.
// Experimental.
func NewEc2ImageBuilderGetImage_Override(e Ec2ImageBuilderGetImage, scope constructs.Construct, id *string, props *Ec2ImageBuilderGetImageProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-proserve-lib.constructs.Ec2ImageBuilderGetImage",
		[]interface{}{scope, id, props},
		e,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func Ec2ImageBuilderGetImage_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEc2ImageBuilderGetImage_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-proserve-lib.constructs.Ec2ImageBuilderGetImage",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2ImageBuilderGetImage) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

