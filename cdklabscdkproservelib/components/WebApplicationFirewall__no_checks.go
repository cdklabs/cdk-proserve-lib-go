//go:build no_runtime_type_checking

package components

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WebApplicationFirewall) validateAssociateParameters(resource awselasticloadbalancingv2.IApplicationLoadBalancer) error {
	return nil
}

func validateWebApplicationFirewall_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewWebApplicationFirewallParameters(scope constructs.Construct, id *string, props *WebApplicationFirewallProps) error {
	return nil
}

