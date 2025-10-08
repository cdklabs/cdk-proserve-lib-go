//go:build no_runtime_type_checking

package aspects

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RdsOracleMultiTenant) validateVisitParameters(node constructs.IConstruct) error {
	return nil
}

func validateNewRdsOracleMultiTenantParameters(props *RdsOracleMultiTenantProps) error {
	return nil
}

