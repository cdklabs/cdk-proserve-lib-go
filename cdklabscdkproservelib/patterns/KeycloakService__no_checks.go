//go:build no_runtime_type_checking

package patterns

// Building without runtime type checking enabled, so all the below just return nil

func validateKeycloakService_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewKeycloakServiceParameters(scope constructs.Construct, id *string, props *KeycloakServiceProps) error {
	return nil
}

