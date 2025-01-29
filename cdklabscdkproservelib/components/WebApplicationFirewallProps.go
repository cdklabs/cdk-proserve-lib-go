package components


// Experimental.
type WebApplicationFirewallProps struct {
	// List of AWS Managed Rule Groups to use for the firewall.
	// Default: [].
	//
	// Experimental.
	AwsManagedRuleGroups *[]interface{} `field:"optional" json:"awsManagedRuleGroups" yaml:"awsManagedRuleGroups"`
	// Whether to enable CloudWatch metrics.
	// Default: false.
	//
	// Experimental.
	CloudWatchMetricsEnabled *bool `field:"optional" json:"cloudWatchMetricsEnabled" yaml:"cloudWatchMetricsEnabled"`
	// Logging configuration for the firewall.
	// Experimental.
	Logging *WebApplicationFirewall_WebApplicationFirewallLoggingConfig `field:"optional" json:"logging" yaml:"logging"`
	// Whether to enable sampled requests.
	// Default: false.
	//
	// Experimental.
	SampledRequestsEnabled *bool `field:"optional" json:"sampledRequestsEnabled" yaml:"sampledRequestsEnabled"`
}

