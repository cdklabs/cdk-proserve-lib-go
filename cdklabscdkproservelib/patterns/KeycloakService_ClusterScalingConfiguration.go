package patterns


// Configuration options for scaling the cluster.
// Experimental.
type KeycloakService_ClusterScalingConfiguration struct {
	// The minimum amount of Keycloak tasks that should be active at any given time.
	// Experimental.
	Maximum *float64 `field:"required" json:"maximum" yaml:"maximum"`
	// The maximum amount of Keycloak tasks that should be active at any given time.
	// Experimental.
	Minimum *float64 `field:"required" json:"minimum" yaml:"minimum"`
	// Configuration options for scaling the cluster based on number of active requests.
	//
	// Scaling is always enabled based on CPU utilization if the scaling bounds have been provided.
	// Experimental.
	RequestCountScaling *KeycloakService_ClusterRequestCountScalingConfiguration `field:"optional" json:"requestCountScaling" yaml:"requestCountScaling"`
}

