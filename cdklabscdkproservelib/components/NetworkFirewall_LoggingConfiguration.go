package components

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Experimental.
type NetworkFirewall_LoggingConfiguration struct {
	// The type of logs to write for the Network Firewall.
	//
	// This can be `TLS`, `FLOW`, or `ALERT`.
	// Experimental.
	LogTypes *[]NetworkFirewall_LogType `field:"required" json:"logTypes" yaml:"logTypes"`
	// Optional KMS key for encrypting Network Firewall logs.
	// Experimental.
	Encryption awskms.IKey `field:"optional" json:"encryption" yaml:"encryption"`
	// Log group to use for Network Firewall Logging.
	//
	// If not specified, a log group is created for you. The encryption key provided will be used to encrypt it if one was provided to the construct.
	// Experimental.
	LogGroup awslogs.ILogGroup `field:"optional" json:"logGroup" yaml:"logGroup"`
	// If you do not specify a log group, the amount of time to keep logs in the automatically created Log Group.
	//
	// Default: one week.
	// Experimental.
	LogRetention awslogs.RetentionDays `field:"optional" json:"logRetention" yaml:"logRetention"`
}

