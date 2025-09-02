package patterns

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds"
)

// Configuration options for a non-serverless database model.
// Experimental.
type KeycloakService_TraditionalDatabaseConfiguration struct {
	// Backup lifecycle plan for the database.
	//
	// If not specified, CDK defaults are used.
	// See: https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.aws_rds.DatabaseCluster.html#backup
	//
	// Experimental.
	Backup *awsrds.BackupProps `field:"optional" json:"backup" yaml:"backup"`
	// How long to retain logs for the database and supporting infrastructure.
	// Default: RetentionDays.ONE_WEEK
	//
	// Experimental.
	LogRetentionDuration awslogs.RetentionDays `field:"optional" json:"logRetentionDuration" yaml:"logRetentionDuration"`
	// Whether a ServerlessV2 Aurora database should be deployed or not.
	// Default: true.
	//
	// Experimental.
	Serverless *bool `field:"optional" json:"serverless" yaml:"serverless"`
	// Alternate database engine version to use.
	// Default: AuroraPostgresEngineVersion.VER_17_5
	//
	// Experimental.
	VersionOverride awsrds.AuroraPostgresEngineVersion `field:"optional" json:"versionOverride" yaml:"versionOverride"`
}

