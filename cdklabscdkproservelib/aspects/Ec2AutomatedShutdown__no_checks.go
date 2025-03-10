//go:build no_runtime_type_checking

package aspects

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_Ec2AutomatedShutdown) validateVisitParameters(node constructs.IConstruct) error {
	return nil
}

func validateNewEc2AutomatedShutdownParameters(props *Ec2AutomatedShutdownProps) error {
	return nil
}

