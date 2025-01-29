//go:build !no_runtime_type_checking

package components

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (n *jsiiProxy_NetworkFirewallEndpoints) validateGetEndpointIdParameters(availabilityZone *string) error {
	if availabilityZone == nil {
		return fmt.Errorf("parameter availabilityZone is required, but nil was provided")
	}

	return nil
}

func validateNetworkFirewallEndpoints_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewNetworkFirewallEndpointsParameters(scope constructs.Construct, id *string, props *NetworkFirewallEndpointsProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

