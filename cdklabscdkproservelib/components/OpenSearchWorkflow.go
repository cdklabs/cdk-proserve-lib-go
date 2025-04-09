package components

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/components/internal"
)

// Create OpenSearch Workflows using the flow framework to automate the provisioning of complex tasks using JSON or YAML.
//
// This construct creates a custom resource that deploys a Flow Framework
// template to an OpenSearch domain. It handles the deployment and lifecycle
// management of the workflow through a Lambda-backed custom resources. You can
// read more about the flow framework on AWS at the reference link below.
//
// Example:
//   import { OpenSearchWorkflow } from '@cdklabs/cdk-proserve-lib/constructs';
//   import { Domain } from 'aws-cdk-lib/aws-opensearchservice';
//   import { Role } from 'aws-cdk-lib/aws-iam';
//
//   const aosDomain = Domain.fromDomainEndpoint(this, 'Domain', 'aos-endpoint');
//   const aosRole = Role.fromRoleName(this, 'Role', 'AosRole');
//
//   // Create OpenSearch Workflow using a YAML workflow template
//   const nlpIngestPipeline = new OpenSearchWorkflow(
//       this,
//       'NlpIngestPipeline',
//       {
//           domain: aosDomain,
//           domainAuthentication: aosRole,
//           flowFrameworkTemplatePath: join(
//               __dirname,
//               'nlp-ingest-pipeline.yaml'
//           )
//       }
//   );
//
//   // Retrieve the deployed model from the OpenSearch Workflow
//   this.embeddingModelId = nlpIngestPipeline.getResourceId(
//       'deploy_sentence_model'
//   );
//
// Experimental.
type OpenSearchWorkflow interface {
	constructs.Construct
	// The Lambda function that will be called to determine if the execution is complete for the custom resource.
	// Experimental.
	IsCompleteHandler() awslambda.IFunction
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// The Lambda function that will handle On Event requests for the custom resource.
	// Experimental.
	OnEventHandler() awslambda.IFunction
	// The unique identifier of the deployed OpenSearch workflow.
	//
	// This ID can be used to reference and manage the workflow after deployment.
	// Experimental.
	WorkflowId() *string
	// Retrieves a created Resource ID from the Workflow by the provided workflowStepId.
	//
	// The workflowStepId is the `id` value of the node in your
	// list of workflow nodes from your workflow template.
	//
	// Returns: string value of the resource id that was created.
	// Experimental.
	GetResourceId(workflowStepId *string) *string
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OpenSearchWorkflow
type jsiiProxy_OpenSearchWorkflow struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_OpenSearchWorkflow) IsCompleteHandler() awslambda.IFunction {
	var returns awslambda.IFunction
	_jsii_.Get(
		j,
		"isCompleteHandler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenSearchWorkflow) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenSearchWorkflow) OnEventHandler() awslambda.IFunction {
	var returns awslambda.IFunction
	_jsii_.Get(
		j,
		"onEventHandler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenSearchWorkflow) WorkflowId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workflowId",
		&returns,
	)
	return returns
}


// Constructor.
// Experimental.
func NewOpenSearchWorkflow(scope constructs.Construct, id *string, props *OpenSearchWorkflowProps) OpenSearchWorkflow {
	_init_.Initialize()

	if err := validateNewOpenSearchWorkflowParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_OpenSearchWorkflow{}

	_jsii_.Create(
		"@cdklabs/cdk-proserve-lib.constructs.OpenSearchWorkflow",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructor.
// Experimental.
func NewOpenSearchWorkflow_Override(o OpenSearchWorkflow, scope constructs.Construct, id *string, props *OpenSearchWorkflowProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-proserve-lib.constructs.OpenSearchWorkflow",
		[]interface{}{scope, id, props},
		o,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func OpenSearchWorkflow_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOpenSearchWorkflow_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-proserve-lib.constructs.OpenSearchWorkflow",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenSearchWorkflow) GetResourceId(workflowStepId *string) *string {
	if err := o.validateGetResourceIdParameters(workflowStepId); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getResourceId",
		[]interface{}{workflowStepId},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenSearchWorkflow) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

