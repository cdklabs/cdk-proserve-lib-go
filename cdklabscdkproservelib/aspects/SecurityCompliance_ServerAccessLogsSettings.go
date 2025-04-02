package aspects


// Experimental.
type SecurityCompliance_ServerAccessLogsSettings struct {
	// The bucket where server access logs will be sent.
	//
	// This must be configured
	// with the correct permissions to allow the target bucket to receive logs.
	//
	// If not specified, server access logs will not be enabled.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/userguide/enable-server-access-logging.html
	//
	// Experimental.
	DestinationBucketName *string `field:"required" json:"destinationBucketName" yaml:"destinationBucketName"`
}

