package constructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
)

// Experimental.
type Ec2ImageBuilderStart_WaitForCompletionProps struct {
	// An SNS Topic that will signal when the pipeline is complete.
	//
	// This is
	// typically configured on your EC2 Image Builder pipeline to trigger an
	// SNS notification when the pipeline completes.
	// Experimental.
	Topic awssns.ITopic `field:"required" json:"topic" yaml:"topic"`
	// The maximum amount of time to wait for the image build pipeline to complete.
	//
	// This is set to a maximum of 12 hours by default.
	// Default: 12 hours.
	//
	// Experimental.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
}

