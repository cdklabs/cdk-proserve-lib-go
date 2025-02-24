package patterns

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Underlying components for the pattern.
// Experimental.
type ApiGatewayStaticHosting_PatternComponents struct {
	// Provides access to the underlying Amazon API Gateway REST API that serves as the distribution endpoint for the static content.
	//
	// WARNING: Making changes to the properties of the underlying components of this pattern may cause it to not
	// behave as expected or designed. You do so at your own risk.
	// Experimental.
	Distribution awsapigateway.RestApi `field:"required" json:"distribution" yaml:"distribution"`
	// Provides access to the underlying AWS Lambda function that proxies the static content from Amazon S3.
	//
	// WARNING: Making changes to the properties of the underlying components of this pattern may cause it to not
	// behave as expected or designed. You do so at your own risk.
	// Experimental.
	Proxy awslambda.Function `field:"required" json:"proxy" yaml:"proxy"`
	// Provides access to the underlying Amazon S3 bucket that stores the static content.
	//
	// WARNING: Making changes to the properties of the underlying components of this pattern may cause it to not
	// behave as expected or designed. You do so at your own risk.
	// Experimental.
	Store awss3.Bucket `field:"required" json:"store" yaml:"store"`
}

