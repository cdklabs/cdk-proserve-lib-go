package constructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsnetworkfirewall"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/constructs/internal"
)

// Creates an AWS Network Firewall using a user-supplied Suricata rules file in a VPC.
//
// Follows guidelines that can be found at:.
//
// Example:
//   import { NetworkFirewall } from '@cdklabs/cdk-proserve-lib/constructs';
//
//   new NetworkFirewall(this, 'Firewall', {
//     vpc,
//     firewallSubnets: vpc.selectSubnets({subnetGroupName: 'firewall'}).subnets,
//     suricataRulesFilePath: './firewall-rules-suricata.txt',
//     suricataRulesCapacity: 1000  // perform your own calculation based on the rules
//   });
//
// See: https://aws.github.io/aws-security-services-best-practices/guides/network-firewall/
//
// Experimental.
type NetworkFirewall interface {
	Construct
	// The underlying CloudFormation Network Firewall resource.
	// Experimental.
	Firewall() awsnetworkfirewall.CfnFirewall
	// The tree node.
	// Experimental.
	Node() Node
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetworkFirewall
type jsiiProxy_NetworkFirewall struct {
	internal.Type__Construct
}

func (j *jsiiProxy_NetworkFirewall) Firewall() awsnetworkfirewall.CfnFirewall {
	var returns awsnetworkfirewall.CfnFirewall
	_jsii_.Get(
		j,
		"firewall",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkFirewall) Node() Node {
	var returns Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Creates an AWS Network Firewall using a user-supplied Suricata rules file in a VPC.
// Experimental.
func NewNetworkFirewall(scope Construct, id *string, props *NetworkFirewallProps) NetworkFirewall {
	_init_.Initialize()

	if err := validateNewNetworkFirewallParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkFirewall{}

	_jsii_.Create(
		"@cdklabs/cdk-proserve-lib.constructs.NetworkFirewall",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Creates an AWS Network Firewall using a user-supplied Suricata rules file in a VPC.
// Experimental.
func NewNetworkFirewall_Override(n NetworkFirewall, scope Construct, id *string, props *NetworkFirewallProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-proserve-lib.constructs.NetworkFirewall",
		[]interface{}{scope, id, props},
		n,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func NetworkFirewall_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetworkFirewall_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-proserve-lib.constructs.NetworkFirewall",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkFirewall) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

