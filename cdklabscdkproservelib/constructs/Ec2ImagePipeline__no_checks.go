//go:build no_runtime_type_checking

package constructs

// Building without runtime type checking enabled, so all the below just return nil

func validateEc2ImagePipeline_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewEc2ImagePipelineParameters(scope Construct, id *string, props *Ec2ImagePipelineProps) error {
	return nil
}

