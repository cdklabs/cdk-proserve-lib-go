//go:build !no_runtime_type_checking

package patterns

import (
	"fmt"
)

func (k *jsiiProxy_KeycloakService_EngineVersion) validateIsParameters(keycloak KeycloakService_EngineVersion) error {
	if keycloak == nil {
		return fmt.Errorf("parameter keycloak is required, but nil was provided")
	}

	return nil
}

func validateKeycloakService_EngineVersion_OfParameters(version *string) error {
	if version == nil {
		return fmt.Errorf("parameter version is required, but nil was provided")
	}

	return nil
}

