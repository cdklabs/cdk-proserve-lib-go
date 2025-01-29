package components


// Configuration for rule overrides.
// Experimental.
type WebApplicationFirewall_OverrideConfig struct {
	// The action to take for the specific rule.
	// Experimental.
	Action WebApplicationFirewall_OverrideAction `field:"required" json:"action" yaml:"action"`
	// The name of the specific rule to override.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
}

