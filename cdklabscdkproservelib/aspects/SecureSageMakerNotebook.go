package aspects

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/aspects/internal"
)

// Aspect that enforces security controls on SageMaker Notebook Instances by requiring VPC placement, disabling direct internet access, and preventing root access to the notebook environment.
//
// This Aspect enforces these settings through a combination of setting
// the CloudFormation properties on the Notebook resource and attaching a
// DENY policy to the role that is used by the notebook. The policy will enforce
// that the following API actions contain the correct properties to ensure
// network isolation and that the VPC subnets are set:
//
// - 'sagemaker:CreateTrainingJob',
// - 'sagemaker:CreateHyperParameterTuningJob',
// - 'sagemaker:CreateModel',
// - 'sagemaker:CreateProcessingJob'.
// Experimental.
type SecureSageMakerNotebook interface {
	awscdk.IAspect
	// All aspects can visit an IConstruct.
	// Experimental.
	Visit(node constructs.IConstruct)
}

// The jsii proxy struct for SecureSageMakerNotebook
type jsiiProxy_SecureSageMakerNotebook struct {
	internal.Type__awscdkIAspect
}

// Experimental.
func NewSecureSageMakerNotebook(props *SecureSageMakerNotebookProps) SecureSageMakerNotebook {
	_init_.Initialize()

	if err := validateNewSecureSageMakerNotebookParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SecureSageMakerNotebook{}

	_jsii_.Create(
		"@cdklabs/cdk-proserve-lib.aspects.SecureSageMakerNotebook",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewSecureSageMakerNotebook_Override(s SecureSageMakerNotebook, props *SecureSageMakerNotebookProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-proserve-lib.aspects.SecureSageMakerNotebook",
		[]interface{}{props},
		s,
	)
}

func (s *jsiiProxy_SecureSageMakerNotebook) Visit(node constructs.IConstruct) {
	if err := s.validateVisitParameters(node); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"visit",
		[]interface{}{node},
	)
}

