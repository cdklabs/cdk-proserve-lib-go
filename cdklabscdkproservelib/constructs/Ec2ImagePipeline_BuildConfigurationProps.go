package constructs


// Experimental.
type Ec2ImagePipeline_BuildConfigurationProps struct {
	// Experimental.
	Start *bool `field:"required" json:"start" yaml:"start"`
	// Experimental.
	Hash *string `field:"optional" json:"hash" yaml:"hash"`
	// Experimental.
	WaitForCompletion *bool `field:"optional" json:"waitForCompletion" yaml:"waitForCompletion"`
}

