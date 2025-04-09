package components

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsopensearchservice"
	"github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/types"
)

// Properties for configuring an OpenSearch workflow.
// Experimental.
type OpenSearchWorkflowProps struct {
	// The OpenSearch domain to deploy the workflow to.
	// Experimental.
	Domain awsopensearchservice.IDomain `field:"required" json:"domain" yaml:"domain"`
	// IAM role used for domain authentication.
	// Experimental.
	DomainAuthentication awsiam.IRole `field:"required" json:"domainAuthentication" yaml:"domainAuthentication"`
	// Path to the Flow Framework template file (YAML or JSON).
	// Experimental.
	FlowFrameworkTemplatePath *string `field:"required" json:"flowFrameworkTemplatePath" yaml:"flowFrameworkTemplatePath"`
	// Whether to allow destructive operations like updating/deleting workflows.
	// Experimental.
	AllowDestructiveOperations types.DestructiveOperation `field:"optional" json:"allowDestructiveOperations" yaml:"allowDestructiveOperations"`
	// Optional KMS key for encryption.
	// Experimental.
	Encryption awskms.IKey `field:"optional" json:"encryption" yaml:"encryption"`
	// Optional lambda configuration settings for the custom resource provider.
	// Experimental.
	LambdaConfiguration *types.LambdaConfiguration `field:"optional" json:"lambdaConfiguration" yaml:"lambdaConfiguration"`
	// Optional asset variables for the workflow.
	//
	// This can either be an AWS CDK
	// S3 Asset object or a string that represents an S3 path (e.g. `s3://my-bucket/my-key`).
	// Your template must be configured to accept these variables using
	// `${{{ my_variable }}}` syntax. For each one of these variables, an S3
	// pre-signed URL will be generated and substituted into your template right
	// before workflow creation time. If you provide an S3 path, you must grant
	// read permissions to the appropriate bucket in order for the custom
	// resource to be able to generate a pre-signed url.
	// Experimental.
	TemplateAssetVariables *map[string]interface{} `field:"optional" json:"templateAssetVariables" yaml:"templateAssetVariables"`
	// Optional creation variables for the workflow. Your template must be configured to accept these variables using `${{{ my_variable }}}` syntax.
	//
	// These variables will be substituted in prior to creation, so that will
	// be available during creation time and provision time.
	// Experimental.
	TemplateCreationVariables *map[string]*string `field:"optional" json:"templateCreationVariables" yaml:"templateCreationVariables"`
	// Optional provisioning variables for the workflow. Your template must be configured to accept these variables using `${{ my_variable }}` syntax.
	//
	// https://opensearch.org/docs/latest/automating-configurations/api/provision-workflow
	// Experimental.
	TemplateProvisionVariables *map[string]*string `field:"optional" json:"templateProvisionVariables" yaml:"templateProvisionVariables"`
}

