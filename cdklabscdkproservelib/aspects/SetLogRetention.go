package aspects

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/aspects/internal"
)

// Aspect that sets the log retention period for CloudWatch log groups to a user-supplied retention period.
//
// Example:
//   import { App, Aspects } from 'aws-cdk-lib';
//   import { RetentionDays } from 'aws-cdk-lib/aws-logs';
//   import { SetLogRetention } from '@cdklabs/cdk-proserve-lib/aspects';
//
//   const app = new App();
//
//   Aspects.of(app).add(
//     new SetLogRetention({ period: RetentionDays.EIGHTEEN_MONTHS })
//   );
//
// Experimental.
type SetLogRetention interface {
	awscdk.IAspect
	// Visits a construct and sets log retention if applicable.
	// Experimental.
	Visit(node constructs.IConstruct)
}

// The jsii proxy struct for SetLogRetention
type jsiiProxy_SetLogRetention struct {
	internal.Type__awscdkIAspect
}

// Creates a new instance of SetLogRetention.
// Experimental.
func NewSetLogRetention(props *SetLogRetentionProps) SetLogRetention {
	_init_.Initialize()

	if err := validateNewSetLogRetentionParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SetLogRetention{}

	_jsii_.Create(
		"@cdklabs/cdk-proserve-lib.aspects.SetLogRetention",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Creates a new instance of SetLogRetention.
// Experimental.
func NewSetLogRetention_Override(s SetLogRetention, props *SetLogRetentionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-proserve-lib.aspects.SetLogRetention",
		[]interface{}{props},
		s,
	)
}

func (s *jsiiProxy_SetLogRetention) Visit(node constructs.IConstruct) {
	if err := s.validateVisitParameters(node); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"visit",
		[]interface{}{node},
	)
}

