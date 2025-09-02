package patterns


// Configuration options for the fabric.
// Experimental.
type KeycloakService_FabricConfiguration struct {
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

