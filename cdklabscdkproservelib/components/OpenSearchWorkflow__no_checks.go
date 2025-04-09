//go:build no_runtime_type_checking

package components

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_OpenSearchWorkflow) validateGetResourceIdParameters(workflowStepId *string) error {
	return nil
}

func validateOpenSearchWorkflow_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewOpenSearchWorkflowParameters(scope constructs.Construct, id *string, props *OpenSearchWorkflowProps) error {
	return nil
}

