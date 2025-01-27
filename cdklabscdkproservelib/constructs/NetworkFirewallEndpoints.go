package constructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/constructs/internal"
)

// Retrieves Network Firewall endpoints so that you can reference them in your other resources.
//
// Uses an AWS Custom Resource to fetch endpoint information from the Network
// Firewall service. This is useful so that you can both create a Network
// Firewall and reference the endpoints it creates, to do things like configure
// routing to the firewall.
//
// Example:
//   import { CfnOutput } from 'aws-cdk-lib';
//   import { NetworkFirewallEndpoints } from '@cdklabs/cdk-proserve-lib/constructs';
//
//   const endpoints = new NetworkFirewallEndpoints(this, 'Endpoints', {
//     firewall: cfnFirewall,  // CfnFirewall resource to find endpoints for
//   });
//   const az1EndpointId = endpoints.getEndpointId('us-east-1a');
//   const az2EndpointId = endpoints.getEndpointId('us-east-1b');
//   new CfnOutput(this, 'Az1EndpointId', { value: az1Endpoint });
//   new CfnOutput(this, 'Az2EndpointId', { value: az2Endpoint });
//
// Experimental.
type NetworkFirewallEndpoints interface {
	Construct
	// The tree node.
	// Experimental.
	Node() Node
	// Gets the endpoint ID for a specific availability zone.
	//
	// Returns: The endpoint ID for the specified availability zone.
	// Experimental.
	GetEndpointId(availabilityZone *string) *string
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetworkFirewallEndpoints
type jsiiProxy_NetworkFirewallEndpoints struct {
	internal.Type__Construct
}

func (j *jsiiProxy_NetworkFirewallEndpoints) Node() Node {
	var returns Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Retrieves Network Firewall endpoints so that you can reference them in your other resources.
// Experimental.
func NewNetworkFirewallEndpoints(scope Construct, id *string, props *NetworkFirewallEndpointsProps) NetworkFirewallEndpoints {
	_init_.Initialize()

	if err := validateNewNetworkFirewallEndpointsParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkFirewallEndpoints{}

	_jsii_.Create(
		"@cdklabs/cdk-proserve-lib.constructs.NetworkFirewallEndpoints",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Retrieves Network Firewall endpoints so that you can reference them in your other resources.
// Experimental.
func NewNetworkFirewallEndpoints_Override(n NetworkFirewallEndpoints, scope Construct, id *string, props *NetworkFirewallEndpointsProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-proserve-lib.constructs.NetworkFirewallEndpoints",
		[]interface{}{scope, id, props},
		n,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func NetworkFirewallEndpoints_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetworkFirewallEndpoints_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-proserve-lib.constructs.NetworkFirewallEndpoints",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkFirewallEndpoints) GetEndpointId(availabilityZone *string) *string {
	if err := n.validateGetEndpointIdParameters(availabilityZone); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getEndpointId",
		[]interface{}{availabilityZone},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkFirewallEndpoints) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

