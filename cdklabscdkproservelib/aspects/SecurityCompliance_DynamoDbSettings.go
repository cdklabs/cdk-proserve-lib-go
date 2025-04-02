package aspects


// Experimental.
type SecurityCompliance_DynamoDbSettings struct {
	// Enables Point-in-Time Recovery for DynamoDB tables.
	//
	// Resolves:
	//   - AwsSolutions-DDB3
	//
	// Defaults to true if not disabled.
	// Experimental.
	PointInTimeRecovery *SecurityCompliance_DisableableSetting `field:"optional" json:"pointInTimeRecovery" yaml:"pointInTimeRecovery"`
}

