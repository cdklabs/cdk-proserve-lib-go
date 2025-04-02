package aspects


// Experimental.
type SecurityComplianceProps struct {
	// Settings for the aspect.
	// Experimental.
	Settings *SecurityCompliance_Settings `field:"optional" json:"settings" yaml:"settings"`
	// Suppressions to add for CDK Nag.
	//
	// You must add your own reasoning to each
	// suppression. These helpers have been created for common nag suppression
	// use-cases. It is recommended to review the suppressions that are added
	// and ensure that they adhere to your organizational level of acceptance.
	// Each suppression must be supplied with a reason for the suppression as
	// a string to each suppression property.
	//
	// If you are not using CDK Nag or do not want to use any suppressions, you
	// can ignore this property.
	// Experimental.
	Suppressions *SecurityCompliance_Suppressions `field:"optional" json:"suppressions" yaml:"suppressions"`
}

