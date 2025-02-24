package patterns

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/types"
)

// Properties for configuring the static hosting pattern.
// Experimental.
type ApiGatewayStaticHostingProps struct {
	// Metadata about the static assets to host.
	// Experimental.
	Asset *ApiGatewayStaticHosting_Asset `field:"required" json:"asset" yaml:"asset"`
	// Configuration information for the distribution endpoint that will be used to serve static content.
	// Experimental.
	Domain interface{} `field:"required" json:"domain" yaml:"domain"`
	// Amazon S3 bucket where access logs should be stored.
	// Default: undefined A new bucket will be created for storing access logs.
	//
	// Experimental.
	AccessLoggingBucket awss3.IBucket `field:"optional" json:"accessLoggingBucket" yaml:"accessLoggingBucket"`
	// Resource access policy to define on the API itself to control who can invoke the endpoint.
	// Experimental.
	AccessPolicy awsiam.PolicyDocument `field:"optional" json:"accessPolicy" yaml:"accessPolicy"`
	// Destination where Amazon API Gateway logs can be sent.
	// Experimental.
	ApiLogDestination awsapigateway.IAccessLogDestination `field:"optional" json:"apiLogDestination" yaml:"apiLogDestination"`
	// Encryption key for protecting the framework resources.
	// Default: undefined AWS service-managed encryption keys will be used where available.
	//
	// Experimental.
	Encryption awskms.IKey `field:"optional" json:"encryption" yaml:"encryption"`
	// Endpoint deployment information for the REST API.
	// Default: undefined Will deploy an edge-optimized API.
	//
	// Experimental.
	Endpoint *awsapigateway.EndpointConfiguration `field:"optional" json:"endpoint" yaml:"endpoint"`
	// Optional configuration settings for the backend handler.
	// Experimental.
	LambdaConfiguration *types.LambdaConfiguration `field:"optional" json:"lambdaConfiguration" yaml:"lambdaConfiguration"`
	// Whether or not to retain the Amazon S3 bucket where static assets are deployed on stack deletion.
	// Default: false The Amazon S3 bucket and all assets contained within will be deleted.
	//
	// Experimental.
	RetainStoreOnDeletion *bool `field:"optional" json:"retainStoreOnDeletion" yaml:"retainStoreOnDeletion"`
	// A version identifier to deploy to the Amazon S3 bucket to help with rapid identification of current deployment This will appear as `metadata.json` at the root of the bucket.
	//
	// Example:
	//   1.0.2
	//
	// Default: undefined No version information will be deployed to the Amazon S3 bucket.
	//
	// Experimental.
	VersionTag *string `field:"optional" json:"versionTag" yaml:"versionTag"`
}

