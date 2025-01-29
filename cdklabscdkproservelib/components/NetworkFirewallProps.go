package components

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Properties for configuring a NetworkFirewall.
// Experimental.
type NetworkFirewallProps struct {
	// List of subnets where the Network Firewall will be placed These should typically be dedicated firewall subnets.
	// Experimental.
	FirewallSubnets *[]awsec2.ISubnet `field:"required" json:"firewallSubnets" yaml:"firewallSubnets"`
	// The capacity to set for the Suricata rule group.
	//
	// This cannot be modified
	// after creation. You should set this to the upper bound of what you expect
	// your firewall rule group to consume.
	// Experimental.
	SuricataRulesCapacity *float64 `field:"required" json:"suricataRulesCapacity" yaml:"suricataRulesCapacity"`
	// Path to the Suricata rules file on the local file system.
	// Experimental.
	SuricataRulesFilePath *string `field:"required" json:"suricataRulesFilePath" yaml:"suricataRulesFilePath"`
	// VPC where the Network Firewall will be deployed.
	// Experimental.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// Network Firewall routing configuration.
	//
	// By configuring these settings,
	// the Construct will automatically setup basic routing statements for you
	// for the provided subnets. This should be used with caution and you should
	// double check the routing is correct prior to deployment.
	// Experimental.
	ConfigureVpcRoutes *NetworkFirewall_NetworkFirewallVpcRouteProps `field:"optional" json:"configureVpcRoutes" yaml:"configureVpcRoutes"`
	// Optional logging configuration for the Network Firewall.
	//
	// If not provided,
	// logs will not be written.
	// Experimental.
	Logging *NetworkFirewall_LoggingConfiguration `field:"optional" json:"logging" yaml:"logging"`
}

