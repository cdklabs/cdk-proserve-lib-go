package aspects

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/aspects/internal"
)

// Automatically shut down EC2 instances when an alarm is triggered based off of a provided metric.
//
// 🚩 If you are applying this Aspect to multiple EC2 instances, you
// will need to configure the CDK context variable flag
// `@aws-cdk/aws-cloudwatch-actions:changeLambdaPermissionLogicalIdForLambdaAction`
// set to `true`. If this is not configured, applying this Aspect to multiple
// EC2 instances will result in a CDK synth error.
//
// Allows for cost optimization and the reduction of resources not being
// actively used. When the EC2 alarm is triggered for a given EC2 instance, it
// will automatically trigger a Lambda function to shutdown the instance.
//
// Example:
//   import { App, Aspects, Duration, Stack } from 'aws-cdk-lib';
//   import { ComparisonOperator, Stats } from 'aws-cdk-lib/aws-cloudwatch';
//   import { Instance } from 'aws-cdk-lib/aws-ec2';
//   import { Ec2AutomatedShutdown } from './src/aspects/ec2-automated-shutdown';
//
//   const app = new App({
//       context: {
//           '@aws-cdk/aws-cloudwatch-actions:changeLambdaPermissionLogicalIdForLambdaAction':
//               true
//       }
//   });
//   const stack = new Stack(app, 'MyStack');
//
//   // Create your EC2 instance(s)
//   const instance = new Instance(stack, 'MyInstance', {
//       // instance properties
//   });
//
//   // Apply the aspect to automatically shut down the EC2 instance when underutilized
//   Aspects.of(stack).add(new Ec2AutomatedShutdown());
//
//   // Or with custom configuration
//   Aspects.of(stack).add(
//       new Ec2AutomatedShutdown({
//           alarmConfig: {
//               metricName: Ec2AutomatedShutdown.Ec2MetricName.NETWORK_IN,
//               period: Duration.minutes(5),
//               statistic: Stats.AVERAGE,
//               threshold: 100, // 100 bytes
//               evaluationPeriods: 6,
//               datapointsToAlarm: 5,
//               comparisonOperator: ComparisonOperator.LESS_THAN_THRESHOLD
//           }
//       })
//   );
//
// Experimental.
type Ec2AutomatedShutdown interface {
	awscdk.IAspect
	// All aspects can visit an IConstruct.
	// Experimental.
	Visit(node constructs.IConstruct)
}

// The jsii proxy struct for Ec2AutomatedShutdown
type jsiiProxy_Ec2AutomatedShutdown struct {
	internal.Type__awscdkIAspect
}

// Experimental.
func NewEc2AutomatedShutdown(props *Ec2AutomatedShutdownProps) Ec2AutomatedShutdown {
	_init_.Initialize()

	if err := validateNewEc2AutomatedShutdownParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2AutomatedShutdown{}

	_jsii_.Create(
		"@cdklabs/cdk-proserve-lib.aspects.Ec2AutomatedShutdown",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewEc2AutomatedShutdown_Override(e Ec2AutomatedShutdown, props *Ec2AutomatedShutdownProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-proserve-lib.aspects.Ec2AutomatedShutdown",
		[]interface{}{props},
		e,
	)
}

func (e *jsiiProxy_Ec2AutomatedShutdown) Visit(node constructs.IConstruct) {
	if err := e.validateVisitParameters(node); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"visit",
		[]interface{}{node},
	)
}

