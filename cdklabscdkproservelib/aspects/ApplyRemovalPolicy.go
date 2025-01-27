package aspects

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/aspects/internal"
)

// Sets a user specified Removal Policy to all resources that the aspect applies to.
//
// This Aspect is useful if you want to enforce a specified removal policy on
// resources. For example, you could ensure that your removal policy is always
// set to RETAIN or DESTROY.
//
// Example:
//   import { App, Aspects, RemovalPolicy } from 'aws-cdk-lib';
//   import { ApplyRemovalPolicy } from '@cdklabs/cdk-proserve-lib/aspects';
//
//   const app = new App();
//
//   Aspects.of(app).add(
//     new ApplyRemovalPolicy({ removalPolicy: RemovalPolicy.DESTROY })
//   );
//
// Experimental.
type ApplyRemovalPolicy interface {
	awscdk.IAspect
	// Visits a construct and applies the removal policy.
	// Experimental.
	Visit(node constructs.IConstruct)
}

// The jsii proxy struct for ApplyRemovalPolicy
type jsiiProxy_ApplyRemovalPolicy struct {
	internal.Type__awscdkIAspect
}

// Creates a new instance of SetRemovalPolicy.
// Experimental.
func NewApplyRemovalPolicy(props *ApplyRemovalPolicyProps) ApplyRemovalPolicy {
	_init_.Initialize()

	if err := validateNewApplyRemovalPolicyParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApplyRemovalPolicy{}

	_jsii_.Create(
		"@cdklabs/cdk-proserve-lib.aspects.ApplyRemovalPolicy",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Creates a new instance of SetRemovalPolicy.
// Experimental.
func NewApplyRemovalPolicy_Override(a ApplyRemovalPolicy, props *ApplyRemovalPolicyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-proserve-lib.aspects.ApplyRemovalPolicy",
		[]interface{}{props},
		a,
	)
}

func (a *jsiiProxy_ApplyRemovalPolicy) Visit(node constructs.IConstruct) {
	if err := a.validateVisitParameters(node); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"visit",
		[]interface{}{node},
	)
}

