package aspects

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/aspects/internal"
)

// Ensures that Lambda log groups are created for all Lambda functions that the aspect applies to.
//
// Example:
//   import { App, Aspects } from 'aws-cdk-lib';
//   import { CreateLambdaLogGroup } from '@cdklabs/cdk-proserve-lib/aspects';
//
//   const app = new App();
//
//   Aspects.of(app).add(new CreateLambdaLogGroup());
//
// Experimental.
type CreateLambdaLogGroup interface {
	awscdk.IAspect
	// Visits a construct and creates a log group if the construct is a Lambda function.
	// Experimental.
	Visit(node constructs.IConstruct)
}

// The jsii proxy struct for CreateLambdaLogGroup
type jsiiProxy_CreateLambdaLogGroup struct {
	internal.Type__awscdkIAspect
}

// Experimental.
func NewCreateLambdaLogGroup() CreateLambdaLogGroup {
	_init_.Initialize()

	j := jsiiProxy_CreateLambdaLogGroup{}

	_jsii_.Create(
		"@cdklabs/cdk-proserve-lib.aspects.CreateLambdaLogGroup",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCreateLambdaLogGroup_Override(c CreateLambdaLogGroup) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-proserve-lib.aspects.CreateLambdaLogGroup",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CreateLambdaLogGroup) Visit(node constructs.IConstruct) {
	if err := c.validateVisitParameters(node); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"visit",
		[]interface{}{node},
	)
}

