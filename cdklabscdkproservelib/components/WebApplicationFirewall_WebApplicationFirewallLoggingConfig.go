package components

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Experimental.
type WebApplicationFirewall_WebApplicationFirewallLoggingConfig struct {
	// Log Group name affix to be appended to aws-waf-logs-<affix>.
	// Experimental.
	LogGroupNameAffix *string `field:"required" json:"logGroupNameAffix" yaml:"logGroupNameAffix"`
	// KMS key to use for encryption of the log group.
	// Experimental.
	EncryptionKey awskms.Key `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// Removal policy for the log group.
	// Default: DESTROY.
	//
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// Retention period for the log group.
	// Default: ONE_MONTH.
	//
	// Experimental.
	Retention awslogs.RetentionDays `field:"optional" json:"retention" yaml:"retention"`
}

