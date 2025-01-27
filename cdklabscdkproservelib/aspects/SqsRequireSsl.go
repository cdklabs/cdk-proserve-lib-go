package aspects

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/aspects/internal"
)

// Enforces SSL/TLS requirements on Simple Queue Service (SQS) for all resources that the aspect applies to.
//
// This is accomplished by adding a resource policy to any SQS queue that denies
// all actions when the request is not made over a secure transport.
//
// Example:
//   import { App, Aspects } from 'aws-cdk-lib';
//   import { SqsRequireSsl } from '@cdklabs/cdk-proserve-lib/aspects';
//
//   const app = new App();
//
//   Aspects.of(app).add(new SqsRequireSsl());
//
// Experimental.
type SqsRequireSsl interface {
	awscdk.IAspect
	// Visits a construct and adds SSL/TLS requirement policy if it's an SQS queue.
	// Experimental.
	Visit(node constructs.IConstruct)
}

// The jsii proxy struct for SqsRequireSsl
type jsiiProxy_SqsRequireSsl struct {
	internal.Type__awscdkIAspect
}

// Experimental.
func NewSqsRequireSsl() SqsRequireSsl {
	_init_.Initialize()

	j := jsiiProxy_SqsRequireSsl{}

	_jsii_.Create(
		"@cdklabs/cdk-proserve-lib.aspects.SqsRequireSsl",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewSqsRequireSsl_Override(s SqsRequireSsl) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-proserve-lib.aspects.SqsRequireSsl",
		nil, // no parameters
		s,
	)
}

func (s *jsiiProxy_SqsRequireSsl) Visit(node constructs.IConstruct) {
	if err := s.validateVisitParameters(node); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"visit",
		[]interface{}{node},
	)
}

