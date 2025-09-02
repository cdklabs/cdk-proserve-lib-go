package patterns


// Overrides for prescribed defaults for the infrastructure.
// Experimental.
type KeycloakService_InfrastructureConfiguration struct {
	// Overrides related to the application hosting infrastructure.
	// Experimental.
	Cluster *KeycloakService_ClusterConfiguration `field:"optional" json:"cluster" yaml:"cluster"`
	// Overrides related to the database infrastructure.
	// Experimental.
	Database interface{} `field:"optional" json:"database" yaml:"database"`
	// Overrides related to the networking infrastructure.
	// Experimental.
	Fabric *KeycloakService_FabricConfiguration `field:"optional" json:"fabric" yaml:"fabric"`
}

