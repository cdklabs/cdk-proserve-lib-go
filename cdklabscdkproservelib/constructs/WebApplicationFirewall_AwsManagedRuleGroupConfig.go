package constructs


// Configuration interface for AWS Managed Rule Groups.
//
// This interface allows you to specify a managed rule group and optionally
// override the default actions for specific rules within that group.
// Experimental.
type WebApplicationFirewall_AwsManagedRuleGroupConfig struct {
	// The AWS Managed Rule Group to apply.
	// Experimental.
	RuleGroup WebApplicationFirewall_AwsManagedRuleGroup `field:"required" json:"ruleGroup" yaml:"ruleGroup"`
	// Optional list of rule action overrides.
	// Experimental.
	RuleGroupActionOverrides *[]*WebApplicationFirewall_OverrideConfig `field:"optional" json:"ruleGroupActionOverrides" yaml:"ruleGroupActionOverrides"`
}

