package aspects

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/types"
)

// Experimental.
type Ec2AutomatedShutdownProps struct {
	// Optional custom metric configuration.
	//
	// If not provided, defaults to CPU utilization with a 5% threshold.
	// Experimental.
	AlarmConfig *AlarmConfig `field:"optional" json:"alarmConfig" yaml:"alarmConfig"`
	// Optional KMS Encryption Key to use for encrypting resources.
	// Experimental.
	Encryption awskms.IKey `field:"optional" json:"encryption" yaml:"encryption"`
	// Optional Lambda configuration settings.
	// Experimental.
	LambdaConfiguration *types.LambdaConfiguration `field:"optional" json:"lambdaConfiguration" yaml:"lambdaConfiguration"`
}

