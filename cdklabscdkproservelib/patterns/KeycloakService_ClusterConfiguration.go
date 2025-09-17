package patterns

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
)

// Configuration options for the cluster.
// Experimental.
type KeycloakService_ClusterConfiguration struct {
	// Environment variables to make accessible to the service containers.
	// Experimental.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// Boundaries for cluster scaling.
	//
	// If not specified, auto scaling is disabled.
	// Experimental.
	Scaling *KeycloakService_ClusterScalingConfiguration `field:"optional" json:"scaling" yaml:"scaling"`
	// Environment variables to make accessible to the service containers via secrets.
	// Experimental.
	Secrets *map[string]awsecs.Secret `field:"optional" json:"secrets" yaml:"secrets"`
	// Resource allocation options for each Keycloak task.
	//
	// If not specified, each task gets 1 vCPU and 2GB memory
	//
	// Guidance on sizing can be found [here](https://www.keycloak.org/high-availability/concepts-memory-and-cpu-sizing)
	// Experimental.
	Sizing *KeycloakService_TaskSizingConfiguration `field:"optional" json:"sizing" yaml:"sizing"`
}

