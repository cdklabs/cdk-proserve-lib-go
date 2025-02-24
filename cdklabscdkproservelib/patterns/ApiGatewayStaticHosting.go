package patterns

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/patterns/internal"
)

// A pattern that deploys resources to support the hosting of static assets within an AWS account.
//
// Unlike the normal pattern for static content hosting (Amazon S3 fronted by Amazon CloudFront), this pattern instead
// uses a combination of Amazon S3, AWS Lambda, and Amazon API Gateway. This can be useful for rapidly deploying a
// static website that follows best practices when Amazon CloudFront is not available.
//
// The construct also handles encryption for the framework resources using either a provided KMS key or an
// AWS managed key.
//
// There are two methods for exposing the URL to consumers - the default API execution endpoint or via a custom domain
// name setup.
//
// If using the default API execution endpoint, you must provide a base path as this will translate to the
// stage name of the REST API. You must also ensure that all relative links in the static content either reference
// the base path in URLs relative to the root (e.g. preceded by '/') or uses URLs that are relative to the current
// directory (e.g. no '/').
//
// If using the custom domain name, then you do not need to provide a base path and relative links in your static
// content will not require modification. You can choose to specify a base path with this option if so desired - in
// that case, similar rules regarding relative URLs in the static content above must be followed.
//
// Example:
//   import { ApiGatewayStaticHosting } from '@cdklabs/cdk-proserve-lib/patterns';
//   import { EndpointType } from 'aws-cdk-lib/aws-apigateway';
//
//   new ApiGatewayStaticHosting(this, 'MyWebsite', {
//       asset: {
//           id: 'Entry',
//           path: join(__dirname, 'assets', 'website', 'dist'),
//           spaIndexPageName: 'index.html'
//       },
//       domain: {
//           basePath: 'public'
//       },
//       endpoint: {
//           types: [EndpointType.REGIONAL]
//       },
//       retainStoreOnDeletion: true,
//       versionTag: '1.0.2'
//   });
//
// Experimental.
type ApiGatewayStaticHosting interface {
	constructs.Construct
	// Provides access to the underlying components of the pattern as an escape hatch.
	//
	// WARNING: Making changes to the properties of the underlying components of this pattern may cause it to not
	// behave as expected or designed. You do so at your own risk.
	// Experimental.
	Components() *ApiGatewayStaticHosting_PatternComponents
	// Alias domain name for the API that distributes the static content.
	//
	// This is only available if the custom domain name configuration was provided to the pattern. In that event, you
	// would then create either a CNAME or ALIAS record in your DNS system that maps your custom domain name to this
	// value.
	// Experimental.
	CustomDomainNameAlias() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// URL for the API that distributes the static content.
	// Experimental.
	Url() *string
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApiGatewayStaticHosting
type jsiiProxy_ApiGatewayStaticHosting struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_ApiGatewayStaticHosting) Components() *ApiGatewayStaticHosting_PatternComponents {
	var returns *ApiGatewayStaticHosting_PatternComponents
	_jsii_.Get(
		j,
		"components",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStaticHosting) CustomDomainNameAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customDomainNameAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStaticHosting) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStaticHosting) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}


// Creates a new static hosting pattern.
// Experimental.
func NewApiGatewayStaticHosting(scope constructs.Construct, id *string, props *ApiGatewayStaticHostingProps) ApiGatewayStaticHosting {
	_init_.Initialize()

	if err := validateNewApiGatewayStaticHostingParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApiGatewayStaticHosting{}

	_jsii_.Create(
		"@cdklabs/cdk-proserve-lib.patterns.ApiGatewayStaticHosting",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Creates a new static hosting pattern.
// Experimental.
func NewApiGatewayStaticHosting_Override(a ApiGatewayStaticHosting, scope constructs.Construct, id *string, props *ApiGatewayStaticHostingProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-proserve-lib.patterns.ApiGatewayStaticHosting",
		[]interface{}{scope, id, props},
		a,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func ApiGatewayStaticHosting_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApiGatewayStaticHosting_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-proserve-lib.patterns.ApiGatewayStaticHosting",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayStaticHosting) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

