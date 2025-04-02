package aspects

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
)

// Experimental.
type SecurityCompliance_StageMethodLogging struct {
	// Sets the setting to disabled.
	//
	// This does not actually make an impact on
	// the setting itself, it just stops this aspect from making changes to
	// the specific setting.
	// Experimental.
	Disabled *bool `field:"optional" json:"disabled" yaml:"disabled"`
	// The logging level to use for the stage method logging. This applies to all resources and methods in all stages.
	//
	// Defaults to MethodLoggingLevel.ERROR if not specified.
	// Experimental.
	LoggingLevel awsapigateway.MethodLoggingLevel `field:"optional" json:"loggingLevel" yaml:"loggingLevel"`
}

