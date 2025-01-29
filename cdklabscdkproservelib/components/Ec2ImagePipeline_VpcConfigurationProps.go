package components

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// VPC Configuration.
// Experimental.
type Ec2ImagePipeline_VpcConfigurationProps struct {
	// Experimental.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// Experimental.
	SecurityGroup awsec2.ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// Experimental.
	Subnet awsec2.ISubnet `field:"optional" json:"subnet" yaml:"subnet"`
}

