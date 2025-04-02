package aspects


// Experimental.
type SecurityCompliance_ReservedConcurrentSettings struct {
	// Sets the setting to disabled.
	//
	// This does not actually make an impact on
	// the setting itself, it just stops this aspect from making changes to
	// the specific setting.
	// Experimental.
	Disabled *bool `field:"optional" json:"disabled" yaml:"disabled"`
	// The number of reserved concurrency executions.
	//
	// Resolves:
	//   - NIST.800.53.R5-LambdaConcurrency
	//
	// Defaults to 1 if not specified.
	// Experimental.
	ConcurrentExecutionCount *float64 `field:"optional" json:"concurrentExecutionCount" yaml:"concurrentExecutionCount"`
}

