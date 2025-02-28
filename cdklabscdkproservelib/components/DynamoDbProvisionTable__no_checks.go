//go:build no_runtime_type_checking

package components

// Building without runtime type checking enabled, so all the below just return nil

func validateDynamoDbProvisionTable_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewDynamoDbProvisionTableParameters(scope constructs.Construct, id *string, props *DynamoDbProvisionTableProps) error {
	return nil
}

