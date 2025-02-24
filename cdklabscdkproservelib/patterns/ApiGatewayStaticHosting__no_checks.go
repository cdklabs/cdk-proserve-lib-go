//go:build no_runtime_type_checking

package patterns

// Building without runtime type checking enabled, so all the below just return nil

func validateApiGatewayStaticHosting_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewApiGatewayStaticHostingParameters(scope constructs.Construct, id *string, props *ApiGatewayStaticHostingProps) error {
	return nil
}

