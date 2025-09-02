package patterns

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Properties for the KeycloakService construct.
// Experimental.
type KeycloakServiceProps struct {
	// Options related to Keycloak.
	// Experimental.
	Keycloak *KeycloakService_KeycloakProps `field:"required" json:"keycloak" yaml:"keycloak"`
	// Network where Keycloak should be deployed.
	// Experimental.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// Key for encrypting resource data.
	//
	// If not specified, a new key will be created.
	// Experimental.
	Encryption awskms.IKey `field:"optional" json:"encryption" yaml:"encryption"`
	// How long to retain logs for all components.
	//
	// If not specified, logs will be retained for one week.
	// Experimental.
	LogRetentionDuration awslogs.RetentionDays `field:"optional" json:"logRetentionDuration" yaml:"logRetentionDuration"`
	// Overrides for prescribed defaults for the infrastructure.
	// Experimental.
	Overrides *KeycloakService_InfrastructureConfiguration `field:"optional" json:"overrides" yaml:"overrides"`
	// Policies to lifecycle various components of the pattern during stack actions.
	//
	// If not specified, resources will be retained.
	// Experimental.
	RemovalPolicies *KeycloakService_RemovalPolicies `field:"optional" json:"removalPolicies" yaml:"removalPolicies"`
}

