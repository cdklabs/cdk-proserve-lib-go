package aspects


// Experimental.
type SecurityCompliance_StepFunctionsSettings struct {
	// Enable or disable X-Ray tracing.
	//
	// Resolves:
	//   - AwsSolutions-SF2
	//
	// Defaults to true if not disabled.
	// Experimental.
	Tracing *SecurityCompliance_DisableableSetting `field:"optional" json:"tracing" yaml:"tracing"`
}

