//go:build !no_runtime_type_checking

package aspects

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (e *jsiiProxy_Ec2AutomatedShutdown) validateVisitParameters(node constructs.IConstruct) error {
	if node == nil {
		return fmt.Errorf("parameter node is required, but nil was provided")
	}

	return nil
}

func validateNewEc2AutomatedShutdownParameters(props *Ec2AutomatedShutdownProps) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

