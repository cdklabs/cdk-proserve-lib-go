package aspects

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Properties for configuring log retention settings.
// Experimental.
type SetLogRetentionProps struct {
	// The retention period for the logs.
	// Experimental.
	Period awslogs.RetentionDays `field:"required" json:"period" yaml:"period"`
}

