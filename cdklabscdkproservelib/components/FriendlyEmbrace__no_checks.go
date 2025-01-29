//go:build no_runtime_type_checking

package components

// Building without runtime type checking enabled, so all the below just return nil

func validateFriendlyEmbrace_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewFriendlyEmbraceParameters(scope constructs.Construct, id *string, props *FriendlyEmbraceProps) error {
	return nil
}

