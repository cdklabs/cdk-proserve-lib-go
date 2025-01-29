package components

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsnetworkfirewall"
	"github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/interfaces"
)

// Properties for the NetworkFirewallEndpoints construct.
// Experimental.
type NetworkFirewallEndpointsProps struct {
	// The AWS Network Firewall to get the Endpoints for.
	// Experimental.
	Firewall awsnetworkfirewall.CfnFirewall `field:"required" json:"firewall" yaml:"firewall"`
	// Optional Lambda configuration settings.
	// Experimental.
	LambdaConfiguration *interfaces.AwsCustomResourceLambdaConfiguration `field:"optional" json:"lambdaConfiguration" yaml:"lambdaConfiguration"`
}

