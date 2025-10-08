package aspects

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/aspects/internal"
)

// Enables Oracle MultiTenant configuration on RDS Oracle database instances.
//
// This Aspect will apply Oracle MultiTenant configuration to multiple RDS Oracle instances across a CDK
// application automatically. When applied to a construct tree, it identifies all RDS Oracle database
// instances and enables MultiTenant architecture on each one.
//
// **NOTE: This should ONLY be used on new Oracle RDS databases, as it takes a backup and can take a
// significant amount of time to complete. This is a 1-way door, after this setting is turned on it
// CANNOT be reversed!**
//
// Example:
//   // Basic usage applied to an entire CDK application:
//
//   import { App, Aspects } from 'aws-cdk-lib';
//   import { RdsOracleMultiTenant } from '@cdklabs/cdk-proserve-lib/aspects';
//
//   const app = new App();
//
//   // Apply to all Oracle instances in the application
//   Aspects.of(app).add(new RdsOracleMultiTenant());
//
// See: {@link https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-multitenant.html Oracle MultiTenant on Amazon RDS}
//
// Experimental.
type RdsOracleMultiTenant interface {
	awscdk.IAspect
	// Configuration properties for the Aspect.
	// Experimental.
	Props() *RdsOracleMultiTenantProps
	// Visits a construct node and applies Oracle MultiTenant configuration if applicable.
	// Experimental.
	Visit(node constructs.IConstruct)
}

// The jsii proxy struct for RdsOracleMultiTenant
type jsiiProxy_RdsOracleMultiTenant struct {
	internal.Type__awscdkIAspect
}

func (j *jsiiProxy_RdsOracleMultiTenant) Props() *RdsOracleMultiTenantProps {
	var returns *RdsOracleMultiTenantProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}


// Creates a new RDS Oracle MultiTenant Aspect that automatically enables Oracle MultiTenant configuration on RDS Oracle database instances found in the construct tree.
// Experimental.
func NewRdsOracleMultiTenant(props *RdsOracleMultiTenantProps) RdsOracleMultiTenant {
	_init_.Initialize()

	if err := validateNewRdsOracleMultiTenantParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_RdsOracleMultiTenant{}

	_jsii_.Create(
		"@cdklabs/cdk-proserve-lib.aspects.RdsOracleMultiTenant",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Creates a new RDS Oracle MultiTenant Aspect that automatically enables Oracle MultiTenant configuration on RDS Oracle database instances found in the construct tree.
// Experimental.
func NewRdsOracleMultiTenant_Override(r RdsOracleMultiTenant, props *RdsOracleMultiTenantProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-proserve-lib.aspects.RdsOracleMultiTenant",
		[]interface{}{props},
		r,
	)
}

func (r *jsiiProxy_RdsOracleMultiTenant) Visit(node constructs.IConstruct) {
	if err := r.validateVisitParameters(node); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"visit",
		[]interface{}{node},
	)
}

