package patterns

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Policies to lifecycle various components of the pattern during stack actions.
// Experimental.
type KeycloakService_RemovalPolicies struct {
	// How to deal with data-related elements.
	// Default: RemovalPolicy.RETAIN
	//
	// Experimental.
	Data awscdk.RemovalPolicy `field:"optional" json:"data" yaml:"data"`
	// How to deal with log-related elements.
	// Default: RemovalPolicy.RETAIN
	//
	// Experimental.
	Logs awscdk.RemovalPolicy `field:"optional" json:"logs" yaml:"logs"`
}

