package patterns


// Configuration options for scaling the cluster based on number of active requests.
// Experimental.
type KeycloakService_ClusterRequestCountScalingConfiguration struct {
	// Whether to enable scaling based on the number of active requests.
	//
	// Scaling is always enabled based on CPU utilization if the scaling bounds have been provided.
	// Experimental.
	Enabled *bool `field:"required" json:"enabled" yaml:"enabled"`
	// The pivotal number of active requests through the load balancer before a scaling action is triggered.
	//
	// Used
	// to fine-tune scaling to your specific capacity needs.
	//
	// If not specified but auto scaling is enabled, then by default scaling out will occur when the number of
	// active requests exceeds 80 and scaling in will occur when this number drops below 80. All scaling activities
	// incur a 5 minute cooldown period.
	// Experimental.
	Threshold *float64 `field:"optional" json:"threshold" yaml:"threshold"`
}

