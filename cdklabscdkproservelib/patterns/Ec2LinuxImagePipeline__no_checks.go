//go:build no_runtime_type_checking

package patterns

// Building without runtime type checking enabled, so all the below just return nil

func validateEc2LinuxImagePipeline_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Ec2LinuxImagePipeline) validateSetImagePipelineArnParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Ec2LinuxImagePipeline) validateSetImagePipelineTopicParameters(val awssns.ITopic) error {
	return nil
}

func validateNewEc2LinuxImagePipelineParameters(scope constructs.Construct, id *string, props *Ec2LinuxImagePipelineProps) error {
	return nil
}

