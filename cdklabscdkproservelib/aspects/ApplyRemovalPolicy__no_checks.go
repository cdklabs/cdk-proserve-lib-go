//go:build no_runtime_type_checking

package aspects

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApplyRemovalPolicy) validateVisitParameters(node constructs.IConstruct) error {
	return nil
}

func validateNewApplyRemovalPolicyParameters(props *ApplyRemovalPolicyProps) error {
	return nil
}

