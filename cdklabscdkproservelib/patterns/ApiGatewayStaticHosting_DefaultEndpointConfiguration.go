package patterns


// Domain configuration when using the Amazon API Gateway Default Execution Endpoint.
// Experimental.
type ApiGatewayStaticHosting_DefaultEndpointConfiguration struct {
	// Base path where all assets will be located.
	//
	// This is because the default execution endpoint does not serve content at the root but off of a stage. As
	// such this base path will be used to create the deployment stage to serve assets from.
	//
	// Example:
	//   /dev/site1
	//
	// Experimental.
	BasePath *string `field:"required" json:"basePath" yaml:"basePath"`
}

