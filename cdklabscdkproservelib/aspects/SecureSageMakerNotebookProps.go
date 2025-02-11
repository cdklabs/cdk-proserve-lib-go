package aspects

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Experimental.
type SecureSageMakerNotebookProps struct {
	// Sets the VPC Subnets that the SageMaker Notebook Instance is allowed to launch training and inference jobs into.
	//
	// This is enforced by adding
	// DENY statements to the existing role that the Notebook Instance is using.
	// Experimental.
	AllowedLaunchSubnets *[]awsec2.ISubnet `field:"required" json:"allowedLaunchSubnets" yaml:"allowedLaunchSubnets"`
	// Sets the VPC Subnet for the Sagemaker Notebook Instance.
	//
	// This ensures the
	// notebook is locked down to a specific VPC/subnet.
	// Experimental.
	NotebookSubnet awsec2.ISubnet `field:"required" json:"notebookSubnet" yaml:"notebookSubnet"`
	// Sets the `directInternetAccess` property on the SageMaker Notebooks.
	//
	// By default, this is set to false to disable internet access on any
	// SageMaker Notebook Instance that this aspect is applied to.
	// Default: false.
	//
	// Experimental.
	DirectInternetAccess *bool `field:"optional" json:"directInternetAccess" yaml:"directInternetAccess"`
	// Sets the `rootAccess` property on the SageMaker Notebooks.
	//
	// By default, this is set to false to disable root access on any
	// SageMaker Notebook Instance that this aspect is applied to.
	// Default: false.
	//
	// Experimental.
	RootAccess *bool `field:"optional" json:"rootAccess" yaml:"rootAccess"`
}

