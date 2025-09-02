package patterns


// Experimental.
type KeycloakService_ManagementConfiguration struct {
	// Port to serve the management web traffic on.
	//
	// Example:
	//   9006
	//
	// Experimental.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Whether the health management API is enabled.
	// Default: false.
	//
	// Experimental.
	Health *bool `field:"optional" json:"health" yaml:"health"`
	// Whether the metrics management API is enabled.
	// Default: false.
	//
	// Experimental.
	Metrics *bool `field:"optional" json:"metrics" yaml:"metrics"`
	// Optional alternative relative path for serving content specifically for management.
	// Default: /.
	//
	// Experimental.
	Path *string `field:"optional" json:"path" yaml:"path"`
}

