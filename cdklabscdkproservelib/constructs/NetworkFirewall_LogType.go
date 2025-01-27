package constructs


// Experimental.
type NetworkFirewall_LogType string

const (
	// Logs for events that are related to TLS inspection.
	//
	// For more information, see Inspecting SSL/TLS traffic with TLS inspection configurations in the Network Firewall Developer Guide .
	// Experimental.
	NetworkFirewall_LogType_TLS NetworkFirewall_LogType = "TLS"
	// Standard network traffic flow logs.
	//
	// The stateful rules engine records flow logs for all network traffic that it receives. Each flow log record captures the network flow for a specific standard stateless rule group.
	// Experimental.
	NetworkFirewall_LogType_FLOW NetworkFirewall_LogType = "FLOW"
	// Logs for traffic that matches your stateful rules and that have an action that sends an alert.
	//
	// A stateful rule sends alerts for the rule actions DROP, ALERT, and REJECT. For more information, see the StatefulRule property.
	// Experimental.
	NetworkFirewall_LogType_ALERT NetworkFirewall_LogType = "ALERT"
)

