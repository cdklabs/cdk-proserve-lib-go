//go:build no_runtime_type_checking

package constructs

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NetworkFirewallEndpoints) validateGetEndpointIdParameters(availabilityZone *string) error {
	return nil
}

func validateNetworkFirewallEndpoints_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewNetworkFirewallEndpointsParameters(scope Construct, id *string, props *NetworkFirewallEndpointsProps) error {
	return nil
}

