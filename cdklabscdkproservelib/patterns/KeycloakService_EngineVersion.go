package patterns

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/jsii"
)

// Versions of the Keycloak application.
// Experimental.
type KeycloakService_EngineVersion interface {
	// The full version string.
	// Experimental.
	Value() *string
	// Determines if the KeycloakVersion matches a specific version.
	//
	// Returns: True if the current version matches the target version, false otherwise.
	// Experimental.
	Is(keycloak KeycloakService_EngineVersion) *bool
}

// The jsii proxy struct for KeycloakService_EngineVersion
type jsiiProxy_KeycloakService_EngineVersion struct {
	_ byte // padding
}

func (j *jsiiProxy_KeycloakService_EngineVersion) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Create a new KeycloakVersion with an arbitrary version.
//
// Returns: KeycloakVersion.
// Experimental.
func KeycloakService_EngineVersion_Of(version *string) KeycloakService_EngineVersion {
	_init_.Initialize()

	if err := validateKeycloakService_EngineVersion_OfParameters(version); err != nil {
		panic(err)
	}
	var returns KeycloakService_EngineVersion

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-proserve-lib.patterns.KeycloakService.EngineVersion",
		"of",
		[]interface{}{version},
		&returns,
	)

	return returns
}

func KeycloakService_EngineVersion_V26_3_2() KeycloakService_EngineVersion {
	_init_.Initialize()
	var returns KeycloakService_EngineVersion
	_jsii_.StaticGet(
		"@cdklabs/cdk-proserve-lib.patterns.KeycloakService.EngineVersion",
		"V26_3_2",
		&returns,
	)
	return returns
}

func (k *jsiiProxy_KeycloakService_EngineVersion) Is(keycloak KeycloakService_EngineVersion) *bool {
	if err := k.validateIsParameters(keycloak); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		k,
		"is",
		[]interface{}{keycloak},
		&returns,
	)

	return returns
}

