package aspects


// Experimental.
type SecurityCompliance_S3Settings struct {
	// Enable server access logs to a destination S3 bucket.
	//
	// Since this requires
	// a destination S3 bucket, it is not set by default. You must set a target
	// S3 bucket to enable access logs.
	//
	// Resolves:
	//   - AwsSolutions-S1
	//   - NIST.800.53.R5-S3BucketLoggingEnabled
	// Experimental.
	ServerAccessLogs *SecurityCompliance_ServerAccessLogsSettings `field:"optional" json:"serverAccessLogs" yaml:"serverAccessLogs"`
	// Enables versioning for S3 buckets.
	//
	// Resolves:
	//   - NIST.800.53.R5-S3BucketVersioningEnabled
	//
	// Defaults to true if not disabled.
	// Experimental.
	Versioning *SecurityCompliance_DisableableSetting `field:"optional" json:"versioning" yaml:"versioning"`
}

