package patterns


// Details for the Keycloak hostname configuration.
//
// [Guide: Configuring the hostname](https://www.keycloak.org/server/hostname)
// Experimental.
type KeycloakService_HostnameConfiguration struct {
	// Hostname for all endpoints.
	//
	// Example:
	//   auth.example.com
	//
	// Experimental.
	Default *string `field:"required" json:"default" yaml:"default"`
	// Optional hostname for the administration endpoint.
	//
	// This allows for the separation of the user and administration endpoints for increased security
	//
	// By default, the administrative endpoints will use the default hostname unless this is specified.
	//
	// Example:
	//   admin.auth.example.com
	//
	// Experimental.
	Admin *string `field:"optional" json:"admin" yaml:"admin"`
}

