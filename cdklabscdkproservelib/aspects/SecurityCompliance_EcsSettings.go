package aspects


// Experimental.
type SecurityCompliance_EcsSettings struct {
	// Enables container insights for ECS clusters.
	//
	// Resolves:
	//   - AwsSolutions-ECS4
	//
	// Defaults to ContainerInsights.ENABLED if not disabled.
	// Experimental.
	ClusterContainerInsights *SecurityCompliance_DisableableSetting `field:"optional" json:"clusterContainerInsights" yaml:"clusterContainerInsights"`
}

