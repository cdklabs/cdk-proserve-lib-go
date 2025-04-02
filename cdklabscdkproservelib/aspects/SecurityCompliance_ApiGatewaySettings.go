package aspects


// Experimental.
type SecurityCompliance_ApiGatewaySettings struct {
	// Enable or disable CloudWatch logging for API Gateway stages.
	//
	// Resolves:
	//   - AwsSolutions-APIG6
	//
	// Defaults to log all errors if not specified or disabled.
	// Experimental.
	StageMethodLogging *SecurityCompliance_StageMethodLogging `field:"optional" json:"stageMethodLogging" yaml:"stageMethodLogging"`
}

