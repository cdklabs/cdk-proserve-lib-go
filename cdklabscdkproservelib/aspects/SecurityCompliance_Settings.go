package aspects


// Configuration settings for the security-compliance aspect.
//
// This interface provides a centralized way to configure security and compliance
// settings for various AWS resources. Each property corresponds to a specific
// AWS service and contains settings that help ensure resources comply with
// security best practices and compliance requirements.
//
// By default, most security settings are enabled unless explicitly disabled.
// Some settings may require additional configuration to be effective.
//
// Example:
//   const securitySettings: Settings = {
//     lambda: {
//       reservedConcurrentExecutions: {
//         concurrentExecutionCount: 5
//       }
//     },
//     s3: {
//       serverAccessLogs: {
//         destinationBucketName: 'access-logs-bucket'
//       }
//     }
//   };
//
// Experimental.
type SecurityCompliance_Settings struct {
	// Security and compliance settings for API Gateway resources.
	//
	// Controls settings like method logging to ensure proper monitoring
	// and auditability of API usage.
	// Experimental.
	ApiGateway *SecurityCompliance_ApiGatewaySettings `field:"optional" json:"apiGateway" yaml:"apiGateway"`
	// Security and compliance settings for DynamoDB tables.
	//
	// Configures features like Point-in-Time Recovery to improve
	// data durability and recoverability.
	// Experimental.
	DynamoDb *SecurityCompliance_DynamoDbSettings `field:"optional" json:"dynamoDb" yaml:"dynamoDb"`
	// Security and compliance settings for ECS clusters and services.
	//
	// Enables features like Container Insights for better
	// monitoring and observability.
	// Experimental.
	Ecs *SecurityCompliance_EcsSettings `field:"optional" json:"ecs" yaml:"ecs"`
	// Security and compliance settings for Lambda functions.
	//
	// Controls execution limits and other settings to improve
	// the security posture of Lambda functions.
	// Experimental.
	Lambda *SecurityCompliance_LambdaSettings `field:"optional" json:"lambda" yaml:"lambda"`
	// Security and compliance settings for S3 buckets.
	//
	// Configures features like versioning and server access logging
	// to improve data protection and meet compliance requirements.
	// Experimental.
	S3 *SecurityCompliance_S3Settings `field:"optional" json:"s3" yaml:"s3"`
	// Security and compliance settings for Step Functions state machines.
	//
	// Controls settings like X-Ray tracing to improve
	// observability and debugging capabilities.
	// Experimental.
	StepFunctions *SecurityCompliance_StepFunctionsSettings `field:"optional" json:"stepFunctions" yaml:"stepFunctions"`
}

