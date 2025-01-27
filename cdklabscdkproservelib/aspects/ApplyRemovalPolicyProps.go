package aspects

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for configuring the removal policy settings.
// Experimental.
type ApplyRemovalPolicyProps struct {
	// The removal policy to apply to the resource.
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `field:"required" json:"removalPolicy" yaml:"removalPolicy"`
}

