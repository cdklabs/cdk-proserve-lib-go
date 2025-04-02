package aspects


// Experimental.
type SecurityCompliance_DisableableSetting struct {
	// Sets the setting to disabled.
	//
	// This does not actually make an impact on
	// the setting itself, it just stops this aspect from making changes to
	// the specific setting.
	// Experimental.
	Disabled *bool `field:"optional" json:"disabled" yaml:"disabled"`
}

