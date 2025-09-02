package patterns

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
)

// Options related to Keycloak.
// Experimental.
type KeycloakService_KeycloakProps struct {
	// Configuration for Keycloak.
	// Experimental.
	Configuration *KeycloakService_ApplicationConfiguration `field:"required" json:"configuration" yaml:"configuration"`
	// Keycloak container image to use.
	// Experimental.
	Image awsecs.ContainerImage `field:"required" json:"image" yaml:"image"`
	// Keycloak version.
	// Experimental.
	Version KeycloakService_EngineVersion `field:"required" json:"version" yaml:"version"`
}

