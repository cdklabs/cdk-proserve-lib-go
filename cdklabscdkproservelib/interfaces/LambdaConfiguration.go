package interfaces

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
)

// Experimental.
type LambdaConfiguration struct {
	// Optional SQS queue to use as a dead letter queue.
	// Experimental.
	DeadLetterQueue awssqs.IQueue `field:"optional" json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// Optional retention period for the Lambda functions log group.
	// Default: RetentionDays.ONE_MONTH
	//
	// Experimental.
	LogGroupRetention awslogs.RetentionDays `field:"optional" json:"logGroupRetention" yaml:"logGroupRetention"`
	// The number of concurrent executions for the provider Lambda function.
	//
	// Default: 5.
	// Experimental.
	ReservedConcurrentExecutions *float64 `field:"optional" json:"reservedConcurrentExecutions" yaml:"reservedConcurrentExecutions"`
	// Security groups to attach to the provider Lambda functions.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Optional subnet selection for the Lambda functions.
	// Experimental.
	Subnets *awsec2.SubnetSelection `field:"optional" json:"subnets" yaml:"subnets"`
	// VPC where the Lambda functions will be deployed.
	// Experimental.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

