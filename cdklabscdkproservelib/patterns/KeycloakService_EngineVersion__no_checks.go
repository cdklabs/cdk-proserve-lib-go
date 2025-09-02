//go:build no_runtime_type_checking

package patterns

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KeycloakService_EngineVersion) validateIsParameters(keycloak KeycloakService_EngineVersion) error {
	return nil
}

func validateKeycloakService_EngineVersion_OfParameters(version *string) error {
	return nil
}

