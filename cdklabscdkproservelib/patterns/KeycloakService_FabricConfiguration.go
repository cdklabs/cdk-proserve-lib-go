package patterns


// Configuration options for the fabric.
// Experimental.
type KeycloakService_FabricConfiguration struct {
	// If specified, an Application Load Balancer will be used for the Keycloak service endpoint instead of a Network Load Balancer.
	//
	// This is useful if you want to have fine grain control over the routes exposed as well
	// as implement application-based firewall rules.
	//
	// The default is to use a Network Load Balancer (Layer 4) with TCP passthrough for performance.
	//
	// NOTE: If you switch to application (layer 7) load balancing, you will not be able to perform mutual TLS
	// authentication and authorization flows at the Keycloak service itself as SSL will be terminated at the load
	// balancer and re-encrypted to the backend which will drop the client certificate.
	// Experimental.
	ApplicationLoadBalancing *KeycloakService_FabricApplicationLoadBalancingConfiguration `field:"optional" json:"applicationLoadBalancing" yaml:"applicationLoadBalancing"`
	// Name of the Route53 DNS Zone where the Keycloak hostnames should be automatically configured if provided.
	//
	// By default, no Route53 records will be created.
	//
	// Example:
	//   example.com
	//
	// Experimental.
	DnsZoneName *string `field:"optional" json:"dnsZoneName" yaml:"dnsZoneName"`
	// Whether or not the load balancer should be exposed to the external network.
	// Default: false.
	//
	// Experimental.
	InternetFacing *bool `field:"optional" json:"internetFacing" yaml:"internetFacing"`
}

