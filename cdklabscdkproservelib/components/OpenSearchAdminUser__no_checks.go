//go:build no_runtime_type_checking

package components

// Building without runtime type checking enabled, so all the below just return nil

func validateOpenSearchAdminUser_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewOpenSearchAdminUserParameters(scope constructs.Construct, id *string, props *OpenSearchAdminUserProps) error {
	return nil
}

