package patterns

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// Configuration for the Keycloak application.
// Experimental.
type KeycloakService_ApplicationConfiguration struct {
	// Hostname configuration for Keycloak.
	// Experimental.
	Hostnames *KeycloakService_HostnameConfiguration `field:"required" json:"hostnames" yaml:"hostnames"`
	// Credentials for bootstrapping a local admin user within Keycloak.
	//
	// Must be a key-value secret with `username` and `password` fields
	//
	// [Guide: Bootstrapping an Admin Account](https://www.keycloak.org/server/hostname)
	//
	// By default, a new secret will be created with a username and randomly generated password.
	// Experimental.
	AdminUser awssecretsmanager.ISecret `field:"optional" json:"adminUser" yaml:"adminUser"`
	// Level of information for Keycloak to log.
	// Default: warn.
	//
	// Experimental.
	LoggingLevel *string `field:"optional" json:"loggingLevel" yaml:"loggingLevel"`
	// Configuration options for the management interface.
	//
	// If not specified, the management interface is disabled.
	// Experimental.
	Management *KeycloakService_ManagementConfiguration `field:"optional" json:"management" yaml:"management"`
	// Optional alternative relative path for serving content.
	// Default: /.
	//
	// Experimental.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Port to serve the standard HTTPS web traffic on.
	// Default: 443.
	//
	// Experimental.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

