package patterns


// Configuration options for scaling the tasks.
// Experimental.
type KeycloakService_TaskSizingConfiguration struct {
	// vCPU allocation for each task.
	//
	// Values match the permitted values for `FargateTaskDefinitionProps.cpu`
	//
	// By default 1 vCPU (1024) is allocated.
	// See: https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.aws_ecs.FargateTaskDefinition.html#cpu
	//
	// Default: 1024.
	//
	// Experimental.
	Cpu *float64 `field:"optional" json:"cpu" yaml:"cpu"`
	// Memory allocation in MiB for each task.
	//
	// Values match the permitted values for `FargateTaskDefinitionProps.memoryLimitMiB`
	//
	// By default 2048 MiB (2GB) is allocated.
	// See: https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.aws_ecs.FargateTaskDefinition.html#memorylimitmib
	//
	// Default: 2048.
	//
	// Experimental.
	MemoryMb *float64 `field:"optional" json:"memoryMb" yaml:"memoryMb"`
}

