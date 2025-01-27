package constructs


// Enum representing possible override actions for WAF rules.
// Experimental.
type WebApplicationFirewall_OverrideAction string

const (
	// Experimental.
	WebApplicationFirewall_OverrideAction_ALLOW WebApplicationFirewall_OverrideAction = "ALLOW"
	// Experimental.
	WebApplicationFirewall_OverrideAction_BLOCK WebApplicationFirewall_OverrideAction = "BLOCK"
	// Experimental.
	WebApplicationFirewall_OverrideAction_COUNT WebApplicationFirewall_OverrideAction = "COUNT"
)

