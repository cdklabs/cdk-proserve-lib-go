package components

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/components/internal"
)

// Controls the contents of an Amazon OpenSearch Service domain from Infrastructure as Code.
//
// This construct allows you to manage indices, component/index templates, cluster settings, Index State Management
// (ISM) policies, role / role mappings, and saved objects for a managed OpenSearch domain from CDK. Within your
// repository, you would create a directory containing the following sub-directories:
//
// - indices
// - ism-policies
// - role-mappings
// - roles
// - saved-objects
// - templates
//   - component
//   - index
//
// Within each subfolder, you can add JSON files to represent objects you want to be provisioned. The schema of the
// JSON file will be specific to the entity being provisioned and you can find more information withiin the OpenSearch
// documentation. The name of each file will be used as the name of the entity that is created within OpenSearch.
//
// The role-mappings entity is special and its structure is not found in the OpenSearch documentation. The name of the
// file should be the name of an internal OpenSearch role and will be used to send a PUT request to
// `/_plugins/_security/api/rolesmapping/<name>`. The contents of the file should be backend role names to map to the
// internal OpenSearch role (where each backend role appears on a separate line). These backend roles can be LDAP
// Distinguished Names, AWS Identity and Access Management (IAM) Role ARNs, etc.
//
// The custom resource property 'dynamicRoleMappings' allows you to supplement role mappings at CDK deployment time.
// This is useful in situations where you are dynamically creating the backend role as part of IaC and its identifier
// will not be known ahead of time. For example, if you create an AWS IAM Role that will be mapped to an internal
// OpenSearch role like `all_access` via CDK, you can pass that Role ARN to the resource through `dynamicRoleMappings`
// as such:
//
// ```
// dynamicRoleMappings: {
//     all_access: [myRole.roleArn]
// }
// ```
//
// The property allows you to map multiple backend roles to a single internal OpenSearch role hence the value being a
// list of strings.
//
// The custom resource proeprty `clusterSettings` allows you to dynamicall configure cluster settings via IaC. Note
// that not all OpenSearch settings will be configurable in the managed Amazon OpenSearch Service and you receive an
// error when trying to do so. Additional details can be found
// [here](https://docs.opensearch.org/docs/latest/api-reference/cluster-api/cluster-settings/)
//
// By default, the custom resource will only modify the domain during AWS CloudFormation CREATE calls. This is to
// prevent potential data loss or issues as the domain will most likely drift from its initial provisioning
// configuration once established and used. If you would like to allow the custom resource to manage the domain
// provisioning during other CloudForamtion lifecycle events, you can do so by setting the `allowDestructiveOperations`
// property on the custom resource.
//
// The construct also handles encryption for the framework resources using either a provided KMS key or an
// AWS managed key.
//
// The recommended pattern for provisioning a managed OpenSearch domain is to leverage this custom resource in a
// separate CDK stack from the one that deploys your domain. Typically OpenSearch domain deployments and teardowns
// take a significant amount of time and so you want to minimize errors in the stack that deploys your domain to
// prevent rollbacks and the need to redeploy. By separating your domain creation and provisioning, failures in
// provisioning will not cause the domain to be destroyed and will save a significant amount of development time.
//
// Example:
//   import { join } from 'node:path';
//   import { OpenSearchProvisionDomain } from '@cdklabs/cdk-proserve-lib/constructs';
//   import { DestructiveOperation } from '@cdklabs/cdk-proserve-lib/types';
//   import { Role } from 'aws-cdk-lib/aws-iam';
//   import { Domain } from 'aws-cdk-lib/aws-opensearchservice';
//
//   const domain = Domain.fromDomainAttributes(this, 'Domain', {
//       domainArn: 'XXXX',
//       domainEndpoint: 'XXXX'
//   });
//
//   const admin = Role.fromRoleArn(this, 'DomainAdmin', 'XXXX');
//   const user = Role.fromRoleArn(this, 'DomainUser', 'XXXX');
//
//   new OpenSearchProvisionDomain(this, 'ProvisionDomain', {
//       domain: domain,
//       domainAdmin: admin,
//       provisioningConfigurationPath: join(
//           __dirname,
//           '..',
//           'dist',
//           'cluster-configuration'
//       ),
//       allowDestructiveOperations: DestructiveOperation.UPDATE,
//       clusterSettings: {
//           persistent: {
//               'plugins.ml_commons.model_access_control_enabled': 'true'
//           }
//       },
//       dynamicRoleMappings: {
//           all_access: [user.roleArn]
//       }
//   });
//
// Experimental.
type OpenSearchProvisionDomain interface {
	constructs.Construct
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OpenSearchProvisionDomain
type jsiiProxy_OpenSearchProvisionDomain struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_OpenSearchProvisionDomain) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Provisions an existing Amazon OpenSearch Service domain with user-specified data.
// Experimental.
func NewOpenSearchProvisionDomain(scope constructs.Construct, id *string, props *OpenSearchProvisionDomainProps) OpenSearchProvisionDomain {
	_init_.Initialize()

	if err := validateNewOpenSearchProvisionDomainParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_OpenSearchProvisionDomain{}

	_jsii_.Create(
		"@cdklabs/cdk-proserve-lib.constructs.OpenSearchProvisionDomain",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Provisions an existing Amazon OpenSearch Service domain with user-specified data.
// Experimental.
func NewOpenSearchProvisionDomain_Override(o OpenSearchProvisionDomain, scope constructs.Construct, id *string, props *OpenSearchProvisionDomainProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-proserve-lib.constructs.OpenSearchProvisionDomain",
		[]interface{}{scope, id, props},
		o,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func OpenSearchProvisionDomain_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOpenSearchProvisionDomain_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-proserve-lib.constructs.OpenSearchProvisionDomain",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenSearchProvisionDomain) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

