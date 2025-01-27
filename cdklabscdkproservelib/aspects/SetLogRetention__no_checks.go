//go:build no_runtime_type_checking

package aspects

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SetLogRetention) validateVisitParameters(node constructs.IConstruct) error {
	return nil
}

func validateNewSetLogRetentionParameters(props *SetLogRetentionProps) error {
	return nil
}

