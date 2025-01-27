package constructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/constructs/internal"
)

// Starts an EC2 Image Builder Pipeline and optionally waits for the build to complete.
//
// This construct is useful if you want to create an image as part of your IaC
// deployment. By waiting for completion of this construct, you can use the
// image in the same deployment by retrieving the AMI and passing it to an EC2
// build step.
//
// Example:
//   import { Duration } from 'aws-cdk-lib';
//   import { Topic } from 'aws-cdk-lib/aws-sns';
//   import { Ec2ImageBuilderStart } from '@cdklabs/cdk-proserve-lib/constructs';
//
//   const topic = Topic.fromTopicArn(
//     this,
//     'MyTopic',
//     'arn:aws:sns:us-east-1:123456789012:my-notification-topic'
//   );
//   new Ec2ImageBuilderStart(this, 'ImageBuilderStart', {
//     pipelineArn:
//       'arn:aws:imagebuilder:us-east-1:123456789012:image-pipeline/my-image-pipeline',
//     waitForCompletion: {
//       topic: topic,
//       timeout: Duration.hours(7)  // wait up to 7 hours for completion
//     }
//   });
//
// Experimental.
type Ec2ImageBuilderStart interface {
	Construct
	// The ARN of the image build version created by the pipeline execution.
	// Experimental.
	ImageBuildVersionArn() *string
	// The tree node.
	// Experimental.
	Node() Node
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Ec2ImageBuilderStart
type jsiiProxy_Ec2ImageBuilderStart struct {
	internal.Type__Construct
}

func (j *jsiiProxy_Ec2ImageBuilderStart) ImageBuildVersionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageBuildVersionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ImageBuilderStart) Node() Node {
	var returns Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Starts an EC2 Image Builder Pipeline and optionally waits for the build to complete.
// Experimental.
func NewEc2ImageBuilderStart(scope Construct, id *string, props *Ec2ImageBuilderStartProps) Ec2ImageBuilderStart {
	_init_.Initialize()

	if err := validateNewEc2ImageBuilderStartParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2ImageBuilderStart{}

	_jsii_.Create(
		"@cdklabs/cdk-proserve-lib.constructs.Ec2ImageBuilderStart",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Starts an EC2 Image Builder Pipeline and optionally waits for the build to complete.
// Experimental.
func NewEc2ImageBuilderStart_Override(e Ec2ImageBuilderStart, scope Construct, id *string, props *Ec2ImageBuilderStartProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-proserve-lib.constructs.Ec2ImageBuilderStart",
		[]interface{}{scope, id, props},
		e,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func Ec2ImageBuilderStart_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEc2ImageBuilderStart_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-proserve-lib.constructs.Ec2ImageBuilderStart",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2ImageBuilderStart) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

