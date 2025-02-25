package patterns

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/patterns/internal"
)

// A pattern to build an EC2 Image Pipeline specifically for Linux.
//
// This pattern contains opinionated code and features to help create a linux
// pipeline. This pattern further simplifies setting up an image pipeline by
// letting you choose specific operating systems and features.
//
// The example below shows how you can configure an image that contains the AWS
// CLI and retains the SSM agent on the image. The image will have a 100GB root
// volume.
//
// Example:
//   import { Ec2LinuxImagePipeline } from '@cdklabs/cdk-proserve-lib/patterns';
//
//   new Ec2LinuxImagePipeline(this, 'ImagePipeline', {
//     version: '0.1.0',
//     operatingSystem:
//       Ec2LinuxImagePipeline.OperatingSystem.RED_HAT_ENTERPRISE_LINUX_8_9,
//     rootVolumeSize: 100,
//       buildConfiguration: {
//         start: true,
//         waitForCompletion: true
//       },
//     features: [
//       Ec2LinuxImagePipeline.Feature.AWS_CLI,
//       Ec2LinuxImagePipeline.Feature.RETAIN_SSM_AGENT
//     ]
//   );
//
// Experimental.
type Ec2LinuxImagePipeline interface {
	constructs.Construct
	// The Amazon Resource Name (ARN) of the Image Pipeline.
	//
	// Used to uniquely identify this image pipeline.
	// Experimental.
	ImagePipelineArn() *string
	// Experimental.
	SetImagePipelineArn(val *string)
	// The SNS Topic associated with this Image Pipeline.
	//
	// Publishes notifications about pipeline execution events.
	// Experimental.
	ImagePipelineTopic() awssns.ITopic
	// Experimental.
	SetImagePipelineTopic(val awssns.ITopic)
	// The latest AMI built by the pipeline.
	//
	// NOTE: You must have enabled the
	// Build Configuration option to wait for image build completion for this
	// property to be available.
	// Experimental.
	LatestAmi() *string
	// Experimental.
	SetLatestAmi(val *string)
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Ec2LinuxImagePipeline
type jsiiProxy_Ec2LinuxImagePipeline struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_Ec2LinuxImagePipeline) ImagePipelineArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagePipelineArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LinuxImagePipeline) ImagePipelineTopic() awssns.ITopic {
	var returns awssns.ITopic
	_jsii_.Get(
		j,
		"imagePipelineTopic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LinuxImagePipeline) LatestAmi() *string {
	var returns *string
	_jsii_.Get(
		j,
		"latestAmi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LinuxImagePipeline) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// A pattern to build an EC2 Image Pipeline specifically for Linux.
// Experimental.
func NewEc2LinuxImagePipeline(scope constructs.Construct, id *string, props *Ec2LinuxImagePipelineProps) Ec2LinuxImagePipeline {
	_init_.Initialize()

	if err := validateNewEc2LinuxImagePipelineParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2LinuxImagePipeline{}

	_jsii_.Create(
		"@cdklabs/cdk-proserve-lib.patterns.Ec2LinuxImagePipeline",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// A pattern to build an EC2 Image Pipeline specifically for Linux.
// Experimental.
func NewEc2LinuxImagePipeline_Override(e Ec2LinuxImagePipeline, scope constructs.Construct, id *string, props *Ec2LinuxImagePipelineProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-proserve-lib.patterns.Ec2LinuxImagePipeline",
		[]interface{}{scope, id, props},
		e,
	)
}

func (j *jsiiProxy_Ec2LinuxImagePipeline)SetImagePipelineArn(val *string) {
	if err := j.validateSetImagePipelineArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imagePipelineArn",
		val,
	)
}

func (j *jsiiProxy_Ec2LinuxImagePipeline)SetImagePipelineTopic(val awssns.ITopic) {
	if err := j.validateSetImagePipelineTopicParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imagePipelineTopic",
		val,
	)
}

func (j *jsiiProxy_Ec2LinuxImagePipeline)SetLatestAmi(val *string) {
	_jsii_.Set(
		j,
		"latestAmi",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func Ec2LinuxImagePipeline_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEc2LinuxImagePipeline_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-proserve-lib.patterns.Ec2LinuxImagePipeline",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2LinuxImagePipeline) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

