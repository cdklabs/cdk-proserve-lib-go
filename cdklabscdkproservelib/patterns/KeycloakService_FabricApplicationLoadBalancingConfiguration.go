package patterns

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
)

// Configuration for using application load balancing (layer 7) for the fabric endpoint.
// Experimental.
type KeycloakService_FabricApplicationLoadBalancingConfiguration struct {
	// TLS certificate to support SSL termination at the load balancer level for the default Keycloak endpoint.
	// Experimental.
	Certificate awscertificatemanager.ICertificate `field:"required" json:"certificate" yaml:"certificate"`
	// TLS certificate to support SSL termination at the load balancer level for the management Keycloak endpoint.
	// Experimental.
	ManagementCertificate awscertificatemanager.ICertificate `field:"optional" json:"managementCertificate" yaml:"managementCertificate"`
}

