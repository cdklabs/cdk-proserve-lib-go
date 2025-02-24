package patterns

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
)

// Domain configuration when using a Custom Domain Name for Amazon API Gateway.
// Experimental.
type ApiGatewayStaticHosting_CustomDomainConfiguration struct {
	// Options for specifying the custom domain name setup.
	// Experimental.
	Options *awsapigateway.DomainNameOptions `field:"required" json:"options" yaml:"options"`
}

