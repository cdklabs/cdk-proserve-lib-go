//go:build no_runtime_type_checking

package components

// Building without runtime type checking enabled, so all the below just return nil

func validateOpenSearchProvisionDomain_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewOpenSearchProvisionDomainParameters(scope constructs.Construct, id *string, props *OpenSearchProvisionDomainProps) error {
	return nil
}

