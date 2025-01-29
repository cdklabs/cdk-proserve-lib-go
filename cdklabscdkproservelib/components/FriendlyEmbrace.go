package components

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambdanodejs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/components/internal"
)

// The Friendly Embrace construct can be used to remove CloudFormation stack dependencies that are based on stack exports and imports.
//
// 🚨 WARNING: This construct is experimental and will directly modify
// CloudFormation stacks in your CDK application via a Lambda-backed Custom
// Resource. It is not recommended to use this construct in a production
// environment.
//
// A custom resource that is designed to remove the "Deadly Embrace" problem that
// occurs when you attempt to update a CloudFormation stack that is exporting
// a resource used by another stack. This custom resource will run before all of
// your stacks deploy/update and remove the dependencies by hardcoding each
// export for all stacks that use it. For this reason, this construct should run
// inside of its own stack and should be the last stack that is instantiated for
// your application. That way the construct will be able to retrieve all of the
// stacks from your CDK resource tree that it needs to update.
//
// > NOTE: You may need to add more permissions to the handler if the custom
// resource cannot update your stacks. You can call upon the `handler` property
// of the class to add more permissions to it.
//
// Example:
//   import { App, Stack } from 'aws-cdk-lib';
//   import { FriendlyEmbrace } from '@cdklabs/cdk-proserve-lib/constructs';
//
//   const app = new App();
//   // ... other stack definitions
//   const embrace = new Stack(app, 'FriendlyEmbrace'); // last stack
//   new FriendlyEmbrace(embrace, 'FriendlyEmbrace'); // only construct in stack
//
// Experimental.
type FriendlyEmbrace interface {
	constructs.Construct
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Handler for the custom resource.
	// Experimental.
	OnEventHandler() awslambdanodejs.NodejsFunction
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FriendlyEmbrace
type jsiiProxy_FriendlyEmbrace struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_FriendlyEmbrace) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FriendlyEmbrace) OnEventHandler() awslambdanodejs.NodejsFunction {
	var returns awslambdanodejs.NodejsFunction
	_jsii_.Get(
		j,
		"onEventHandler",
		&returns,
	)
	return returns
}


// The Friendly Embrace construct can be used to remove CloudFormation stack dependencies that are based on stack exports and imports.
//
// 🚨 WARNING: This construct is experimental and will directly modify
// CloudFormation stacks in your CDK application via a Lambda-backed Custom
// Resource. It is not recommended to use this construct in a production
// environment.
// Experimental.
func NewFriendlyEmbrace(scope constructs.Construct, id *string, props *FriendlyEmbraceProps) FriendlyEmbrace {
	_init_.Initialize()

	if err := validateNewFriendlyEmbraceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_FriendlyEmbrace{}

	_jsii_.Create(
		"@cdklabs/cdk-proserve-lib.constructs.FriendlyEmbrace",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// The Friendly Embrace construct can be used to remove CloudFormation stack dependencies that are based on stack exports and imports.
//
// 🚨 WARNING: This construct is experimental and will directly modify
// CloudFormation stacks in your CDK application via a Lambda-backed Custom
// Resource. It is not recommended to use this construct in a production
// environment.
// Experimental.
func NewFriendlyEmbrace_Override(f FriendlyEmbrace, scope constructs.Construct, id *string, props *FriendlyEmbraceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-proserve-lib.constructs.FriendlyEmbrace",
		[]interface{}{scope, id, props},
		f,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func FriendlyEmbrace_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFriendlyEmbrace_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-proserve-lib.constructs.FriendlyEmbrace",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FriendlyEmbrace) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

