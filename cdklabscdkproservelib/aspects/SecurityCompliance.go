package aspects

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/aspects/internal"
)

// Applies best practice security settings to be in compliance with security tools such as CDK Nag.
//
// This aspect automatically implements AWS security best practices and compliance
// requirements for various AWS services used in your CDK applications.
// It can be configured with custom settings and supports suppressing specific
// CDK Nag warnings with proper justification.
//
// Example:
//   import { App, Stack, Aspects } from 'aws-cdk-lib';
//   import { Function, Runtime, Code } from 'aws-cdk-lib/aws-lambda';
//   import { Bucket } from 'aws-cdk-lib/aws-s3';
//   import { SecurityCompliance } from '../../../src/aspects/security-compliance';
//
//   const app = new App();
//   const stack = new Stack(app, 'MySecureStack');
//
//   // Create resources
//   const myBucket = new Bucket(stack, 'MyBucket');
//   const myFunction = new Function(stack, 'MyFunction', {
//       runtime: Runtime.NODEJS_18_X,
//       handler: 'index.handler',
//       code: Code.fromInline(
//           'exports.handler = async () => { return { statusCode: 200 }; }'
//       )
//   });
//
//   // Apply the SecurityCompliance aspect with custom settings
//   const securityAspect = new SecurityCompliance({
//       settings: {
//           s3: {
//               serverAccessLogs: {
//                   destinationBucketName: 'my-access-logs-bucket'
//               },
//               versioning: {
//                   disabled: false
//               }
//           },
//           lambda: {
//               reservedConcurrentExecutions: {
//                   concurrentExecutionCount: 5
//               }
//           }
//       },
//       suppressions: {
//           lambdaNotInVpc:
//               'This is a development environment where VPC is not required',
//           iamNoInlinePolicies: 'Inline policies are acceptable for this use case'
//       }
//   });
//
//   // Apply the aspect to the stack
//   Aspects.of(app).add(securityAspect);
//
// Experimental.
type SecurityCompliance interface {
	awscdk.IAspect
	// Apply the aspect to the node.
	// Experimental.
	Visit(node constructs.IConstruct)
}

// The jsii proxy struct for SecurityCompliance
type jsiiProxy_SecurityCompliance struct {
	internal.Type__awscdkIAspect
}

// Experimental.
func NewSecurityCompliance(props *SecurityComplianceProps) SecurityCompliance {
	_init_.Initialize()

	if err := validateNewSecurityComplianceParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SecurityCompliance{}

	_jsii_.Create(
		"@cdklabs/cdk-proserve-lib.aspects.SecurityCompliance",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewSecurityCompliance_Override(s SecurityCompliance, props *SecurityComplianceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-proserve-lib.aspects.SecurityCompliance",
		[]interface{}{props},
		s,
	)
}

func (s *jsiiProxy_SecurityCompliance) Visit(node constructs.IConstruct) {
	if err := s.validateVisitParameters(node); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"visit",
		[]interface{}{node},
	)
}

