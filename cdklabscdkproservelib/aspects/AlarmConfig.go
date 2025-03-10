package aspects

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Optional custom metric configuration for CloudWatch Alarms.
//
// If not provided, defaults to CPU utilization with a 5% threshold.
// Experimental.
type AlarmConfig struct {
	// The comparison operator to use for the alarm.
	// Default: = ComparisonOperator.LESS_THAN_THRESHOLD
	//
	// Experimental.
	ComparisonOperator awscloudwatch.ComparisonOperator `field:"required" json:"comparisonOperator" yaml:"comparisonOperator"`
	// The number of datapoints that must go past/below the threshold to trigger the alarm.
	// Default: = 2.
	//
	// Experimental.
	DatapointsToAlarm *float64 `field:"required" json:"datapointsToAlarm" yaml:"datapointsToAlarm"`
	// The number of periods over which data is compared to the specified threshold.
	// Default: = 3.
	//
	// Experimental.
	EvaluationPeriods *float64 `field:"required" json:"evaluationPeriods" yaml:"evaluationPeriods"`
	// The name of the CloudWatch metric to monitor.
	// Default: = CPUUtilization.
	//
	// Experimental.
	MetricName Ec2AutomatedShutdown_Ec2MetricName `field:"required" json:"metricName" yaml:"metricName"`
	// The period over which the metric is measured.
	// Default: = 1 minute.
	//
	// Experimental.
	Period awscdk.Duration `field:"required" json:"period" yaml:"period"`
	// The CloudWatch metric statistic to use.
	//
	// Use the `aws_cloudwatch.Stats` helper class to construct valid input
	// strings.
	// Default: = 'Average'.
	//
	// Experimental.
	Statistic *string `field:"required" json:"statistic" yaml:"statistic"`
	// The threshold value for the alarm.
	// Default: = 5%.
	//
	// Experimental.
	Threshold *float64 `field:"required" json:"threshold" yaml:"threshold"`
}

