package constructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/interfaces"
)

// Experimental.
type NetworkFirewall_NetworkFirewallVpcRouteProps struct {
	// Subnets that will sit behind the network firewall and should have routes to the Network Firewall.
	//
	// By supplying this parameter, routes will
	// be created for these subnets to the Network Firewall. Specify the
	// optional `destinationCidr` parameter if you want to restrict the
	// routes to a specific CIDR block. By default, routes will be created
	// for all outbound traffic (0.0.0.0/0) to the firewall.
	// Experimental.
	ProtectedSubnets *[]awsec2.ISubnet `field:"required" json:"protectedSubnets" yaml:"protectedSubnets"`
	// The destination CIDR block for the firewall (protectedSubnets) route.
	//
	// If not specified, defaults to '0.0.0.0/0' (all IPv4 traffic).
	// Experimental.
	DestinationCidr *string `field:"optional" json:"destinationCidr" yaml:"destinationCidr"`
	// Configuration for the Lambda function that will be used to retrieve info about the AWS Network Firewall in order to setup the routing.
	// Experimental.
	LambdaConfiguration *interfaces.AwsCustomResourceLambdaConfiguration `field:"optional" json:"lambdaConfiguration" yaml:"lambdaConfiguration"`
	// Subnets that should have routes back to the protected subnets.
	//
	// Since
	// traffic is flowing through the firewall, routes should be put into the
	// subnets where traffic is returning to. This is most likely your public
	// subnets in the VPC. By supplying this parameter, routes will be created
	// that send all traffic destined for the `protectedSubnets` back to the
	// firewall for proper routing.
	// Experimental.
	ReturnSubnets *[]awsec2.ISubnet `field:"optional" json:"returnSubnets" yaml:"returnSubnets"`
}

