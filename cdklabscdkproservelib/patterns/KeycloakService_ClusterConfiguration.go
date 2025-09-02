package patterns


// Configuration options for the cluster.
// Experimental.
type KeycloakService_ClusterConfiguration struct {
	// Boundaries for cluster scaling.
	//
	// If not specified, auto scaling is disabled.
	// Experimental.
	Scaling *KeycloakService_ClusterScalingConfiguration `field:"optional" json:"scaling" yaml:"scaling"`
	// Resource allocation options for each Keycloak task.
	//
	// If not specified, each task gets 1 vCPU and 2GB memory
	//
	// Guidance on sizing can be found [here](https://www.keycloak.org/high-availability/concepts-memory-and-cpu-sizing)
	// Experimental.
	Sizing *KeycloakService_TaskSizingConfiguration `field:"optional" json:"sizing" yaml:"sizing"`
}

