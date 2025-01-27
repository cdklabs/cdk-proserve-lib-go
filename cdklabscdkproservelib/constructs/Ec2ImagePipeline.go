package constructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/constructs/internal"
)

// An EC2 Image Pipeline that can be used to build a Amazon Machine Image (AMI) automatically.
//
// This construct simplifies the process of creating an EC2 Image Pipeline and
// provides all of the available components that can be used that are maintained
// by AWS.
//
// Example:
//   import { CfnOutput } from 'aws-cdk-lib';
//   import { Ec2ImagePipeline } from '@cdklabs/cdk-proserve-lib/constructs';
//
//   const pipeline = new Ec2ImagePipeline(this, 'ImagePipeline', {
//     version: '0.1.0',
//     buildConfiguration: {
//       start: true,
//       waitForCompletion: true
//     },
//     components: [
//       Ec2ImagePipeline.Component.AWS_CLI_VERSION_2_LINUX,
//       Ec2ImagePipeline.Component.DOCKER_CE_LINUX
//     ]
//   });
//   new CfnOutput(this, 'ImagePipelineAmi', { value: pipeline.latestAmi! });
//
// Experimental.
type Ec2ImagePipeline interface {
	Construct
	// The Image Pipeline ARN that gets created.
	// Experimental.
	ImagePipelineArn() *string
	// The Image Pipeline Topic that gets created.
	// Experimental.
	ImagePipelineTopic() awssns.ITopic
	// The latest AMI built by the pipeline.
	//
	// NOTE: You must have enabled the
	// Build Configuration option to wait for image build completion for this
	// property to be available.
	// Experimental.
	LatestAmi() *string
	// The tree node.
	// Experimental.
	Node() Node
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Ec2ImagePipeline
type jsiiProxy_Ec2ImagePipeline struct {
	internal.Type__Construct
}

func (j *jsiiProxy_Ec2ImagePipeline) ImagePipelineArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagePipelineArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ImagePipeline) ImagePipelineTopic() awssns.ITopic {
	var returns awssns.ITopic
	_jsii_.Get(
		j,
		"imagePipelineTopic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ImagePipeline) LatestAmi() *string {
	var returns *string
	_jsii_.Get(
		j,
		"latestAmi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ImagePipeline) Node() Node {
	var returns Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// An EC2 Image Pipeline that can be used to build a Amazon Machine Image (AMI) automatically.
// Experimental.
func NewEc2ImagePipeline(scope Construct, id *string, props *Ec2ImagePipelineProps) Ec2ImagePipeline {
	_init_.Initialize()

	if err := validateNewEc2ImagePipelineParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2ImagePipeline{}

	_jsii_.Create(
		"@cdklabs/cdk-proserve-lib.constructs.Ec2ImagePipeline",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// An EC2 Image Pipeline that can be used to build a Amazon Machine Image (AMI) automatically.
// Experimental.
func NewEc2ImagePipeline_Override(e Ec2ImagePipeline, scope Construct, id *string, props *Ec2ImagePipelineProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-proserve-lib.constructs.Ec2ImagePipeline",
		[]interface{}{scope, id, props},
		e,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func Ec2ImagePipeline_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEc2ImagePipeline_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-proserve-lib.constructs.Ec2ImagePipeline",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2ImagePipeline) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

