package aspects


// Experimental.
type SecurityCompliance_LambdaSettings struct {
	// Enables reserved concurrent executions for Lambda Functions.
	//
	// Resolves:
	//   - NIST.800.53.R5-LambdaConcurrency
	//
	// Defaults to 1 if not disabled or set.
	// Experimental.
	ReservedConcurrentExecutions *SecurityCompliance_ReservedConcurrentSettings `field:"optional" json:"reservedConcurrentExecutions" yaml:"reservedConcurrentExecutions"`
}

