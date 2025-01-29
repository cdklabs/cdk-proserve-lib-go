//go:build no_runtime_type_checking

package components

// Building without runtime type checking enabled, so all the below just return nil

func validateNetworkFirewall_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewNetworkFirewallParameters(scope constructs.Construct, id *string, props *NetworkFirewallProps) error {
	return nil
}

