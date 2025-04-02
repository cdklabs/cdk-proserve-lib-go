//go:build no_runtime_type_checking

package aspects

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SecurityCompliance) validateVisitParameters(node constructs.IConstruct) error {
	return nil
}

func validateNewSecurityComplianceParameters(props *SecurityComplianceProps) error {
	return nil
}

